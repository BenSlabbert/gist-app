package githubapi

import (
	"encoding/json"
	"fmt"
	"github.com/BenSlabbert/gist-app/pkg/env"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

var api *Api

func TestMain(m *testing.M) {
	a, err := NewApi(os.Getenv(env.GithubApiUsername), os.Getenv(env.GithubApiToken))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	api = a
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

func TestApi_PutGist(t *testing.T) {
	gist, err := api.PutGist("file.txt", "content", "description", true)
	if err != nil {
		t.Fatal(err)
	}

	if gist.Id == "" {
		t.Fatal()
	}

	// clean up
	err = api.DeleteGist(gist.Id)
	if err != nil {
		t.Fatal(err)
	}
}

func TestFiles_MarshalJSON(t *testing.T) {
	x := &gistCreate{
		Description: "new file from app",
		Public:      true,
		Files: CustomFilesMarshaller{
			filename: "new_filename.sh",
			files: Files{
				Name: Name{
					Content: "file content",
				},
			},
		},
	}

	bytes, err := json.Marshal(x)
	if err != nil {
		t.Fatal(err)
	}

	s := string(bytes)
	if strings.Contains(s, "replace_me") {
		t.Fatal()
	}
	if !strings.Contains(s, "new_filename.sh") {
		t.Fatal()
	}
}

func TestApi_GetRateLimit(t *testing.T) {
	// todo implement a rate limiter
	//  use this endpoint when starting to see how many requests we can make
	//  then limit the requests accordingly

	rl, err := api.GetRateLimit()
	if err != nil {
		t.Fatal(err)
	}

	if rl == nil {
		t.Fatal()
	}

	i, err := strconv.ParseInt(fmt.Sprintf("%d", rl.Resources.Core.Reset), 10, 64)
	if err != nil {
		panic(err)
	}

	tm := time.Unix(i, 0)
	fmt.Println(tm)
	s := tm.String()
	fmt.Println(s)
	if tm.Before(time.Now()) {
		t.Fatal()
	}
}
