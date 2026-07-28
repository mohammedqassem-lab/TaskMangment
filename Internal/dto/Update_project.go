package dto

type UpdateProject struct {
	Id          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Version     int64  `json:"version"`
}
