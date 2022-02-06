package githubapi

import (
	"encoding/json"
	"testing"
)

var getGistResponseJson = `
[
  {
    "url": "https://api.github.com/gists/8994456e1864e19eb40783a3ef27ea82",
    "forks_url": "https://api.github.com/gists/8994456e1864e19eb40783a3ef27ea82/forks",
    "commits_url": "https://api.github.com/gists/8994456e1864e19eb40783a3ef27ea82/commits",
    "id": "8994456e1864e19eb40783a3ef27ea82",
    "node_id": "MDQ6R2lzdDg5OTQ0NTZlMTg2NGUxOWViNDA3ODNhM2VmMjdlYTgy",
    "git_pull_url": "https://gist.github.com/8994456e1864e19eb40783a3ef27ea82.git",
    "git_push_url": "https://gist.github.com/8994456e1864e19eb40783a3ef27ea82.git",
    "html_url": "https://gist.github.com/8994456e1864e19eb40783a3ef27ea82",
    "files": {
      "ec2-local-port-forward.sh": {
        "filename": "ec2-local-port-forward.sh",
        "type": "application/x-sh",
        "language": "Shell",
        "raw_url": "https://gist.githubusercontent.com/BenSlabbert/8994456e1864e19eb40783a3ef27ea82/raw/243aa787990efbfbc570ffcaea57397ff20e251b/ec2-local-port-forward.sh",
        "size": 121
      }
    },
    "public": true,
    "created_at": "2021-03-01T15:52:44Z",
    "updated_at": "2021-03-01T15:52:51Z",
    "description": "port forward the local port on an ec2 instance to local machine",
    "comments": 0,
    "user": null,
    "comments_url": "https://api.github.com/gists/8994456e1864e19eb40783a3ef27ea82/comments",
    "owner": {
      "login": "BenSlabbert",
      "id": 12390879,
      "node_id": "MDQ6VXNlcjEyMzkwODc5",
      "avatar_url": "https://avatars.githubusercontent.com/u/12390879?v=4",
      "gravatar_id": "",
      "url": "https://api.github.com/users/BenSlabbert",
      "html_url": "https://github.com/BenSlabbert",
      "followers_url": "https://api.github.com/users/BenSlabbert/followers",
      "following_url": "https://api.github.com/users/BenSlabbert/following{/other_user}",
      "gists_url": "https://api.github.com/users/BenSlabbert/gists{/gist_id}",
      "starred_url": "https://api.github.com/users/BenSlabbert/starred{/owner}{/repo}",
      "subscriptions_url": "https://api.github.com/users/BenSlabbert/subscriptions",
      "organizations_url": "https://api.github.com/users/BenSlabbert/orgs",
      "repos_url": "https://api.github.com/users/BenSlabbert/repos",
      "events_url": "https://api.github.com/users/BenSlabbert/events{/privacy}",
      "received_events_url": "https://api.github.com/users/BenSlabbert/received_events",
      "type": "User",
      "site_admin": false
    },
    "truncated": false
  }
]
`

func TestGetGistsResponse(t *testing.T) {
	g := new(GetGistsResponse)
	err := json.Unmarshal([]byte(getGistResponseJson), g)
	if err != nil {
		t.Fatal(err)
	}

	responses := *g

	if len(responses) != 1 {
		t.Fatal()
	}

	if responses[0].ID != "8994456e1864e19eb40783a3ef27ea82" {
		t.Fatal()
	}

	if responses[0].Files["ec2-local-port-forward.sh"].Filename != "ec2-local-port-forward.sh" {
		t.Fatal()
	}
}
