package doorkeeper

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	baseURL    = "https://api.doorkeeper.jp"
	timeFormat = "2006-01-02T15:04:05Z"
)

// A Client manages communication with the Doorkeeper API
type Client struct {
	accessToken string
	client      *http.Client
	UserAgent   string
}

// NewClient returns a new API Client instance
func NewClient(accessToken string) *Client {
	userAgent := fmt.Sprintf("go-doorkeeper/%s (+https://github.com/sue445/go-doorkeeper)", Version)

	return &Client{accessToken: accessToken, UserAgent: userAgent, client: &http.Client{}}
}

// GetGroup returns a specific group
func (c *Client) GetGroup(groupName string, options ...OptionFunc) (*Group, *RateLimit, error) {
	params := optionsToParams(options)

	body, rateLimit, err := c.get("/groups/"+groupName, params)

	if err != nil {
		return nil, nil, err
	}

	var rawGroup rawGroup
	err = json.Unmarshal(body, &rawGroup)

	if err != nil {
		return nil, nil, err
	}

	return &rawGroup.Group, rateLimit, nil
}

func (c *Client) get(path string, params map[string]string) ([]byte, *RateLimit, error) {
	req, err := http.NewRequest("GET", c.buildURL(path, params), nil)

	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("User-Agent", c.UserAgent)

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	xRateLimit := resp.Header.Get("X-Ratelimit")
	rawRateLimit, err := newRawRateLimitFromJSON(xRateLimit)

	if err != nil {
		return nil, nil, err
	}

	rateLimit, err := rawRateLimit.toRateLimit()

	if err != nil {
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, err
	}

	return body, rateLimit, nil
}

func (c *Client) buildURL(path string, params map[string]string) string {
	if len(params) == 0 {
		return baseURL + path
	}

	query := make([]string, len(params))
	for k := range params {
		query = append(query, k+"="+params[k])
	}
	return baseURL + path + "?" + strings.Join(query, "&")
}
