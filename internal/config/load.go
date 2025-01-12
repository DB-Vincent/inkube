package config

import (
	"github.com/BurntSushi/toml"
)

func Load(configPath string) (*Config, error) {
	var cfg Config
	if _, err := toml.DecodeFile(configPath, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
