package githubapi

import (
	"github.com/BenSlabbert/gist-app/pkg/env"
	"os"
	"testing"
)

var api *Api

func TestMain(m *testing.M) {
	api = NewApi("BenSlabbert", os.Getenv(env.GithubApiToken))
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestApi_FetchPrivateGists(t *testing.T) {
	gists, err := api.FetchPrivateGists()

	if err != nil {
		t.Fatal(err)
	}

	if !(len(gists) > 0) {
		t.Fatal("gists should not be empty")
	}
}

func TestApi_FetchPrivateGist(t *testing.T) {
	gistId := "cd67c07a821d74bec9e5e10bc4b92e2d"
	gist, err := api.FetchPrivateGist(gistId)

	if err != nil {
		t.Fatal(err)
	}

	if gist.ID != gistId {
		t.Fatal("incorrect gist id")
	}
}
