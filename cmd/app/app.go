package app

import (
	"library_api/config"
	"library_api/internal/server"
)

func Start() {
	config.Load()
	server.Run()

}
