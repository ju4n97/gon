package config

import "fmt"

func NewConfigLoaderFactory(source string) (ConfigLoader, error) {
	switch source {
	case "env":
		return &EnvConfigLoader{}, nil
	case "json":
		return &JsonConfigLoader{}, nil
	case "yaml":
		return &YamlConfigLoader{}, nil
	default:
		return nil, fmt.Errorf("Invalid config source: %s", source)
	}
}
