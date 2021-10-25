package pullreqsum

import "github.com/google/go-github/github"

type SummarySubject struct {
	Observers    []Observer
	PullRequests []*github.PullRequest
}

func (s *SummarySubject) GetPullRequests() []*github.PullRequest {
	return s.PullRequests
}

func (s *SummarySubject) Attach(obs Observer) {
	s.Observers = append(s.Observers, obs)
}

func (s *SummarySubject) Notify() {
	for _, o := range s.Observers {
		o.Update(s)
	}
}
