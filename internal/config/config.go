package config

import (
	"io"
	"os"

	pconfig "github.com/go-void/portal/pkg/config"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	DNSOptions    *pconfig.Config `toml:"dns"`
	RouterOptions RouterOptions   `toml:"router"`
	StoreOptions  StoreOptions    `toml:"store"`
}

type RouterOptions struct {
	Port int    `toml:"port"`
	Path string `toml:"path"`
}

type StoreOptions struct {
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
	Backend  string `toml:"backend"`
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
}

// Default returns a config with default values
func Default() *Config {
	return &Config{
		DNSOptions: pconfig.Default(),
		RouterOptions: RouterOptions{
			Port: 8080,
			Path: "",
		},
	}
}

// Read reads a TOML config file
func Read(path string) (*Config, error) {
	if path == "" {
		return Default(), nil
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}

	cfg := new(Config)
	err = toml.Unmarshal(b, cfg)
	return cfg, err
}

// Write writes a TOML config file
func Write(path string, c *Config) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	b, err := toml.Marshal(c)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	return err
}

// Validate validates the config
func (c *Config) Validate() error {
	return c.DNSOptions.Validate()
}
