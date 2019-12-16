package model

// Tag is a struct to define OAS3's Tag Data Type
type Tag struct {
	Name         *string       `yaml:"name,omitempty"`
	Description  *string       `yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty"`
}
