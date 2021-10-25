package pullreqsum

import "testing"

func TestParseGithubRepositoryPath(t *testing.T) {
	path := "test/example"

	owner, repo := parseGithubRepositoryPath(path)
	if owner != "test" {
		t.Fatalf("Got %s expected %s", owner, "test")
	}
	if repo != "example" {
		t.Fatalf("Got %s expected %s", repo, "example")
	}
}
