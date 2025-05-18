package main

import (
	"fmt"
	"log"
	"os"

	"cmd/netbirdfunc"
)

func main() {
	apiToken := os.Getenv("NETBIRD_API_TOKEN")
	if apiToken == "" {
		log.Fatal("NETBIRD_API_TOKEN environment variable not set")
	}

	client := netbirdfunc.NewClient(apiToken)

	// Get Users
	users, err := client.GetUsers()
	if err != nil {
		log.Fatalf("Error fetching users: %v", err)
	}
	fmt.Println("Users:")
	for _, user := range users {
		fmt.Printf(" - %s (%s)\n", user.Name, user.Email)
	}

	// Get Groups
	groups, err := client.GetGroups()
	if err != nil {
		log.Fatalf("Error fetching groups: %v", err)
	}
	fmt.Println("\nGroups:")
	for _, group := range groups {
		fmt.Printf(" - %s (Peers: %d)\n", group.Name, group.PeersCount)
	}
}
