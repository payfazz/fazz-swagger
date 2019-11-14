package model

type Tag struct {
	Name         *string       `yaml:"name,omitempty"`
	Description  *string       `yaml:"description,omitempty"`
	ExternalDocs *ExternalDocs `yaml:"externalDocs,omitempty"`
}
