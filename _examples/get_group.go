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

	group, rateLimit, err := client.GetGroup(groupName)

	if err != nil {
		panic(fmt.Sprintf("err=%+v", err))
	}

	fmt.Printf("group=%+v\n", group)
	fmt.Printf("rateLimit=%+v\n", rateLimit)
}
