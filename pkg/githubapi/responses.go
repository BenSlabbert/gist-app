package githubapi

import "time"

type FileDetails struct {
	Filename string `json:"filename"`
	Type     string `json:"type"`
	Language string `json:"language"`
	RawUrl   string `json:"raw_url"`
	Size     int    `json:"size"`
}

type GetGistsResponse []struct {
	Comments    int64                  `json:"comments"`
	CommentsURL string                 `json:"comments_url"`
	CommitsURL  string                 `json:"commits_url"`
	CreatedAt   string                 `json:"created_at"`
	Description string                 `json:"description"`
	Files       map[string]FileDetails `json:"files"`
	ForksURL    string                 `json:"forks_url"`
	GitPullURL  string                 `json:"git_pull_url"`
	GitPushURL  string                 `json:"git_push_url"`
	HTMLURL     string                 `json:"html_url"`
	ID          string                 `json:"id"`
	NodeID      string                 `json:"node_id"`
	Owner       struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		NodeID            string `json:"node_id"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"owner"`
	Public    bool        `json:"public"`
	Truncated bool        `json:"truncated"`
	UpdatedAt string      `json:"updated_at"`
	URL       string      `json:"url"`
	User      interface{} `json:"user"`
}

type GistResponse struct {
	Comments    int64         `json:"comments"`
	CommentsURL string        `json:"comments_url"`
	CommitsURL  string        `json:"commits_url"`
	CreatedAt   string        `json:"created_at"`
	Description string        `json:"description"`
	Files       interface{}   `json:"files"`
	Forks       []interface{} `json:"forks"`
	ForksURL    string        `json:"forks_url"`
	GitPullURL  string        `json:"git_pull_url"`
	GitPushURL  string        `json:"git_push_url"`
	History     []struct {
		ChangeStatus struct {
			Additions int64 `json:"additions"`
			Deletions int64 `json:"deletions"`
			Total     int64 `json:"total"`
		} `json:"change_status"`
		CommittedAt string `json:"committed_at"`
		URL         string `json:"url"`
		User        struct {
			AvatarURL         string `json:"avatar_url"`
			EventsURL         string `json:"events_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			GravatarID        string `json:"gravatar_id"`
			HTMLURL           string `json:"html_url"`
			ID                int64  `json:"id"`
			Login             string `json:"login"`
			NodeID            string `json:"node_id"`
			OrganizationsURL  string `json:"organizations_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			ReposURL          string `json:"repos_url"`
			SiteAdmin         bool   `json:"site_admin"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			Type              string `json:"type"`
			URL               string `json:"url"`
		} `json:"user"`
		Version string `json:"version"`
	} `json:"history"`
	HTMLURL string `json:"html_url"`
	ID      string `json:"id"`
	NodeID  string `json:"node_id"`
	Owner   struct {
		AvatarURL         string `json:"avatar_url"`
		EventsURL         string `json:"events_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		GravatarID        string `json:"gravatar_id"`
		HTMLURL           string `json:"html_url"`
		ID                int64  `json:"id"`
		Login             string `json:"login"`
		NodeID            string `json:"node_id"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"owner"`
	Public    bool        `json:"public"`
	Truncated bool        `json:"truncated"`
	UpdatedAt string      `json:"updated_at"`
	URL       string      `json:"url"`
	User      interface{} `json:"user"`
}

type GistCreateResponse struct {
	Url         string    `json:"url"`
	ForksUrl    string    `json:"forks_url"`
	CommitsUrl  string    `json:"commits_url"`
	Id          string    `json:"id"`
	NodeId      string    `json:"node_id"`
	GitPullUrl  string    `json:"git_pull_url"`
	GitPushUrl  string    `json:"git_push_url"`
	HtmlUrl     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	Comments    int       `json:"comments"`
	CommentsUrl string    `json:"comments_url"`
}

type RateLimitResponse struct {
	Resources struct {
		Core struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"core"`
		Search struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"search"`
		Graphql struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"graphql"`
		IntegrationManifest struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"integration_manifest"`
		SourceImport struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"source_import"`
		CodeScanningUpload struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"code_scanning_upload"`
		ActionsRunnerRegistration struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"actions_runner_registration"`
		Scim struct {
			Limit     int `json:"limit"`
			Used      int `json:"used"`
			Remaining int `json:"remaining"`
			Reset     int `json:"reset"`
		} `json:"scim"`
	} `json:"resources"`
	Rate struct {
		Limit     int `json:"limit"`
		Used      int `json:"used"`
		Remaining int `json:"remaining"`
		Reset     int `json:"reset"`
	} `json:"rate"`
}
