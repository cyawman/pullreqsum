package pullreqsum

import (
	"time"

	"github.com/google/go-github/github"
)

type LatestPullRequests struct {
	CutoffDate time.Time
}

func (lpr *LatestPullRequests) Filter(pullRequests []*github.PullRequest) []*github.PullRequest {
	var latestPullRequests []*github.PullRequest
	for _, pr := range pullRequests {
		if pr.CreatedAt.After(lpr.CutoffDate) {
			latestPullRequests = append(latestPullRequests, pr)
		}
	}
	return latestPullRequests
}
