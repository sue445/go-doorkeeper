package doorkeeper

import (
	"github.com/cockroachdb/errors"
	"time"
)

// A Event represents doorkeeper event
type Event struct {
	Title        string
	ID           int
	StartsAt     time.Time
	EndsAt       time.Time
	VenueName    string
	Address      string
	Lat          *float64
	Long         *float64
	PublishedAt  time.Time
	UpdatedAt    time.Time
	Group        int
	Description  string
	PublicURL    string
	Participants int
	Waitlisted   int
	TicketLimit  int
}

type rawEvent struct {
	Title        string   `json:"title"`
	ID           int      `json:"id"`
	StartsAt     string   `json:"starts_at"`
	EndsAt       string   `json:"ends_at"`
	VenueName    string   `json:"venue_name"`
	Address      string   `json:"address"`
	Lat          *float64 `json:"lat"`
	Long         *float64 `json:"long"`
	PublishedAt  string   `json:"published_at"`
	UpdatedAt    string   `json:"updated_at"`
	Group        int      `json:"group"`
	Description  string   `json:"description"`
	PublicURL    string   `json:"public_url"`
	Participants int      `json:"participants"`
	Waitlisted   int      `json:"waitlisted"`
	TicketLimit  int      `json:"ticket_limit"`
}

func (e *rawEvent) toEvent() (*Event, error) {
	startsAt, err := time.Parse(timeFormatMs, e.StartsAt)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	endsAt, err := time.Parse(timeFormatMs, e.EndsAt)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	publishedAt, err := time.Parse(timeFormatMs, e.PublishedAt)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	updatedAt, err := time.Parse(timeFormatMs, e.UpdatedAt)

	if err != nil {
		return nil, errors.WithStack(err)
	}

	event := &Event{
		Title:        e.Title,
		ID:           e.ID,
		StartsAt:     startsAt,
		EndsAt:       endsAt,
		VenueName:    e.VenueName,
		Address:      e.Address,
		PublishedAt:  publishedAt,
		UpdatedAt:    updatedAt,
		Group:        e.Group,
		Description:  e.Description,
		PublicURL:    e.PublicURL,
		Participants: e.Participants,
		Waitlisted:   e.Waitlisted,
		TicketLimit:  e.TicketLimit,
		Lat:          e.Lat,
		Long:         e.Long,
	}

	return event, nil
}
