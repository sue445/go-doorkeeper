package main

import (
	"fmt"
	"github.com/sue445/go-doorkeeper"
	"os"
)

func main() {
	accessToken := os.Getenv("DOORKEEPER_ACCESS_TOKEN")
	if accessToken == "" {
		panic("DOORKEEPER_ACCESS_TOKEN is required")
	}

	groupName := "trbmeetup"
	if os.Getenv("DOORKEEPER_GROUP") != "" {
		groupName = os.Getenv("DOORKEEPER_GROUP")
	}

	client := doorkeeper.NewClient(accessToken)

	if os.Getenv("DOORKEEPER_API_ENDPOINT") != "" {
		client.APIEndpoint = os.Getenv("DOORKEEPER_API_ENDPOINT")
	}

	events, rateLimit, err := client.GetGroupEvents(groupName, &doorkeeper.GetEventsParams{Sort: doorkeeper.SortByPublishedAt()})

	if err != nil {
		panic(fmt.Sprintf("err=%+v", err))
	}

	for i, event := range events {
		fmt.Printf("event[%d]=%+v\n", i, event)
	}
	fmt.Printf("rateLimit=%+v\n", rateLimit)
}
