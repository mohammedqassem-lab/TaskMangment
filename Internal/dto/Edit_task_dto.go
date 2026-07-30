package dto

type EditTask struct {
	Id          int64  `json:"id" binding:"required"`
	Titel       string `json:"titel"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priorty"`
	Version     int64  `json:"version" binding:"required"`
	UserId      int64  `json:"-" swaggerignore:"true"`
}
