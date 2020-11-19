package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
	"sync"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/k0kubun/pp"
	"github.com/olekukonko/tablewriter"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type reviewRequest struct {
	RequestedReviewer struct {
		User struct {
			Name  string
			Login string
		} `graphql:"... on User"`
	}
}

type pullRequest struct {
	Title      githubv4.String
	CreatedAt  githubv4.String
	Number     githubv4.Int
	Repository struct {
		Name  githubv4.String
		Owner struct {
			Login githubv4.String
		}
	}
	Labels struct {
		Nodes []struct {
			Name githubv4.String
		}
	} `graphql:"labels(last: 10)"`
	Author struct {
		Login githubv4.String
	}
	ReviewRequests struct {
		Nodes []reviewRequest
	} `graphql:"reviewRequests(first: 10)"`
}

func (pr *pullRequest) fromContributor() bool {
	for _, label := range pr.Labels.Nodes {
		if label.Name == "contribution" {
			return true
		}
	}
	return false
}

type review struct {
	OccurredAt  githubv4.String
	PullRequest pullRequest
}

type repo struct {
	Owner string
	Name  string
	Issue int
}

type config struct {
	ReportPostTo  *repo    `yaml:"report_post_to"`
	TrackedDays   int      `yaml:"tracked_days"`
	TrackedRepos  []repo   `yaml:"tracked_repos"`
	LoginList     []string `yaml:"login_list"`
	OrgList       []string `yaml:"org_list"`
	OrgID         map[string]githubv4.ID
	PostToIssueID githubv4.ID
}

var cfg config

var client *githubv4.Client

func initOrgID() error {
	var query struct {
		Organization struct {
			ID githubv4.ID
		} `graphql:"organization(login: $login)"`
	}
	cfg.OrgID = make(map[string]githubv4.ID)
	for _, org := range cfg.OrgList {
		err := client.Query(context.Background(), &query, map[string]interface{}{
			"login": githubv4.String(org),
		})
		if err != nil {
			return err
		}
		cfg.OrgID[org] = query.Organization.ID
	}
	return nil
}

func initPostIssueID() error {
	if cfg.ReportPostTo == nil {
		return nil
	}
	var query struct {
		Repository struct {
			Issue struct {
				ID githubv4.ID
			} `graphql:"issue(number: $number)"`
		} `graphql:"repository(owner: $owner, name: $name)"`
	}
	err := client.Query(context.Background(), &query, map[string]interface{}{
		"name":   githubv4.String(cfg.ReportPostTo.Name),
		"owner":  githubv4.String(cfg.ReportPostTo.Owner),
		"number": githubv4.Int(cfg.ReportPostTo.Issue),
	})
	if err != nil {
		return err
	}
	cfg.PostToIssueID = query.Repository.Issue.ID
	return nil
}

func init() {
	at := os.Getenv("GITHUB_TOKEN")
	if len(at) == 0 {
		panic(fmt.Errorf("doesn't have GITHUB_TOKEN"))
	}
	src := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: at},
	)
	httpClient := oauth2.NewClient(context.Background(), src)

	client = githubv4.NewClient(httpClient)

	yml, err := ioutil.ReadFile("./config.yml")
	if err != nil {
		panic(err)
	}
	if err = yaml.Unmarshal(yml, &cfg); err != nil {
		panic(err)
	}
	pp.Println(cfg)
	sort.SliceStable(cfg.LoginList, func(i, j int) bool {
		return cfg.LoginList[i] < cfg.LoginList[j]
	})

	err = initOrgID()
	if err != nil {
		panic(err)
	}
	err = initPostIssueID()
	if err != nil {
		panic(err)
	}
}

type statistics struct {
	login          string
	reviewsByDay   [][]review
	total          int
	reqestedReview []pullRequest
}

var reviewers map[string]*statistics

func newStatistics(login string) *statistics {
	s := statistics{login: login}
	s.reviewsByDay = make([][]review, cfg.TrackedDays)
	for i := 0; i < cfg.TrackedDays; i++ {
		s.reviewsByDay[i] = make([]review, 0)
	}
	return &s
}

