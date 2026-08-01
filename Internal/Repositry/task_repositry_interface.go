package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
)

type ITaskRepositry interface {
	Create(ctx context.Context, Task *model.Task) error
	Update(ctx context.Context, Task *dto.EditTask) error
	Delete(ctx context.Context, Id int64, WorkspaceId int64) error

	CheckProject(ctx context.Context, projectId, workspaceId int64) error
	CheckUser(ctx context.Context, Parent_task_id int64, AssigneeId int64, workspace_id int64) error
	GetAll(ctx context.Context, TaskFilter dto.TaskFilter) ([]*model.Task, error)
	GetOverDueTasks(ctx context.Context) ([]*int64, error)
	MakeTaskOverDeue(ctx context.Context, id *int64) error
}
