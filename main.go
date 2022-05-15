package main

import (
	"awesomeProject3/config"
	"awesomeProject3/routes"
)

func main() {
	config.Init()
	r := routes.NewRouter()
	_ = r.Run(config.HttpPort)
}
