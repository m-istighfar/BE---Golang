package main

import (
	"DRX_Test/internal/config"
	"DRX_Test/internal/delivery/http/server"
	"DRX_Test/internal/pkg/logger"
)

func main() {
	cfg := config.InitConfig()

	logger.SetLogrusLogger(cfg)
	server.StartGinHttpServer(cfg)
}
