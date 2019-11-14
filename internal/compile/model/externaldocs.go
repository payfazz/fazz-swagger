package model

type ExternalDocs struct {
	URL         *string `yaml:"url,omitempty"`
	Description *string `yaml:"description,omitempty"`
}
