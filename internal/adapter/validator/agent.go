package validator

type AgentAddRequest struct {
	Name        string `json:"name" validate:"required"`
	Age         int    `json:"age" validate:"required"`
	Role        string `json:"role" validate:"required"`
	Language    string `json:"language" validate:"required"`
	Description string `json:"description" validate:"omitempty"`
	Status      string `json:"status" validate:"required"`
}

type AgentRemoveRequest struct {
	ID uint `json:"id" validate:"required"`
}
