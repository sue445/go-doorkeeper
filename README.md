# go-doorkeeper
[Doorkeeper API](https://www.doorkeeper.jp/developer/api) client for Go

[![Build Status](https://github.com/sue445/go-doorkeeper/workflows/test/badge.svg?branch=master)](https://github.com/sue445/go-doorkeeper/actions?query=workflow%3Atest)
[![Maintainability](https://api.codeclimate.com/v1/badges/fcf6115e2d1c43780cb8/maintainability)](https://codeclimate.com/github/sue445/go-doorkeeper/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/fcf6115e2d1c43780cb8/test_coverage)](https://codeclimate.com/github/sue445/go-doorkeeper/test_coverage)
[![GoDoc](https://godoc.org/github.com/sue445/go-doorkeeper?status.svg)](https://godoc.org/github.com/sue445/go-doorkeeper)
[![Go Report Card](https://goreportcard.com/badge/github.com/sue445/go-doorkeeper)](https://goreportcard.com/report/github.com/sue445/go-doorkeeper)

## Example
```go
package main

import (
	"github.com/sue445/go-doorkeeper"
	"os"
)

func main() {
	accessToken := os.Getenv("DOORKEEPER_ACCESS_TOKEN")
	if accessToken == "" {
		panic("DOORKEEPER_ACCESS_TOKEN is required")
	}

	client := doorkeeper.NewClient(accessToken)

	// List all featured events
	events, rateLimit, err := client.GetEvents(&doorkeeper.GetEventsParams{})
	// more options
	events, rateLimit, err := client.GetEvents(&doorkeeper.GetEventsParams{Query: "golang", Sort: doorkeeper.SortByPublishedAt()})

	// List a community's events
	groupName := "trbmeetup"
	events, rateLimit, err := client.GetGroupEvents(groupName, &doorkeeper.GetEventsParams{})

	// Show a specific event
	eventID := 28319
	event, rateLimit, err := client.GetEvent(eventID)
	// or
	event, rateLimit, err := client.GetEvent(eventID, doorkeeper.WithLocale("en"))

	// Show a specific group
	groupName := "trbmeetup"
	group, rateLimit, err := client.GetGroup(groupName)
	// or
	group, rateLimit, err := client.GetGroup(groupName, doorkeeper.WithLocale("en"))
}
```

## Reference
https://godoc.org/github.com/sue445/go-doorkeeper
