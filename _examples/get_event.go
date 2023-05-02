package main

import (
	"fmt"
	"github.com/sue445/go-doorkeeper"
	"os"
	"strconv"
)

func main() {
	accessToken := os.Getenv("DOORKEEPER_ACCESS_TOKEN")
	if accessToken == "" {
		panic("DOORKEEPER_ACCESS_TOKEN is required")
	}

	eventID := 28319
	if os.Getenv("DOORKEEPER_EVENT_ID") != "" {
		eventID, _ = strconv.Atoi(os.Getenv("DOORKEEPER_EVENT_ID"))
	}

	client := doorkeeper.NewClient(accessToken)

	if os.Getenv("DOORKEEPER_API_ENDPOINT") != "" {
		client.APIEndpoint = os.Getenv("DOORKEEPER_API_ENDPOINT")
	}

	event, rateLimit, err := client.GetEvent(eventID)

	if err != nil {
		panic(fmt.Sprintf("err=%+v", err))
	}

	fmt.Printf("event=%+v\n", event)
	fmt.Printf("rateLimit=%+v\n", rateLimit)
}
