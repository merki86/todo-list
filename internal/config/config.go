package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"dev"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
	Server      `yaml:"server"`
}

type Server struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"30s"`
}

func MustRead() *Config {
	// Get the path to the configuration file from the environment variable
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	// Check if the configuration file exists
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("The configuration file does not exist: %s", configPath)
	}

	// Read the configuration file
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read the configuration file: %v", err)
	}

	return &cfg
}
