package config

import (
	"os"

	"go.uber.org/config"
)

// Config struct to config file
type Config struct {
	Population struct {
		Size uint
	}
	Mutation struct {
		Cycles int
		Chance int
		Genes  int
	}
	Files struct {
		Stores string
		Order  string
	}
}

var configs Config

// GetConfig returns config instance
func GetConfig() Config {
	return configs
}

// InitConfigs init settings
func InitConfigs() error {
	provider, err := config.NewYAML(config.File(os.Getenv("CONFIG_FILE")))

	if err != nil {
		return err
	}

	if err := provider.Get(config.Root).Populate(&configs); err != nil {
		return err
	}

	return nil
}
