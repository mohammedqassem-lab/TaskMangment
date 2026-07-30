package dto

import "time"

type AddTask struct {
	ProjectId      int64     `json:"project_id" binding:"required"`
	Titel          string    `json:"titel" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	AssigneeId     int64     `json:"assigneeid" binding:"required"`
	Parent_task_id int64     `json:"parent_task_id"`
	DueDate        time.Time `json:"due_date" binding:"required" time_format:"2006-01-02"`
	CreateUserId   int64     `json:"-" swaggerignore:"true"`
	WorkSpaceId    int64     `json:"-" swaggerignore:"true"`
}
