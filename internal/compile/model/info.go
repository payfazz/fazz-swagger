package model

// Info is a struct to define OAS3's Info Data Type
type Info struct {
	Title          *string  `yaml:"title,omitempty"`
	Description    *string  `yaml:"description,omitempty"`
	TermsOfService *string  `yaml:"termsOfService,omitempty"`
	Contact        *Contact `yaml:"contact,omitempty"`
	License        *License `yaml:"license,omitempty"`
	Version        *string  `yaml:"version,omitempty"`
}

// Contact is a struct to define OAS3's Contact Data Type
type Contact struct {
	Name  *string `yaml:"name,omitempty"`
	URL   *string `yaml:"url,omitempty"`
	Email *string `yaml:"email,omitempty"`
}

// License is a struct to define OAS3's License Data Type
type License struct {
	Name *string `yaml:"name,omitempty"`
	URL  *string `yaml:"url,omitempty"`
}
