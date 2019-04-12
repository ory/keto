package ladon

type Context map[string]interface{}

const (
	Allow = "allow"
	Deny  = "deny"
)

// Input for checking if a request is allowed or not.
//
// swagger:ignore
type Input struct {
	// Resource is the resource that access is requested to.
	Resource string `json:"resource"`

	// Action is the action that is requested on the resource.
	Action string `json:"action"`

	// Subject is the subject that is requesting access.
	Subject string `json:"subject"`

	// Context is the request's environmental context.
	Context map[string]interface{} `json:"context"`
}
