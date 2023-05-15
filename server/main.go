package main

import (
	"github.com/marijakljestan/golang-web-app/server/startup"
	"github.com/marijakljestan/golang-web-app/server/startup/config"
)

func main() {
	config := config.NewLocalConfig()
	server := startup.NewServer(config)
	server.Start()
}
