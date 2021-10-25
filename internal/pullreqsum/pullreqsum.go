package pullreqsum

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/github"
)

type Observer interface {
	Update(s Subject)
}

type Subject interface {
	GetPullRequests() []*github.PullRequest
	Attach(obs Observer)
	Notify()
}

type PullRequestDecorator interface {
	Filter(pullRequests []*github.PullRequest) []*github.PullRequest
}

type Config struct {
	GithubRepository  string
	MessageSender     string
	MessageRecipients []string
	MessageSubject    string
}

func Run(config Config) {
	owner, repo := parseGithubRepositoryPath(config.GithubRepository)

	allPullRequests, err := getAllPullRequestsFromGithub(owner, repo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	lastWeek := time.Now().AddDate(0, 0, -7)
	lpr := LatestPullRequests{CutoffDate: lastWeek}
	latestPullRequests := lpr.Filter(allPullRequests)

	consolePrinter := ConsolePrinter{
		MessageSender:     config.MessageSender,
		MessageRecipients: config.MessageRecipients,
		MessageSubject:    config.MessageSubject,
	}

	summary := SummarySubject{PullRequests: latestPullRequests}
	summary.Attach(&consolePrinter)
	summary.Notify()
}

func parseGithubRepositoryPath(s string) (owner string, repo string) {
	gh := strings.Split(s, "/")
	return gh[0], gh[1]
}

func getAllPullRequestsFromGithub(owner string, repo string) ([]*github.PullRequest, error) {
	ctx := context.Background()
	opts := github.PullRequestListOptions{State: "all"}

	client := github.NewClient(nil)
	allPullRequests, _, err := client.PullRequests.List(ctx, owner, repo, &opts)
	if err != nil {
		return nil, err
	}
	return allPullRequests, nil
}
