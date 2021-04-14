package doorkeeper

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
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
	body, rateLimit, err := c.get(path, params.toValues())

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	var res []rawGetEventResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	events := []*Event{}

	for _, e := range res {
		event, err := e.Event.toEvent()

		if err != nil {
			return nil, nil, errors.WithStack(err)
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
	values := optionsToValues(options)

	body, rateLimit, err := c.get(fmt.Sprintf("/events/%d", eventID), values)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	var res rawGetEventResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	event, err := res.Event.toEvent()

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return event, rateLimit, nil
}

type rawGetGroupResponse struct {
	Group Group `json:"group"`
}

// GetGroup returns a specific group
func (c *Client) GetGroup(groupName string, options ...OptionFunc) (*Group, *RateLimit, error) {
	values := optionsToValues(options)

	body, rateLimit, err := c.get("/groups/"+groupName, values)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	var res rawGetGroupResponse
	err = json.Unmarshal(body, &res)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return &res.Group, rateLimit, nil
}

func (c *Client) get(path string, values url.Values) ([]byte, *RateLimit, error) {
	url, err := c.buildURL(path, values)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, nil, err
	}

	req.Header.Add("Authorization", "Bearer "+c.accessToken)
	req.Header.Add("User-Agent", c.UserAgent)

	resp, err := c.client.Do(req)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	xRateLimit := resp.Header.Get("X-Ratelimit")

	if xRateLimit == "" {
		var rateLimit RateLimit
		return body, &rateLimit, nil
	}

	rawRateLimit, err := newRawRateLimitFromJSON(xRateLimit)

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	rateLimit, err := rawRateLimit.toRateLimit()

	if err != nil {
		return nil, nil, errors.WithStack(err)
	}

	return body, rateLimit, nil
}

func (c *Client) buildURL(path string, values url.Values) (string, error) {
	u, err := url.Parse(baseURL + path)

	if err != nil {
		return "", err
	}

	u.RawQuery = values.Encode()

	return u.String(), nil
}
