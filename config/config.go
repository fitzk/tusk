package config

import (
	"github.com/rliebz/tusk/config/option"
	"github.com/rliebz/tusk/config/task"
)

// Config is a struct representing the format for configuration settings.
type Config struct {
	Name  string `yaml:"name"`
	Usage string `yaml:"usage"`

	Tasks   map[string]*task.Task `yaml:"tasks"`
	Options option.Options        `yaml:"options,omitempty"`
}

// UnmarshalYAML unmarshals and assigns names to options and tasks.
func (c *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	type configType Config // Use new type to avoid recursion
	if err := unmarshal((*configType)(c)); err != nil {
		return err
	}

	for name, t := range c.Tasks {
		t.Name = name
	}

	return nil
}
