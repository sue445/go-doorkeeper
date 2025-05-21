package doorkeeper

import (
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func readTestData(filename string) string {
	buf, err := os.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	return string(buf)
}

func fp(f float64) *float64 {
	return &f
}

func tp(t time.Time) *time.Time {
	return &t
}

func TestClient_GetEvents(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/events?page=1",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "events.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	events, rateLimit, err := c.GetEvents(&GetEventsParams{Page: 1})

	require.NoError(t, err)

	wantEvent := &Event{
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
		Description:  "<h2>アジェンダ</h2>\n\n<h3><small>19:00 〜 19:45</small> 会場</h3>\n\n<p>飲み物を片手にRubyist同士の交流を深めて下さい。</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">中村浩士、なひ</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">トレジャーデータ</a>はクラウド型のデータマネジメントサービスです。分析対象データとして、2015年7月当初時点で16兆レコードを格納しており、さらに毎秒90万レコードの勢いで増加していっています。このデータを処理するため、我々はRuby、Java、JRubyそして各種OSSコンポーネントを活用しています。この発表ではまずそのシステムアーキテクチャを紹介し、続いて堅牢、高性能かつ柔軟なデータマネジメント基盤を構築するために、これらプログラミング言語とOSSコンポーネントをどのように利用しているかについて説明します。</p>\n\n<h4>プロフィール</h4>\n\n<p>トレジャーデータ株式会社ソフトウェアエンジニア、OSS開発者、CRubyおよびJRubyコミッタ、情報セキュリティスペシャリスト</p>\n\n<h3><small>20:15 〜 22:00</small> オープンネットワーク</h3>\n\n<p>参加者同士で当日のプレゼンやRubyに関することについてご歓談下さい。</p>\n\n<h2>会場の注意</h2>\n\n<ul>\n<li>19時以降、正面玄関は外からは開きません。中から人が出てくるのを待てば、入館できます。</li>\n<li>21時以降、正面玄関はロックされてしまいますので、21時より前に到着するようにして下さい。</li>\n</ul>",
		PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
		Participants: 48,
		Waitlisted:   0,
		TicketLimit:  50,
	}
	assert.Len(t, events, 1)
	assert.Equal(t, wantEvent, events[0])

	wantRateLimit := &RateLimit{
		Name:      "authenticated API",
		Period:    300,
		Limit:     300,
		Remaining: 299,
		Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
	}
	assert.Equal(t, wantRateLimit, rateLimit)
}

func TestClient_GetGroupEvents(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup/events?page=1",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "events.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	events, rateLimit, err := c.GetGroupEvents("trbmeetup", &GetEventsParams{Page: 1})

	require.NoError(t, err)

	wantEvent := &Event{
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
		Description:  "<h2>アジェンダ</h2>\n\n<h3><small>19:00 〜 19:45</small> 会場</h3>\n\n<p>飲み物を片手にRubyist同士の交流を深めて下さい。</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">中村浩士、なひ</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">トレジャーデータ</a>はクラウド型のデータマネジメントサービスです。分析対象データとして、2015年7月当初時点で16兆レコードを格納しており、さらに毎秒90万レコードの勢いで増加していっています。このデータを処理するため、我々はRuby、Java、JRubyそして各種OSSコンポーネントを活用しています。この発表ではまずそのシステムアーキテクチャを紹介し、続いて堅牢、高性能かつ柔軟なデータマネジメント基盤を構築するために、これらプログラミング言語とOSSコンポーネントをどのように利用しているかについて説明します。</p>\n\n<h4>プロフィール</h4>\n\n<p>トレジャーデータ株式会社ソフトウェアエンジニア、OSS開発者、CRubyおよびJRubyコミッタ、情報セキュリティスペシャリスト</p>\n\n<h3><small>20:15 〜 22:00</small> オープンネットワーク</h3>\n\n<p>参加者同士で当日のプレゼンやRubyに関することについてご歓談下さい。</p>\n\n<h2>会場の注意</h2>\n\n<ul>\n<li>19時以降、正面玄関は外からは開きません。中から人が出てくるのを待てば、入館できます。</li>\n<li>21時以降、正面玄関はロックされてしまいますので、21時より前に到着するようにして下さい。</li>\n</ul>",
		PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
		Participants: 48,
		Waitlisted:   0,
		TicketLimit:  50,
	}
	assert.Len(t, events, 1)
	assert.Equal(t, wantEvent, events[0])

	wantRateLimit := &RateLimit{
		Name:      "authenticated API",
		Period:    300,
		Limit:     300,
		Remaining: 299,
		Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
	}
	assert.Equal(t, wantRateLimit, rateLimit)
}

func TestClient_GetEvent(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/events/28319",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "event-ja.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	event, rateLimit, err := c.GetEvent(28319)

	require.NoError(t, err)

	wantEvent := &Event{
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
		Description:  "<h2>アジェンダ</h2>\n\n<h3><small>19:00 〜 19:45</small> 会場</h3>\n\n<p>飲み物を片手にRubyist同士の交流を深めて下さい。</p>\n\n<h3><small>19:45 〜 20:15</small> 900K records per second with Ruby, Java, and JRuby<small> <a href=\"https://twitter.com/nahi\" rel=\"nofollow\">中村浩士、なひ</a></small></h3>\n\n<p><a href=\"http://TreasureData.com\" rel=\"nofollow\">トレジャーデータ</a>はクラウド型のデータマネジメントサービスです。分析対象データとして、2015年7月当初時点で16兆レコードを格納しており、さらに毎秒90万レコードの勢いで増加していっています。このデータを処理するため、我々はRuby、Java、JRubyそして各種OSSコンポーネントを活用しています。この発表ではまずそのシステムアーキテクチャを紹介し、続いて堅牢、高性能かつ柔軟なデータマネジメント基盤を構築するために、これらプログラミング言語とOSSコンポーネントをどのように利用しているかについて説明します。</p>\n\n<h4>プロフィール</h4>\n\n<p>トレジャーデータ株式会社ソフトウェアエンジニア、OSS開発者、CRubyおよびJRubyコミッタ、情報セキュリティスペシャリスト</p>\n\n<h3><small>20:15 〜 22:00</small> オープンネットワーク</h3>\n\n<p>参加者同士で当日のプレゼンやRubyに関することについてご歓談下さい。</p>\n\n<h2>会場の注意</h2>\n\n<ul>\n<li>19時以降、正面玄関は外からは開きません。中から人が出てくるのを待てば、入館できます。</li>\n<li>21時以降、正面玄関はロックされてしまいますので、21時より前に到着するようにして下さい。</li>\n</ul>",
		PublicURL:    "https://trbmeetup.doorkeeper.jp/events/28319",
		Participants: 48,
		Waitlisted:   0,
		TicketLimit:  50,
	}
	assert.Equal(t, wantEvent, event)

	wantRateLimit := &RateLimit{
		Name:      "authenticated API",
		Period:    300,
		Limit:     300,
		Remaining: 299,
		Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
	}
	assert.Equal(t, wantRateLimit, rateLimit)
}

func TestClient_GetEvent_WithLocale(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/events/28319?locale=en",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "event-en.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	event, rateLimit, err := c.GetEvent(28319, WithLocale("en"))

	require.NoError(t, err)

	wantEvent := &Event{
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
	}
	assert.Equal(t, wantEvent, event)

	wantRateLimit := &RateLimit{
		Name:      "authenticated API",
		Period:    300,
		Limit:     300,
		Remaining: 299,
		Until:     time.Date(2019, 7, 29, 15, 15, 0, 0, time.UTC),
	}
	assert.Equal(t, wantRateLimit, rateLimit)
}

func TestClient_GetGroup(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "group-ja.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	group, rateLimit, err := c.GetGroup("trbmeetup")

	require.NoError(t, err)

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
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup?locale=en",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "group-en.json")))
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	group, rateLimit, err := c.GetGroup("trbmeetup", WithLocale("en"))

	require.NoError(t, err)

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

func TestClient_GetGroup_WithoutRateLimitHeader(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/trbmeetup?locale=en",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(200, readTestData(filepath.Join("testdata", "group-en.json")))
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	group, rateLimit, err := c.GetGroup("trbmeetup", WithLocale("en"))

	require.NoError(t, err)

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

	wantRateLimit := &RateLimit{}
	assert.Equal(t, wantRateLimit, rateLimit)
}

func TestClient_GetGroup_NotFound(t *testing.T) {
	httpmock.Activate(t)

	httpmock.RegisterResponder("GET", "https://api.doorkeeper.jp/groups/not-found",
		func(_ *http.Request) (*http.Response, error) {
			resp := httpmock.NewStringResponse(404, "")
			resp.Header.Set("X-Ratelimit", `{"name":"authenticated API","period":300,"limit":300,"remaining":299,"until":"2019-07-29T15:15:00Z"}`)
			return resp, nil
		},
	)

	c := NewClient("DOORKEEPER_ACCESS_TOKEN")
	group, rateLimit, err := c.GetGroup("not-found")

	if assert.EqualError(t, err, "404 Not Found") {
		assert.Nil(t, group)
		assert.Nil(t, rateLimit)
	}
}

func TestClient_buildURL(t *testing.T) {
	c := NewClient("DOORKEEPER_ACCESS_TOKEN")

	type args struct {
		path   string
		values url.Values
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "empty values",
			args: args{
				path:   "/events",
				values: url.Values{},
			},
			want: "https://api.doorkeeper.jp/events",
		},
		{
			name: "sort and since and until",
			args: args{
				path: "/groups/trbmeetup/events",
				values: url.Values{
					"sort":  []string{"published_at"},
					"since": []string{"2015-03-03"},
					"until": []string{"2015-08-30"},
				},
			},
			want: "https://api.doorkeeper.jp/groups/trbmeetup/events?since=2015-03-03&sort=published_at&until=2015-08-30",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.buildURL(tt.args.path, tt.args.values)

			if assert.NoError(t, err) {
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
