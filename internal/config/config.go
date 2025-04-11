package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	Rpc RpcConfig
}

type RpcConfig struct {
	Host string
	Port int
}

func NewConfig() *Config {

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Panic("error loading .env")
	}
	viper.AutomaticEnv()

	rpc := RpcConfig{
		Host: viper.GetString("RPC_HOST"),
		Port: viper.GetInt("RPC_PORT"),
	}

	return &Config{
		Rpc: rpc,
	}
}
