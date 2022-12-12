package main

import (
	"diamond/api"
	"diamond/config"
	log "github.com/sirupsen/logrus"
)

var version string = "dev"

func main() {
	log.Info("current Version " + version)
	api.Serve(config.HostAndPort())
}
