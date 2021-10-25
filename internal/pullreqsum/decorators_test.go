package pullreqsum

import (
	"testing"
	"time"

	"github.com/google/go-github/github"
)

func TestLatestPullRequestsFilter(t *testing.T) {
	cutoffDate := time.Now().AddDate(0, 0, -7)

	lastMonth := time.Now().AddDate(0, -1, 0)
	lastWeek := time.Now().AddDate(0, 0, -6)
	yesterday := time.Now().AddDate(0, 0, -1)

	lastMonthPR := github.PullRequest{CreatedAt: &lastMonth}
	lastWeekPR := github.PullRequest{CreatedAt: &lastWeek}
	yesterdayPR := github.PullRequest{CreatedAt: &yesterday}

	var pullRequests []*github.PullRequest
	pullRequests = append(pullRequests, &lastMonthPR, &lastWeekPR, &yesterdayPR)

	decorator := LatestPullRequests{CutoffDate: cutoffDate}
	latestPRs := decorator.Filter(pullRequests)

	if len(latestPRs) != 2 {
		t.Fatalf("Got %d pull requests expected %d", len(latestPRs), 2)
	}

}
