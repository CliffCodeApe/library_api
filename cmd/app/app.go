package app

import (
	"library_api/config"
	"library_api/internal/server"
	"library_api/pkg/token"
)

func Start() {
	config.Load()
	token.Load()
	server.Run()

}
