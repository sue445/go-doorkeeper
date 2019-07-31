package doorkeeper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strings"
)

const (
	baseURL      = "https://api.doorkeeper.jp"
	timeFormat   = "2006-01-02T15:04:05Z"
	timeFormatMs = "2006-01-02T15:04:05.999Z"
	dateFormat   = "2006-01-02"
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

// GetEvents returns events
func (c *Client) GetEvents(params *GetEventsParams) ([]*Event, *RateLimit, error) {
	return c.getEvents("/events", params)
}

// GetGroupEvents returns group events
func (c *Client) GetGroupEvents(groupName string, params *GetEventsParams) ([]*Event, *RateLimit, error) {
	return c.getEvents(fmt.Sprintf("/groups/%s/events", groupName), params)
}

func (c *Client) getEvents(path string, params *GetEventsParams) ([]*Event, *RateLimit, error) {
	body, rateLimit, err := c.get(path, params.toMap())

	if err != nil {
		return nil, nil, err
	}

	var res []rawGetEventResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, nil, err
	}

	events := []*Event{}

	for _, e := range res {
		event, err := e.Event.toEvent()

		if err != nil {
			return nil, nil, err
		}

		events = append(events, event)
	}

	return events, rateLimit, nil
}

type rawGetEventResponse struct {
	Event rawEvent `json:"event"`
}

// GetEvent returns a specific event
func (c *Client) GetEvent(eventID int, options ...OptionFunc) (*Event, *RateLimit, error) {
	params := optionsToParams(options)

	body, rateLimit, err := c.get(fmt.Sprintf("/events/%d", eventID), params)

	if err != nil {
		return nil, nil, err
	}

	var res rawGetEventResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, nil, err
	}

	event, err := res.Event.toEvent()

	if err != nil {
		return nil, nil, err
	}

	return event, rateLimit, nil
}

type rawGetGroupResponse struct {
	Group Group `json:"group"`
}

// GetGroup returns a specific group
func (c *Client) GetGroup(groupName string, options ...OptionFunc) (*Group, *RateLimit, error) {
	params := optionsToParams(options)

	body, rateLimit, err := c.get("/groups/"+groupName, params)

	if err != nil {
		return nil, nil, err
	}

	var res rawGetGroupResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, nil, err
	}

	return &res.Group, rateLimit, nil
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

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, nil, errors.New(resp.Status)
	}

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

	var query []string
	for k := range params {
		query = append(query, k+"="+params[k])
	}
	sort.Strings(query)
	return baseURL + path + "?" + strings.Join(query, "&")
}
