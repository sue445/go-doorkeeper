package doorkeeper

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
	"time"
)

func readTestData(filename string) string {
	buf, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(buf)
}

func TestClient_GetGroup(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "GetGroup-ja.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	group, rateLimit, err := c.GetGroup("trbmeetup")

	assert.NoError(t, err)

	wantGroup := &Group{
		ID:           24,
		Name:         "Tokyo Rubyist Meetup",
		CountryCode:  "JP",
		Logo:         "https://dzpp79ucibp5a.cloudfront.net/groups_logos/24_normal_1371637374_200px-Ruby_logo.png",
		Description:  "<p>Tokyo Rubyist Meetup (trbmeetup)は、日本のRubyistと世界のRubyistとをつなげるための場になることを目指して設立されました。定例会には、東京近郊に住んでいる海外出身のRubyistたちと日本人Rubyistたちが参加します。例会の公用語は英語になりますが、英語が苦手な方も、一緒に英語の練習をするくらいのつもりでお気軽にご参加ください。</p>\n",
		PublicURL:    "https://trbmeetup.doorkeeper.jp/",
		MembersCount: 2056,
	}
	assert.Equal(t, wantGroup, group)

	wantRateLimit := &RateLimit{
		Name:      "authenticated API",
		Period:    300,
		Limit:     300,
		Remaining: 299,
		Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
	}
	assert.Equal(t, wantRateLimit, rateLimit)
}

func TestClient_GetGroup_WithLocale(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup?locale=en",
		func(req *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "GetGroup-en.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	group, rateLimit, err := c.GetGroup("trbmeetup", WithLocale("en"))

	assert.NoError(t, err)

	wantGroup := &Group{
		ID:           24,
		Name:         "Tokyo Rubyist Meetup",
		CountryCode:  "JP",
		Logo:         "https://dzpp79ucibp5a.cloudfront.net/groups_logos/24_normal_1371637374_200px-Ruby_logo.png",
		Description:  "<p>Tokyo Rubyist Meetup (trbmeetup) is an event that seeks to help bridge the Japan and international ruby and ruby on rails community. It will hold regular meetings where Japanese Rubyists can communicate with international Rubyists living in Tokyo. Meetings will be held in English, but anyone is encouraged to participate regardless of their ability.</p>\n",
		PublicURL:    "https://trbmeetup.doorkeeper.jp/",
		MembersCount: 2056,
	}
	assert.Equal(t, wantGroup, group)

	wantRateLimit := &RateLimit{
		Name:      "authenticated API",
		Period:    300,
		Limit:     300,
		Remaining: 299,
		Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
	}
	assert.Equal(t, wantRateLimit, rateLimit)
}
