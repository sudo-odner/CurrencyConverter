package config

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ProtHttp ProtHttp `yaml:"protocol_http"`
}

type ProtHttp struct {
	Url Url `yaml:"url"`
}

type Url struct {
	Url_get_fiat             string `yaml:"url_get_fiat"`
	Url_get_cryptocurrencies string `yaml:"url_get_cryptocurrencies"`
	Url_convert_one_to_one   string `yaml:"url_convert_one_to_one"`
}

func Load() (*Config, error) {

	// Провекрка существования .env
	if err := godotenv.Load(); err != nil {
		log.Print(".env file found")
		return nil, err
	}

	configPath, exists := os.LookupEnv("CONFIG_PATH")
	if !exists {
		log.Fatal("CONFIG_PATH is not found")
		return nil, errors.New("CONFIG_PATH is not found")
	}

	// Провека на существования файла
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath)
		return nil, err
	}

	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
		return nil, err
	}

	return &cfg, nil
}
