package config

import "fmt"

type JsonConfigLoader struct{}

// LoadConfig loads the global configuration from a JSON file.
func (l *JsonConfigLoader) LoadConfig() (*Config, error) {
	// TODO: Implement logic to load the global configuration from a JSON file
	Global = &Config{}

	if err := Global.Validate(); err != nil {
		return nil, fmt.Errorf("Invalid config: %v", err)
	}

	return nil, nil
}
