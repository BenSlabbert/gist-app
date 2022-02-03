package githubapi

// https://json2struct.mervine.net/

type GetGistsResponse []struct {
	Comments    int64       `json:"comments"`
	CommentsURL string      `json:"comments_url"`
	CommitsURL  string      `json:"commits_url"`
	CreatedAt   string      `json:"created_at"`
	Description string      `json:"description"`
	Files       interface{} `json:"files"`
	ForksURL    string      `json:"forks_url"`
	GitPullURL  string      `json:"git_pull_url"`
	GitPushURL  string      `json:"git_push_url"`
	HTMLURL     string      `json:"html_url"`
	ID          string      `json:"id"`
	NodeID      string      `json:"node_id"`
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
