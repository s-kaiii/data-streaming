package config

import (
	"log"

	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ServerAddr   string   `yaml:"server.address"`
	KafkaBrokers []string `yaml:"kafka.brokers"`
	LoggerConfig string   `yaml:"logger.level"`
}

// TODO check log
func LoadConfig() *Config {
	file, err := os.Open("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	config := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(config); err != nil {
		log.Fatalf("Failed to decode config: %v", err)
	}

	return config
}
