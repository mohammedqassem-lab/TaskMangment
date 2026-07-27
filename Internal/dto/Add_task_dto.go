package dto

import "time"

type AddTask struct {
	ProjectId      int64     `json:"project_id"`
	Titel          string    `json:"titel"`
	Description    string    `json:"description"`
	AssigneeId     int64     `json:"assigneeid"`
	Parent_task_id int64     `json:"parent_task_id"`
	DueDate        time.Time `json:"due_date"`
	CreateUserId   int64
	WorkSpaceId    int64
}
