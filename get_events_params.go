package doorkeeper

import (
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

func (p *GetEventsParams) toMap() map[string]string {
	params := map[string]string{}

	if p.Page > 0 {
		params["page"] = strconv.Itoa(p.Page)
	}

	if p.Locale != "" {
		params["locale"] = p.Locale
	}

	if p.Sort.GetValue() != "" {
		params["sort"] = p.Sort.GetValue()
	}

	if p.Since != nil {
		params["since"] = p.Since.Format(dateFormat)
	}

	if p.Until != nil {
		params["until"] = p.Until.Format(dateFormat)
	}

	if p.Query != "" {
		params["q"] = p.Query
	}

	if p.Callback != "" {
		params["callback"] = p.Callback
	}

	return params
}
