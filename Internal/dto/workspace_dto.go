package dto

type CreateWorkspaceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
