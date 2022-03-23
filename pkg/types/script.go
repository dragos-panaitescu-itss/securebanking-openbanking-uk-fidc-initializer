package types

// RequestScript model
type RequestScript struct {
	ID          string `json:"_id,omitempty"`
	Name        string `json:"name"`
	Script      string `json:"script"`
	Language    string `json:"language"`
	Context     string `json:"context"`
	Description string `json:"description"`
	CreatedBy   string `json:"createdBy,omitempty"`
}
