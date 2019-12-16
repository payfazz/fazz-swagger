package model

// Component is a struct to define OAS3's Component Data Type
type Component struct {
	Schemas         map[string]interface{} `yaml:"schemas,omitempty"`
	Responses       map[string]interface{} `yaml:"responses,omitempty"`
	Parameters      map[string]interface{} `yaml:"parameters,omitempty"`
	Examples        map[string]interface{} `yaml:"examples,omitempty"`
	RequestBodies   map[string]interface{} `yaml:"requestBodies,omitempty"`
	Headers         map[string]interface{} `yaml:"headers,omitempty"`
	SecuritySchemes map[string]interface{} `yaml:"securitySchemes,omitempty"`
	Links           map[string]interface{} `yaml:"links,omitempty"`
	Callbacks       map[string]interface{} `yaml:"callbacks,omitempty"`
}
