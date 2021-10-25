package pullreqsum

import (
	"testing"

	"github.com/google/go-github/github"
)

type mockObserver struct{}

func (m *mockObserver) Update(s Subject) {
	pr := s.GetPullRequests()
	if pr == nil {
		panic("No Pull Requests found")
	}
}

func TestSummarySubjectNotifier(t *testing.T) {
	pr := github.PullRequest{}
	pullRequests := make([]*github.PullRequest, 1)
	pullRequests[0] = &pr

	o := mockObserver{}
	s := SummarySubject{PullRequests: pullRequests}
	s.Attach(&o)
	s.Notify()
}
