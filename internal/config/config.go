package config

import(
	"flag"
	"os"
)

type Config struct {
	Grpc GRPCConfig
}

type GRPCConfig struct {
	Host string
	Port int
	Timeout int
}

func LoadConfig() *Config {
	return nil
}
