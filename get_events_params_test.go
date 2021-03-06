package doorkeeper

import (
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
	"time"
)

func TestGetEventsParams_toMap(t *testing.T) {
	type fields struct {
		Page     int
		Locale   string
		Sort     SortEnum
		Since    *time.Time
		Until    *time.Time
		Query    string
		Callback string
	}
	tests := []struct {
		name   string
		fields fields
		want   url.Values
	}{
		{
			name: "full values",
			fields: fields{
				Page:     1,
				Locale:   "en",
				Sort:     SortByUpdatedAt(),
				Since:    tp(time.Date(2015, 8, 13, 10, 0, 0, 0, time.UTC)),
				Until:    tp(time.Date(2015, 9, 13, 10, 0, 0, 0, time.UTC)),
				Query:    "test",
				Callback: "func",
			},
			want: url.Values{
				"page":     []string{"1"},
				"locale":   []string{"en"},
				"sort":     []string{"updated_at"},
				"since":    []string{"2015-08-13"},
				"until":    []string{"2015-09-13"},
				"q":        []string{"test"},
				"callback": []string{"func"},
			},
		},
		{
			name:   "no values",
			fields: fields{},
			want:   url.Values{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &GetEventsParams{
				Page:     tt.fields.Page,
				Locale:   tt.fields.Locale,
				Sort:     tt.fields.Sort,
				Since:    tt.fields.Since,
				Until:    tt.fields.Until,
				Query:    tt.fields.Query,
				Callback: tt.fields.Callback,
			}

			got := p.toValues()
			assert.Equal(t, tt.want, got)
		})
	}
}
