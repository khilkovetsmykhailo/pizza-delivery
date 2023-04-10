package main

import "github.com/marijakljestan/golang-web-app/src/startup"

func main() {
	server := startup.NewServer()
	server.Start()
}
