package doorkeeper

import (
	"net/url"
	"strconv"
	"time"
)

// GetEventsParams manages params of GetEvents
type GetEventsParams struct {
	Page     int
	Locale   string
	Sort     SortEnum
	Since    *time.Time
	Until    *time.Time
	Query    string
	Callback string
}

func (p *GetEventsParams) toValues() url.Values {
	v := url.Values{}

	if p.Page > 0 {
		v.Set("page", strconv.Itoa(p.Page))
	}

	if p.Locale != "" {
		v.Set("locale", p.Locale)
	}

	if p.Sort.GetValue() != "" {
		v.Set("sort", p.Sort.GetValue())
	}

	if p.Since != nil {
		v.Set("since", p.Since.Format(dateFormat))
	}

	if p.Until != nil {
		v.Set("until", p.Until.Format(dateFormat))
	}

	if p.Query != "" {
		v.Set("q", p.Query)
	}

	if p.Callback != "" {
		v.Set("callback", p.Callback)
	}

	return v
}
