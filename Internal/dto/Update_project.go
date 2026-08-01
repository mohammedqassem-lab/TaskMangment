package dto

type UpdateProject struct {
	Id          int64  `json:"id" binding:"required"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     int64  `json:"version" binding:"required"`
	WorkspaceId int64  `json:"-" swaggerignore:"true"`
}
