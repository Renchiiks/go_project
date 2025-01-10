package main

type Config struct {
	PostgresURL string

	Host string
	Port string
}

func initConfig() Config {
	var config Config

	config.PostgresURL = "db://localhost:9091"
	config.Host = "localhost"
	config.Port = "8080"

	return config
}
