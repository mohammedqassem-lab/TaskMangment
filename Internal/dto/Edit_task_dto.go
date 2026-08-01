package dto

type EditTask struct {
	Id          int64  `json:"id" binding:"required"`
	Titel       string `json:"titel"`
	Description string `json:"description"`
	Status      string `json:"status" binding:"omitempty,oneof=Todo In Progress Done"`
	Priority    string `json:"priority" binding:"omitempty,oneof=Low Medium High Hihg"`
	Version     int64  `json:"version" binding:"required"`
	UserId      int64  `json:"-" swaggerignore:"true"`
	WorkSpaceId int64  `json:"-" swaggerignore:"true"`
}
