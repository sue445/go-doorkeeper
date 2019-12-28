package doorkeeper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_rawEvent_toEvent(t *testing.T) {
	type fields struct {
		Title        string
		ID           int
		StartsAt     string
		EndsAt       string
		VenueName    string
		Address      string
		Lat          *float64
		Long         *float64
		PublishedAt  string
		UpdatedAt    string
		Group        int
		Description  string
		PublicURL    string
		Participants int
		Waitlisted   int
		TicketLimit  int
	}
	tests := []struct {
		name    string
		fields  fields
		want    *Event
		wantErr bool
	}{
		{
			name: "normal",
			fields: fields{
				Title:        "900K records per second with Ruby, Java, and JRuby",
				ID:           28319,
				StartsAt:     "2015-08-13T10:00:00.000Z",
				EndsAt:       "2015-08-13T13:00:00.000Z",
				VenueName:    "VOYAGE GROUP",
				Address:      "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
				Lat:          fp(35.6553195),
				Long:         fp(139.6937795),
				PublishedAt:  "2015-07-13T23:48:29.463Z",
				UpdatedAt:    "2018-05-11T00:07:44.270Z",
				Group:        24,
				Description:  "<h2>Agenda</h2>\n\n<h3><small>19:00 〜 19:45</small> Doors open</h3>\n\n<p>Grab a drink and catch up with your fellow Rubyists.</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">Hiroshi Nakamura (@nahi)</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">Treasure Data</a> is cloud analytics infrastructure. As of the beginning of July 2015, it stores 16 trillion records of customers data, and grows by 900,000 records per second. For handling this data, we are leveraging Ruby, Java, JRuby and many OSS components. In this presentation I'll introduce the system architecture and will discuss how we utilize these languages and OSS components to create a robust, performant, and flexible infrastructure.</p>\n\n<h4>Profile</h4>\n\n<p>Software Engineer at Treasure Data, OSS developer and enthusiast, committer of CRuby and JRuby, and Information Security Specialist</p>\n\n<h3><small>20:15 〜 22:00</small> Open Networking</h3>\n\n<p>Discuss the presentations or anything else Ruby related with the other attendees.</p>\n\n<h2>Warning about the Venue</h2>\n\n<ul>\n<li>From 7pm, the front doors will no longer open automatically from the outside. Wait for someone to come out, and then you can go in.</li>\n<li>From 9pm, the front door will be locked, so please make sure you arrive before then.</li>\n</ul>",
				PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
				Participants: 48,
				Waitlisted:   0,
				TicketLimit:  50,
			},
			want: &Event{
				Title:        "900K records per second with Ruby, Java, and JRuby",
				ID:           28319,
				StartsAt:     time.Date(2015, 8, 13, 10, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2015, 8, 13, 13, 0, 0, 0, time.UTC),
				VenueName:    "VOYAGE GROUP",
				Address:      "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
				Lat:          fp(35.6553195),
				Long:         fp(139.6937795),
				PublishedAt:  time.Date(2015, 7, 13, 23, 48, 29, 463000000, time.UTC),
				UpdatedAt:    time.Date(2018, 5, 11, 0, 7, 44, 270000000, time.UTC),
				Group:        24,
				Description:  "<h2>Agenda</h2>\n\n<h3><small>19:00 〜 19:45</small> Doors open</h3>\n\n<p>Grab a drink and catch up with your fellow Rubyists.</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">Hiroshi Nakamura (@nahi)</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">Treasure Data</a> is cloud analytics infrastructure. As of the beginning of July 2015, it stores 16 trillion records of customers data, and grows by 900,000 records per second. For handling this data, we are leveraging Ruby, Java, JRuby and many OSS components. In this presentation I'll introduce the system architecture and will discuss how we utilize these languages and OSS components to create a robust, performant, and flexible infrastructure.</p>\n\n<h4>Profile</h4>\n\n<p>Software Engineer at Treasure Data, OSS developer and enthusiast, committer of CRuby and JRuby, and Information Security Specialist</p>\n\n<h3><small>20:15 〜 22:00</small> Open Networking</h3>\n\n<p>Discuss the presentations or anything else Ruby related with the other attendees.</p>\n\n<h2>Warning about the Venue</h2>\n\n<ul>\n<li>From 7pm, the front doors will no longer open automatically from the outside. Wait for someone to come out, and then you can go in.</li>\n<li>From 9pm, the front door will be locked, so please make sure you arrive before then.</li>\n</ul>",
				PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
				Participants: 48,
				Waitlisted:   0,
				TicketLimit:  50,
			},
		},
		{
			name: "lat and long are empty",
			fields: fields{
				Title:        "900K records per second with Ruby, Java, and JRuby",
				ID:           28319,
				StartsAt:     "2015-08-13T10:00:00.000Z",
				EndsAt:       "2015-08-13T13:00:00.000Z",
				VenueName:    "VOYAGE GROUP",
				Address:      "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
				Lat:          nil,
				Long:         nil,
				PublishedAt:  "2015-07-13T23:48:29.463Z",
				UpdatedAt:    "2018-05-11T00:07:44.270Z",
				Group:        24,
				Description:  "<h2>Agenda</h2>\n\n<h3><small>19:00 〜 19:45</small> Doors open</h3>\n\n<p>Grab a drink and catch up with your fellow Rubyists.</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">Hiroshi Nakamura (@nahi)</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">Treasure Data</a> is cloud analytics infrastructure. As of the beginning of July 2015, it stores 16 trillion records of customers data, and grows by 900,000 records per second. For handling this data, we are leveraging Ruby, Java, JRuby and many OSS components. In this presentation I'll introduce the system architecture and will discuss how we utilize these languages and OSS components to create a robust, performant, and flexible infrastructure.</p>\n\n<h4>Profile</h4>\n\n<p>Software Engineer at Treasure Data, OSS developer and enthusiast, committer of CRuby and JRuby, and Information Security Specialist</p>\n\n<h3><small>20:15 〜 22:00</small> Open Networking</h3>\n\n<p>Discuss the presentations or anything else Ruby related with the other attendees.</p>\n\n<h2>Warning about the Venue</h2>\n\n<ul>\n<li>From 7pm, the front doors will no longer open automatically from the outside. Wait for someone to come out, and then you can go in.</li>\n<li>From 9pm, the front door will be locked, so please make sure you arrive before then.</li>\n</ul>",
				PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
				Participants: 48,
				Waitlisted:   0,
				TicketLimit:  50,
			},
			want: &Event{
				Title:        "900K records per second with Ruby, Java, and JRuby",
				ID:           28319,
				StartsAt:     time.Date(2015, 8, 13, 10, 0, 0, 0, time.UTC),
				EndsAt:       time.Date(2015, 8, 13, 13, 0, 0, 0, time.UTC),
				VenueName:    "VOYAGE GROUP",
				Address:      "東京都渋谷区神泉町8-16 渋谷ファーストプレイス8F",
				Lat:          nil,
				Long:         nil,
				PublishedAt:  time.Date(2015, 7, 13, 23, 48, 29, 463000000, time.UTC),
				UpdatedAt:    time.Date(2018, 5, 11, 0, 7, 44, 270000000, time.UTC),
				Group:        24,
				Description:  "<h2>Agenda</h2>\n\n<h3><small>19:00 〜 19:45</small> Doors open</h3>\n\n<p>Grab a drink and catch up with your fellow Rubyists.</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">Hiroshi Nakamura (@nahi)</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">Treasure Data</a> is cloud analytics infrastructure. As of the beginning of July 2015, it stores 16 trillion records of customers data, and grows by 900,000 records per second. For handling this data, we are leveraging Ruby, Java, JRuby and many OSS components. In this presentation I'll introduce the system architecture and will discuss how we utilize these languages and OSS components to create a robust, performant, and flexible infrastructure.</p>\n\n<h4>Profile</h4>\n\n<p>Software Engineer at Treasure Data, OSS developer and enthusiast, committer of CRuby and JRuby, and Information Security Specialist</p>\n\n<h3><small>20:15 〜 22:00</small> Open Networking</h3>\n\n<p>Discuss the presentations or anything else Ruby related with the other attendees.</p>\n\n<h2>Warning about the Venue</h2>\n\n<ul>\n<li>From 7pm, the front doors will no longer open automatically from the outside. Wait for someone to come out, and then you can go in.</li>\n<li>From 9pm, the front door will be locked, so please make sure you arrive before then.</li>\n</ul>",
				PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
				Participants: 48,
				Waitlisted:   0,
				TicketLimit:  50,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &rawEvent{
				Title:        tt.fields.Title,
				ID:           tt.fields.ID,
				StartsAt:     tt.fields.StartsAt,
				EndsAt:       tt.fields.EndsAt,
				VenueName:    tt.fields.VenueName,
				Address:      tt.fields.Address,
				Lat:          tt.fields.Lat,
				Long:         tt.fields.Long,
				PublishedAt:  tt.fields.PublishedAt,
				UpdatedAt:    tt.fields.UpdatedAt,
				Group:        tt.fields.Group,
				Description:  tt.fields.Description,
				PublicURL:    tt.fields.PublicURL,
				Participants: tt.fields.Participants,
				Waitlisted:   tt.fields.Waitlisted,
				TicketLimit:  tt.fields.TicketLimit,
			}
			got, err := e.toEvent()

			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
