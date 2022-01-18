package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Item struct {
	title string
	body  string
}

func main() {
	var reply Item
	var db []Item

	client, err := rpc.DialHTTP("tcp", "localhost:4040")

	if err != nil {
		log.Fatal("Connection Error: ", err)
	}

	a := Item{"First", "A First Item"}

	client.Call("API.AddItem", a, &reply)
	client.Call("API.GetDB", "", &db)

	fmt.Println("Database: ", db)
}
