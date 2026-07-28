package dto

type EditTask struct {
	Id          int64  `json:"id"`
	Titel       string `json:"titel"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Priority    string `json:"priorty"`
	Version     int64  `json:"version"`
	UserId      int64
}
