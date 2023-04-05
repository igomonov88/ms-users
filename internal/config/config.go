package config

import (
	"github.com/spf13/pflag"

	"github.com/igomonov88/ms-users/internal/db"
	"github.com/igomonov88/ms-users/internal/ms"
)

type Config struct {
	MS ms.Config
	DB db.Config
}

func Must(cfg *Config, err error) Config {
	if err != nil {
		panic(err)
	}

	return *cfg
}

func Read() (*Config, error) {
	pflag.Parse()

	src, err := NewSource()
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := src.ReadConfig(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
