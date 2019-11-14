package model

type Info struct {
	Title          *string  `yaml:"title,omitempty"`
	Description    *string  `yaml:"description,omitempty"`
	TermsOfService *string  `yaml:"termsOfService,omitempty"`
	Contact        *Contact `yaml:"contact,omitempty"`
	License        *License `yaml:"license,omitempty"`
	Version        *string  `yaml:"version,omitempty"`
}

type Contact struct {
	Name  *string `yaml:"name,omitempty"`
	URL   *string `yaml:"url,omitempty"`
	Email *string `yaml:"email,omitempty"`
}

type License struct {
	Name *string `yaml:"name,omitempty"`
	URL  *string `yaml:"url,omitempty"`
}
