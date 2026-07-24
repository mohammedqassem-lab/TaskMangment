package dto

type CreateWorkspaceRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	OwnerId     int    `json:"owner_id"`
}
