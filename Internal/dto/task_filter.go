package dto

type TaskFilter struct {
	ProjectId   int64
	Status      string
	Priorty     string
	AssigneeId  int64
	Serch       string
	SortBy      string
	Order       string
	Limit       int64
	Offset      int64
	WorkSpaceId int64 `json:"-" swaggerignore:"true"`
}
