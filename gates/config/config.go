package config

import (
	"github.com/b0ralgin/pow_test/domain"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
)


type Config struct {
	Port       string       `envconfig:"PORT" default:"8080"`
	Difficulty uint8        `envconfig:"DIFFICULTY" default:"3"`
	Size       uint8        `envconfig:"SIZE" default:"4"`
	Algo       domain.Algoritm       `envconfig:"ALGO" default:"DefaultAlgorithm"`
}

// LoadConfig загружает конфигурацию из переменных окружения
func LoadConfig() (*Config, error) {
	var cfg Config

	// Считываем переменные окружения в структуру
	err := envconfig.Process("", &cfg)
	if err != nil {
	return nil ,errors.Wrap(err, "Loading config error")
	}

	return &cfg, nil
}