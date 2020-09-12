# review-promoter

Review promoter is a github action that could gather the review activity of contributors, which is meant to friendly remind reviewers to routinely review assigned pull requests.

To use this action, you need to fork this repository, and update the configuration file `config.yml`, where

`report_post_to` is a github repository in which an issue will be updated at every run of the action, with detailed review activity and review requests.

An example is

```
report_post_to:
  owner: ichn-hu
  name: review-promoter
  issue: 1
```

where issue 1 of repository `ichn-hu/review-promoter` will be updated.

Drop this field if you don't want to report the details to a github issue (which will automatically link the pull requests mentioned if the repository posted to is public).

`tracked_days` is the number of days tracked back from the point of executing the action, default is 7, meaning the last week.

The printed table has several fields, where O stands for number of open PRs that is waiting for review, C means number of PRs out of them are from contributors, then 1 to 7 means number of review contributions made in the last 7 days, and T means total review contribution in the last 7 days. 

`login_list` is the login name of github users that you want to track.

`org_list` is the organizations that the reviewers contribute to, for each of the review action happend under the organizations in the list, it will be recorded.

`tracked_repos` contains the repositories that are tracked, if a tracked user has a review request in the tracked repositories, it will be reported.

