package main

import (
	"./Server"
)

func main() {
	go Server.ClientPost()
	go Server.GetRequest()
	Server.HandServer()
}