func (s *statistics) getReviewRecordByOrg(client *githubv4.Client, userLogin string, organizationID githubv4.ID) error {
	type edge struct {
		Cursor githubv4.String
		Node   review
	}
	var query struct {
		User struct {
			Name                    githubv4.String
			ContributionsCollection struct {
				PullRequestReviewContributions struct {
					Edges []edge
				} `graphql:"pullRequestReviewContributions(first: 100, after: $cursor)"`
			} `graphql:"contributionsCollection(organizationID: $organizationID)"`
		} `graphql:"user(login: $userLogin)"`
	}
	cursor := githubv4.String("")
	now := time.Now()
	contd := true
	for contd {
		contd = false
		err := client.Query(context.Background(), &query, map[string]interface{}{
			"userLogin":      githubv4.String(userLogin),
			"organizationID": organizationID,
			"cursor":         cursor,
		})
		if err != nil {
			return fmt.Errorf("error query: %v", err)
		}
		contribs := query.User.ContributionsCollection.PullRequestReviewContributions.Edges
		for i, contrib := range contribs {
			t, err := time.Parse(time.RFC3339, string(contrib.Node.OccurredAt))
			if err != nil {
				return fmt.Errorf("error parse time: %v", err)
			}
			hours := now.Sub(t).Hours()
			if i == len(contribs)-1 && hours < float64(cfg.TrackedDays*24) {
				cursor = contrib.Cursor
				contd = true
			}
			d := int(hours / 24)
			if d < cfg.TrackedDays && string(contrib.Node.PullRequest.Author.Login) != s.login {
				s.reviewsByDay[d] = append(s.reviewsByDay[d], contrib.Node)
				s.total++
			}
		}
	}
	return nil
}

func (s *statistics) printOpenRequest() string {
	header := []string{"repo", "PR", "C", "lasted"}
	data := make([][]string, len(s.reqestedReview))
	lasted := make([]float64, len(s.reqestedReview))
	for i, rr := range s.reqestedReview {
		createdAt, err := time.Parse(time.RFC3339, string(rr.CreatedAt))
		if err != nil {
			fmt.Println("time parse err", err)
		}
		lasted[i] = time.Now().Sub(createdAt).Hours()
		hours := int(lasted[i])
		days := hours / 24
		lastedStr := ""
		if days > 0 {
			lastedStr += fmt.Sprintf("%dd", days)
		}
		lastedStr += fmt.Sprintf("%dh", hours%24)
		data[i] = []string{
			fmt.Sprintf("%s/%d", rr.Repository.Name, rr.Number),
			fmt.Sprintf("[%s](https://github.com/%s/%s/pull/%d)", rr.Title, rr.Repository.Owner.Login, rr.Repository.Name, rr.Number),
			string(" Y"[map[bool]int{false: 0, true: 1}[rr.fromContributor()]]),
			lastedStr,
		}
	}
	sort.SliceStable(data, func(i, j int) bool {
		return lasted[i] > lasted[j]
	})
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader(header)
	table.SetColWidth(100000) // don't break line
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return buf.String()
}

func (s *statistics) printReviewed() string {
	header := []string{"repo", "PR", "C", "D", "R"}
	data := make([][]string, s.total)
	k := 0
	for i := 0; i < cfg.TrackedDays; i++ {
		for _, r := range s.reviewsByDay[i] {
			rr := r.PullRequest
			createdAt, err := time.Parse(time.RFC3339, string(rr.CreatedAt))
			if err != nil {
				fmt.Println("time parse err", err)
			}
			occuredAt, err := time.Parse(time.RFC3339, string(r.OccurredAt))
			if err != nil {
				fmt.Println("time parsed err", err)
			}
			dur := occuredAt.Sub(createdAt).Hours()
			hours := int(dur)
			days := hours / 24
			lastedStr := ""
			if days > 0 {
				lastedStr += fmt.Sprintf("%dd", days)
			}
			lastedStr += fmt.Sprintf("%dh", hours%24)
			data[k] = []string{
				fmt.Sprintf("%s/%d", rr.Repository.Name, rr.Number),
				fmt.Sprintf("[%s](https://github.com/%s/%s/pull/%d)", rr.Title, rr.Repository.Owner.Login, rr.Repository.Name, rr.Number),
				string(" Y"[map[bool]int{false: 0, true: 1}[rr.fromContributor()]]),
				fmt.Sprintf("%d", i+1),
				lastedStr,
			}
			k++
		}
	}
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	table.SetHeader(header)
	table.SetColWidth(100000) // don't break line
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render()
	return buf.String()
}

func (s *statistics) printDetail() string {

	report := fmt.Sprintf(`
<details> 
  <summary>%s's detailed review report</summary> 

## To Be Reviewed

%s

## Reviewed in Last 7 Days

%s

</details> 

`, s.login, s.printOpenRequest(), s.printReviewed())

	// yesterday's review

	return report
}

