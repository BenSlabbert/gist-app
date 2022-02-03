package githubapi

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

// basic auth: curl -u username:token https://api.github.com/user

const GistApiBaseUrl = "https://api.github.com/gists"
const JsonFormat = "application/vnd.github.v3+json"

type Api struct {
	username string
	token    string

	httpClient *http.Client
}

func NewApi(username string, token string) *Api {
	a := &Api{username: username, token: token}
	a.init()
	return a
}

func (api *Api) init() {
	var netTransport = &http.Transport{
		DialContext: (&net.Dialer{
			Timeout:   5 * time.Second,
			KeepAlive: 5 * time.Second,
		}).DialContext,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	api.httpClient = &http.Client{
		Transport: netTransport,
		Timeout:   time.Second * 10,
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (api *Api) FetchPrivateGists() (GetGistsResponse, error) {
	req, err := http.NewRequest(http.MethodGet, GistApiBaseUrl, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(api.username, api.token))
	response, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected http status %d but was %d", http.StatusOK, response.StatusCode)
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	obj := new(GetGistsResponse)
	err = json.Unmarshal(all, obj)
	if err != nil {
		return nil, err
	}

	return *obj, nil
}

func (api *Api) FetchPrivateGist(id string) (*GistResponse, error) {
	req, err := http.NewRequest(http.MethodGet, GistApiBaseUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Basic "+basicAuth(api.username, api.token))
	response, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected http status %d but was %d", http.StatusOK, response.StatusCode)
	}

	all, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	gist := new(GistResponse)
	err = json.Unmarshal(all, gist)
	if err != nil {
		return nil, err
	}

	return gist, nil
}
