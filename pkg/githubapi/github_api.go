package githubapi

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/BenSlabbert/gist-app/pkg/util"
	"go.uber.org/ratelimit"
	"io/ioutil"
	"log"
	"math"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"
)

const GistApiBaseUrl = "https://api.github.com/gists"
const RateLimitApiUrl = "https://api.github.com/rate_limit"
const GithubJsonFormat = "application/vnd.github.v3+json"

type Api struct {
	username string
	token    string

	httpClientMtx sync.Mutex
	httpClient    *http.Client

	rateLimitInfo *rateLimitInfo
}

type rateLimitInfo struct {
	limit     int
	remaining int
	reset     time.Time
}

func (rli *rateLimitInfo) update(header http.Header) {
	limit, err := strconv.ParseInt(header.Get("X-Ratelimit-Limit"), 10, 32)
	if err == nil {
		rli.limit = int(limit)
	}

	remaining, err := strconv.ParseInt(header.Get("X-Ratelimit-Remaining"), 10, 32)
	if err == nil {
		rli.remaining = int(remaining)
	}

	tm, err := util.UnixTimestampStringToTime(header.Get("X-Ratelimit-Reset"))
	if err == nil {
		rli.reset = tm
	}
}

func NewApi(username string, token string) (*Api, error) {
	if username == "" || token == "" {
		return nil, fmt.Errorf("username and token are required values")
	}

	a := &Api{
		username:      username,
		token:         token,
		httpClientMtx: sync.Mutex{},
		rateLimitInfo: &rateLimitInfo{},
	}

	err := a.init()
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (api *Api) init() error {
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

	// initial update of rate limiter
	_, err := api.GetRateLimit()
	if err != nil {
		return err
	}

	return nil
}

func (api *Api) FetchPrivateGists() (GetGistsResponse, error) {
	req, err := api.buildFetchPrivateGistsRequest()
	if err != nil {
		return nil, err
	}

	resp, err := api.executeRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected http status %d but was %d", http.StatusOK, resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)
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
	req, err := api.buildFetchPrivateGistRequest(id)
	if err != nil {
		return nil, err
	}

	resp, err := api.executeRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("expected http status %d but was %d", http.StatusOK, resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)
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

func (api *Api) PutGist(fileName, content, description string, public bool) (*GistCreateResponse, error) {
	req, err := api.buildPutGistRequest(description, fileName, content, public)
	if err != nil {
		return nil, err
	}

	resp, err := api.executeRequest(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("expected http status %d but was %d", http.StatusCreated, resp.StatusCode)
	}

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	gist := new(GistCreateResponse)
	err = json.Unmarshal(all, gist)
	if err != nil {
		return nil, err
	}

	return gist, nil
}

func (api *Api) DeleteGist(id string) error {
	req, err := api.buildDeleteGistRequest(id)
	if err != nil {
		return err
	}

	resp, err := api.executeRequest(req)
	if err != nil {
		return err
	}

	if resp.StatusCode == http.StatusNoContent || resp.StatusCode == http.StatusNotModified {
		return nil
	}

	return fmt.Errorf("expected http status %d or %d but was %d", http.StatusNoContent, http.StatusNotModified, resp.StatusCode)
}

func (api *Api) GetRateLimit() (*RateLimitResponse, error) {
	req, err := api.buildGetRateLimitRequest()
	if err != nil {
		return nil, err
	}

	// thi request is excluded from rate limiting:
	// https://docs.github.com/en/rest/overview/resources-in-the-rest-api#checking-your-rate-limit-status
	resp, err := api.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	api.rateLimitInfo.update(resp.Header)

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	gist := new(RateLimitResponse)
	err = json.Unmarshal(all, gist)
	if err != nil {
		return nil, err
	}

	return gist, nil
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (api *Api) addHeaders(req *http.Request) {
	req.Header.Add("Accept", GithubJsonFormat)
	req.Header.Add("Authorization", "Basic "+basicAuth(api.username, api.token))
}

func (api *Api) buildDeleteGistRequest(id string) (*http.Request, error) {
	if id == "" {
		return nil, fmt.Errorf("must provide a gist id")
	}

	req, err := http.NewRequest(http.MethodDelete, GistApiBaseUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	api.addHeaders(req)

	return req, nil
}

func (api *Api) buildFetchPrivateGistRequest(id string) (*http.Request, error) {
	if id == "" {
		return nil, fmt.Errorf("must provide a gist id")
	}

	req, err := http.NewRequest(http.MethodGet, GistApiBaseUrl+"/"+id, nil)
	if err != nil {
		return nil, err
	}

	api.addHeaders(req)

	return req, nil
}

func (api *Api) buildGetRateLimitRequest() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, RateLimitApiUrl, nil)
	if err != nil {
		return nil, err
	}

	api.addHeaders(req)

	return req, nil
}

func (api *Api) buildFetchPrivateGistsRequest() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, GistApiBaseUrl, nil)
	if err != nil {
		return nil, err
	}

	api.addHeaders(req)

	return req, nil
}

func (api *Api) buildPutGistRequest(description, fileName, content string, public bool) (*http.Request, error) {
	reqBody := &gistCreate{
		Description: description,
		Public:      public,
		Files: CustomFilesMarshaller{
			filename: fileName,
			files: Files{
				Name: Name{
					Content: content,
				},
			},
		},
	}

	marshal, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}

	buffer := bytes.NewBuffer(marshal)
	req, err := http.NewRequest(http.MethodPost, GistApiBaseUrl, buffer)
	if err != nil {
		return nil, err
	}

	api.addHeaders(req)

	return req, nil
}

func (api *Api) executeRequest(req *http.Request) (*http.Response, error) {
	api.httpClientMtx.Lock()
	defer api.httpClientMtx.Unlock()

	seconds := time.Until(api.rateLimitInfo.reset).Seconds()
	reqPerSeconds := float64(api.rateLimitInfo.remaining) / seconds
	log.Printf("requests per second available: %.2f", reqPerSeconds)
	reqPerSeconds = math.Round(reqPerSeconds)

	// slightly fewer requests so we don't run into trouble
	if reqPerSeconds > 1 {
		reqPerSeconds--
	}

	rl := ratelimit.New(int(reqPerSeconds), ratelimit.WithSlack(5))

	log.Println("waiting for token from rate limiter")
	start := time.Now()
	_ = rl.Take()
	log.Printf("got token from rate limiter, executing request (wait time: %s)", time.Since(start))

	resp, err := api.httpClient.Do(req)
	api.rateLimitInfo.update(resp.Header)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

type Name struct {
	Content string `json:"content"`
}

type Files struct {
	Name Name `json:"replace_me"`
}

type CustomFilesMarshaller struct {
	filename string
	files    Files
}

// MarshalJSON https://stackoverflow.com/a/42547226/4841710
// implements custom JSON formatting, as we want to change the json tag of the filename dynamically
func (u CustomFilesMarshaller) MarshalJSON() ([]byte, error) {
	m, err := json.Marshal(u.files)
	if err != nil {
		return nil, err
	}

	var a interface{}
	err = json.Unmarshal(m, &a)
	if err != nil {
		return nil, err
	}
	b := a.(map[string]interface{})

	b[u.filename] = b["replace_me"]
	delete(b, "replace_me")

	return json.Marshal(b)
}

type gistCreate struct {
	Description string                `json:"description"`
	Public      bool                  `json:"public"`
	Files       CustomFilesMarshaller `json:"files"`
}
