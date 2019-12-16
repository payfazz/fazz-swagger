package value

// ComponentType is enum to defined type of OAS3's Components
type ComponentType string

const (
	COMPONENT_SCHEMAS         ComponentType = "schemas"
	COMPONENT_RESPONSES       ComponentType = "responses"
	COMPONENT_PARAMETERS      ComponentType = "parameters"
	COMPONENT_EXAMPLES        ComponentType = "examples"
	COMPONENT_REQUESTBODIES   ComponentType = "requestbodies"
	COMPONENT_HEADERS         ComponentType = "headers"
	COMPONENT_SECURITYSCHEMES ComponentType = "securityschemes"
	COMPONENT_LINKS           ComponentType = "links"
	COMPONENT_CALLBACKS       ComponentType = "callbacks"
)
