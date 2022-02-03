package githubapi

import (
	"os"
	"testing"
)

func TestApi_FetchPrivateGists(t *testing.T) {
	api := NewApi("BenSlabbert", os.Getenv("GITHUB_API_TOKEN"))
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
	api := NewApi("BenSlabbert", os.Getenv("GITHUB_API_TOKEN"))
	gist, err := api.FetchPrivateGist(gistId)

	if err != nil {
		t.Fatal(err)
	}

	if gist.ID != gistId {
		t.Fatal("incorrect gist id")
	}
}
