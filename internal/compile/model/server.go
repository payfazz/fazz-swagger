package model

// Server is a struct to define OAS3's Server Data Type
type Server struct {
	URL         *string                     `yaml:"url,omitempty"`
	Description *string                     `yaml:"description,omitempty"`
	Variables   []map[string]ServerVariable `yaml:"variables,omitempty"`
}

// ServerVariable is a struct to define ServerVariable Data Type
type ServerVariable struct {
	Enum        []string `yaml:"enum,omitempty"`
	Default     *string  `yaml:"default,omitempty"`
	Description *string  `yaml:"description,omitempty"`
}
