package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App        AppConfig
	HttpServer HttpServerConfig
	Database   DatabaseConfig
}

type AppConfig struct {
	Environment string
	BCryptCost  int
	BaseURL     string
}

type HttpServerConfig struct {
	Host              string
	Port              int
	GracePeriod       int
	MaxUploadFileSize int64
}

type DatabaseConfig struct {
	Port                  int
	Host                  string
	DbName                string
	Username              string
	Password              string
	Sslmode               string
	MaxIdleConn           int
	MaxOpenConn           int
	MaxConnLifetimeMinute int
}

func InitConfig() *Config {
	env := os.Getenv("APP_ENVIRONMENT")

	if env == "development" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Warning: Error loading .env file (development mode)")
		}
	}

	return &Config{
		App:        initAppConfig(),
		Database:   initDbConfig(),
		HttpServer: initHttpServerConfig(),
	}
}

func initDbConfig() DatabaseConfig {
	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	name := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSL_MODE")

	port, err := strconv.ParseInt(os.Getenv("DB_PORT"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_PORT")
	}

	maxIdleConn, err := strconv.ParseInt(os.Getenv("DB_MAX_IDLE_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_IDLE_CONN")
	}

	maxOpenConn, err := strconv.ParseInt(os.Getenv("DB_MAX_OPEN_CONN"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_MAX_OPEN_CONN")
	}

	connMaxLifetime, err := strconv.ParseInt(os.Getenv("DB_CONN_MAX_LIFETIME"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse DB_CONN_MAX_LIFETIME")
	}

	return DatabaseConfig{
		Port:                  int(port),
		Host:                  host,
		DbName:                name,
		Username:              user,
		Password:              password,
		Sslmode:               sslMode,
		MaxIdleConn:           int(maxIdleConn),
		MaxOpenConn:           int(maxOpenConn),
		MaxConnLifetimeMinute: int(connMaxLifetime),
	}
}

func initHttpServerConfig() HttpServerConfig {
	host := os.Getenv("HTTP_SERVER_HOST")

	portStr := os.Getenv("PORT")
	if portStr == "" {
		portStr = os.Getenv("HTTP_SERVER_PORT")
	}
	if portStr == "" {
		portStr = "8000"
	}

	port, err := strconv.ParseInt(portStr, 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_PORT")
	}

	gracePeriod, err := strconv.ParseInt(os.Getenv("HTTP_SERVER_GRACE_PERIOD"), 10, 32)
	if err != nil {
		log.Fatal("cannot parse HTTP_SERVER_GRACE_PERIOD")
	}

	return HttpServerConfig{
		Host:        host,
		Port:        int(port),
		GracePeriod: int(gracePeriod),
	}
}

func initAppConfig() AppConfig {
	environment := os.Getenv("APP_ENVIRONMENT")

	return AppConfig{
		Environment: environment,
	}
}
