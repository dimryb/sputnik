package config

import (
	"fmt"
	"path"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	CalendarConfig struct {
		HTTP `yaml:"http"`
		Log  `yaml:"log"`
	}

	Log struct {
		Level string `yaml:"level" env:"LOG_LEVEL"`
	}

	HTTP struct {
		Host              string        `yaml:"host" env:"HTTP_HOST"`
		Port              string        `yaml:"port" env:"HTTP_PORT"`
		ReadTimeout       time.Duration `yaml:"readTimeout"`
		WriteTimeout      time.Duration `yaml:"writeTimeout"`
		IdleTimeout       time.Duration `yaml:"idleTimeout"`
		ReadHeaderTimeout time.Duration `yaml:"readHeaderTimeout"`
	}
)

func Load(configPath string, target any) error {
	err := cleanenv.ReadConfig(path.Join("./", configPath), target)
	if err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(target)
	if err != nil {
		return fmt.Errorf("error updating env: %w", err)
	}

	return nil
}
