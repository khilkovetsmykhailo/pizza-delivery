package main

import (
	"github.com/marijakljestan/golang-web-app/src/startup"
	"github.com/marijakljestan/golang-web-app/src/startup/config"
)

func main() {
	config := config.NewLocalConfig()
	server := startup.NewServer(config)
	server.Start()
}
