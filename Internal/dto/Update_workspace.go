package dto

type UpdateWorkspaceDto struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     int64  `json:"version" binding:"required"`
}
