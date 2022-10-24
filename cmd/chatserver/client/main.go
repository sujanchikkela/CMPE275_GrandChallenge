package main

import (
	"flag"

	client "grandChallenge1/grpcchat123/internal/chatserver/client"
)

func main() {

	//Read inputs from the user for name and Chat room name
	name := flag.String("N", "Satvik", "The name of the user")
	room_name := flag.String("R", "default", "The name of the chat room")
	blocked_name := flag.String("B", "NA", "The name of the blocked user")
	flag.Parse()

	client := &client.Client{
		Name:        *name,
		RoomName:    *room_name,
		BlockedName: *blocked_name,
	}
	client.Exec()
}
