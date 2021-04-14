package doorkeeper

import (
	"encoding/json"
	"github.com/pkg/errors"
	"time"
)

// A RateLimit represents API Rate Limit
type RateLimit struct {
	Name      string
	Period    int
	Limit     int
	Remaining int
	Until     time.Time
}

type rawRateLimit struct {
	Name      string `json:"name"`
	Period    int    `json:"period"`
	Limit     int    `json:"limit"`
	Remaining int    `json:"remaining"`
	Until     string `json:"until"`
}

func newRawRateLimitFromJSON(jsonText string) (*rawRateLimit, error) {
	jsonBlob := []byte(jsonText)

	var data rawRateLimit
	err := json.Unmarshal(jsonBlob, &data)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &data, nil
}

func (r *rawRateLimit) toRateLimit() (*RateLimit, error) {
	until, err := time.Parse(timeFormat, r.Until)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &RateLimit{
		Name:      r.Name,
		Period:    r.Period,
		Limit:     r.Limit,
		Remaining: r.Remaining,
		Until:     until,
	}, nil
}
