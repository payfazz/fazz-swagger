package model

// ExternalDocs is a struct to define OAS3's External Docs Data Type
type ExternalDocs struct {
	URL         *string `yaml:"url,omitempty"`
	Description *string `yaml:"description,omitempty"`
}
