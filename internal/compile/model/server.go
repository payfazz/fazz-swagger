package model

type Server struct {
	URL         *string                     `yaml:"url,omitempty"`
	Description *string                     `yaml:"description,omitempty"`
	Variables   []map[string]ServerVariable `yaml:"variables,omitempty"`
}

type ServerVariable struct {
	Enum        []string `yaml:"enum,omitempty"`
	Default     *string  `yaml:"default,omitempty"`
	Description *string  `yaml:"description,omitempty"`
}