func getPullRequestByRepo(mutex *sync.Mutex, owner, name string) error {
	type edge struct {
		Cursor githubv4.String
		Node   pullRequest
	}
	var query struct {
		Repository struct {
			PullRequests struct {
				Edges []edge
			} `graphql:"pullRequests(states: OPEN, first: 100, after: $cursor)"`
		} `graphql:"repository(name: $name, owner: $owner)"`
	}

	cursor := (*githubv4.String)(nil)
	contd := true
	total := 0

	for contd {
		param := map[string]interface{}{
			"name":   githubv4.String(name),
			"owner":  githubv4.String(owner),
			"cursor": cursor,
		}

		err := client.Query(context.Background(), &query, param)
		if err != nil {
			return err
		}

		prs := query.Repository.PullRequests.Edges
		cnt := len(prs)
		total += cnt
		contd = cnt == 100

		mutex.Lock()
		for _, pr := range prs {
			for _, req := range pr.Node.ReviewRequests.Nodes {
				s, ok := reviewers[req.RequestedReviewer.User.Login]
				if ok {
					s.reqestedReview = append(s.reqestedReview, pr.Node)
				}
			}
		}
		mutex.Unlock()
		if cnt != 0 {
			cursor = &prs[cnt-1].Cursor
		}
	}

	fmt.Printf("fetched %d open PRs from %s/%s\n", total, owner, name)
	return nil
}

func getPullRequest() error {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, r := range cfg.TrackedRepos {
		wg.Add(1)
		go func(r repo) {
			err := getPullRequestByRepo(&mutex, r.Owner, r.Name)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}(r)
	}
	wg.Wait()
	return nil
}

func getReviewRecord() {
	reviewers = make(map[string]*statistics)
	var mutex sync.Mutex
	var wg sync.WaitGroup
	for _, user := range cfg.LoginList {
		wg.Add(1)
		go func(user string) {
			s := newStatistics(user)
			for _, ID := range cfg.OrgID {
				err := s.getReviewRecordByOrg(client, user, ID)
				if err != nil {
					fmt.Println(err)
				}
			}
			fmt.Printf("%s analyzed\n", user)
			mutex.Lock()
			reviewers[user] = s
			mutex.Unlock()
			wg.Done()
		}(user)
	}
	wg.Wait()
}

func reportToSlack(webhook, message string) error {
	var data = []byte(fmt.Sprintf("{\"text\": \"%s\"}", message))
	req, err := http.NewRequest("POST", webhook, bytes.NewBuffer(data))
	if err != nil {
		return err
	}
	req.Header.Set("Content-type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}

func reportToIssue(report string) (url string, err error) {
	now := time.Now()
	title := fmt.Sprintf("Detailed Review Report on %d.%d.%d", now.Year(), now.Month(), now.Day())

	var m struct {
		UpdateIssue struct {
			Issue struct {
				Url githubv4.String
			}
		} `graphql:"updateIssue(input: $input)"`
	}
	input := githubv4.UpdateIssueInput{
		ID:    cfg.PostToIssueID,
		Title: githubv4.NewString(githubv4.String(title)),
		Body:  githubv4.NewString(githubv4.String(report)),
	}
	err = client.Mutate(context.Background(), &m, input, nil)
	if err != nil {
		return
	}
	url = string(m.UpdateIssue.Issue.Url)
	return
}

func main() {

	getReviewRecord()
	getPullRequest()

	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	header := []string{"Reviewer", "O(C)"}
	data := make([][]string, len(cfg.LoginList))
	numLogin := len(cfg.LoginList)
	for i := 0; i < numLogin; i++ {
		d := make([]string, 0)
		login := cfg.LoginList[i]
		d = append(d, login)
		s := reviewers[login]
		numOpen := len(s.reqestedReview)
		numCont := 0
		for _, rr := range s.reqestedReview {
			if rr.fromContributor() {
				numCont++
			}
		}
		d = append(d, fmt.Sprintf("%2d (%2d)", numOpen, numCont))
		data[i] = d
	}
	for i := 1; i <= cfg.TrackedDays; i++ {
		header = append(header, fmt.Sprint(i))
		for j, login := range cfg.LoginList {
			data[j] = append(data[j], fmt.Sprintf("%d", len(reviewers[login].reviewsByDay[i-1])))
		}
	}
	header = append(header, "T")
	for i, login := range cfg.LoginList {
		data[i] = append(data[i], fmt.Sprintf("%d", reviewers[login].total))
	}
	table.SetHeader(header)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: false})
	table.SetCenterSeparator("|")
	table.AppendBulk(data)
	table.Render() // Send output

	report := buf.String() + "\n"

	for _, login := range cfg.LoginList {
		report += reviewers[login].printDetail()
	}

	fmt.Println(report)
	if err := ioutil.WriteFile("index.md", []byte(report), 0644); err != nil {
		log.Fatal(err)
	}
	var err error
	if cfg.ReportPostTo != nil {
		_, err = reportToIssue(report)
		if err != nil {
			panic(err)
		}
	}

	webhook := os.Getenv("SLACK_WEBHOOK")
	if len(webhook) != 0 {
		fmt.Println("report to slack")
		err := reportToSlack(webhook, "<!channel>, review reports here\n```\n"+buf.String()+"\n```\nFor a more detailed report, please see https://ichn-hu.github.io/review-promoter/")
		if err != nil {
			panic(err)
		}
	}
}
