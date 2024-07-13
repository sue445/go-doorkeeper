package doorkeeper

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func Test_newRawRateLimitFromJson(t *testing.T) {
	type args struct {
		jsonText string
	}
	tests := []struct {
		name string
		args args
		want *rawRateLimit
	}{
		{
			name: "successful",
			args: args{
				jsonText: `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`,
			},
			want: &rawRateLimit{
				Name:      "authenticated API",
				Period:    300,
				Limit:     300,
				Remaining: 299,
				Until:     "2019-07-29T15:15:00Z",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newRawRateLimitFromJSON(tt.args.jsonText)

			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_rawRateLimit_toRateLimit(t *testing.T) {
	type fields struct {
		name      string
		period    int
		limit     int
		remaining int
		until     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *RateLimit
		wantErr bool
	}{
		{
			name: "successful",
			fields: fields{
				name:      "authenticated API",
				period:    300,
				limit:     300,
				remaining: 299,
				until:     "2019-07-29T15:15:00Z",
			},
			want: &RateLimit{
				Name:      "authenticated API",
				Period:    300,
				Limit:     300,
				Remaining: 299,
				Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &rawRateLimit{
				Name:      tt.fields.name,
				Period:    tt.fields.period,
				Limit:     tt.fields.limit,
				Remaining: tt.fields.remaining,
				Until:     tt.fields.until,
			}
			got, err := r.toRateLimit()

			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
