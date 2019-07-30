package doorkeeper

import (
	"strconv"
	"time"
)

// Sort represents sort type of events API
type Sort int

const (
	// Nothing represents no sort
	Nothing Sort = iota

	// PublishedAt represents sort by published_at
	PublishedAt

	// StartsAt represents sort by starts_at
	StartsAt

	// UpdatedAt represents sort by updated_at
	UpdatedAt
)

func (s Sort) String() string {
	switch s {
	case PublishedAt:
		return "published_at"
	case StartsAt:
		return "starts_at"
	case UpdatedAt:
		return "updated_at"
	case Nothing:
		return ""
	default:
		return ""
	}
}

// GetEventsParams manages params of GetEvents
type GetEventsParams struct {
	Page     int
	Locale   string
	Sort     Sort
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

	if p.Sort.String() != "" {
		params["sort"] = p.Sort.String()
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
