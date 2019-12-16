package model

// Swagger is a struct to define OAS3's Data Structure
type Swagger struct {
	OpenApi      string                  `yaml:"openapi,omitempty"`
	Info         *Info                   `yaml:"info,omitempty"`
	Servers      []Server                `yaml:"servers,omitempty"`
	Tags         []Tag                   `yaml:"tags,omitempty"`
	ExternalDocs *ExternalDocs           `yaml:"externalDocs,omitempty"`
	Security     []map[string][]string   `yaml:"security,omitempty"`
	Paths        *map[string]interface{} `yaml:"paths,omitempty"`
	Components   *Component              `yaml:"components,omitempty"`
}
