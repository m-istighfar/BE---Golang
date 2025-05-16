package main

import (
	"Yuk-Ujian/internal/config"
	"Yuk-Ujian/internal/delivery/http/server"
	"Yuk-Ujian/internal/pkg/logger"
)

func main() {
	cfg := config.InitConfig()

	logger.SetLogrusLogger(cfg)
	server.StartGinHttpServer(cfg)
}
