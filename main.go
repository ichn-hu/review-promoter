package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/goccy/go-yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type user struct {
	Name  string
	Login string
}
type request struct {
	RequestedReviewer struct {
		User user `graphql:"... on User"`
	}
}
type review struct {
	OccurredAt  githubv4.String
	PullRequest struct {
		Number githubv4.Int
		URL    githubv4.String
		Title  githubv4.String
		Author struct {
			Login githubv4.String
		}
		ReviewRequests struct {
			Nodes []request
		} `graphql:"reviewRequests(first: 10)"`
	}
}

type config struct {
	TrackDays int      `yaml:"track_days"`
	LoginList []string `yaml:"login_list"`
	OrgList   []string `yaml:"org_list"`
	OrgID     map[string]githubv4.ID
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
	fmt.Println("parsed configuration: ", cfg)

	err = initOrgID()
	if err != nil {
		panic(err)
	}
}

type Statistics struct {
	login        string
	reviewsByDay [][]review
}

func NewStatistics(login string) *Statistics {
	s := Statistics{login: login}
	s.reviewsByDay = make([][]review, cfg.TrackDays)
	for i := 0; i < cfg.TrackDays; i++ {
		s.reviewsByDay[i] = make([]review, 0)
	}
	return &s
}

func (s *Statistics) getReviewRecordByOrg(client *githubv4.Client, userLogin string, organizationID githubv4.ID) error {
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
			if i == len(contribs)-1 && hours < float64(cfg.TrackDays*24) {
				cursor = contrib.Cursor
				contd = true
			}
			d := int(hours / 24)
			if d < cfg.TrackDays {
				s.reviewsByDay[d] = append(s.reviewsByDay[d], contrib.Node)
			}
		}
	}
	return nil
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

func main() {
	data := make([][]string, len(cfg.LoginList))
	for i, user := range cfg.LoginList {
		data[i] = make([]string, cfg.TrackDays+1)
		data[i][0] = user
		s := NewStatistics(user)
		for _, ID := range cfg.OrgID {
			err := s.getReviewRecordByOrg(client, user, ID)
			if err != nil {
				fmt.Println(err)
			}
		}
		fmt.Printf("%s analyzed\n", user)
		for j := 0; j < cfg.TrackDays; j++ {
			data[i][j+1] = fmt.Sprint(len(s.reviewsByDay[j]))
		}
	}
	var buf bytes.Buffer
	table := tablewriter.NewWriter(&buf)
	header := []string{"Contributor"}
	for i := 1; i <= cfg.TrackDays; i++ {
		var h string
		if i == 1 {
			h = "Yesterday"
		} else {
			h = fmt.Sprintf("%d days ago", i)
		}
		header = append(header, h)
	}
	table.SetHeader(header)
	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output

	fmt.Println(buf.String())
	webhook := os.Getenv("SLACK_WEBHOOK")
	if len(webhook) != 0 {
		fmt.Println("report to slack")
		err := reportToSlack(webhook, "<!channel>, review reports here\n```\n"+buf.String()+"\n```")
		if err != nil {
			panic(err)
		}
	}
}
