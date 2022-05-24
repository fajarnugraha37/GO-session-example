package main

import (
	route "go-server-session/route"
)

func main() {
	server := route.GetServer()

	server.Listen(":5000")
}
