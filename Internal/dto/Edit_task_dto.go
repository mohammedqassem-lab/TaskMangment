package dto

type EditTask struct {
	Id          int64  `json:"id" binding:"required"`
	Titel       string `json:"titel" binding:"required"`
	Description string `json:"description" binding:"required"`
	Status      string `json:"status" binding:"required"`
	Priority    string `json:"priorty" binding:"required"`
	Version     int64  `json:"version" binding:"required"`
	UserId      int64
}
