package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
)

type ITaskRepositry interface {
	Create(ctx context.Context, Task *dto.AddTask) error
	Update(ctx context.Context, Task *dto.EditTask) error
	Delete(ctx context.Context, Id int64) error

	CheckProject(ctx context.Context, projectId, workspaceId int64) error
	CheckUser(ctx context.Context, Parent_task_id int64, AssigneeId int64, workspace_id int64) error
	GetAll(ctx context.Context, TaskFilter dto.TaskFilter) ([]*model.Task, error)
}
