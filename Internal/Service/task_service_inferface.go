package service

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
)

type ITaskService interface {
	Create(ctx context.Context, Task *dto.AddTask) error
	Edit(ctx context.Context, Task *dto.EditTask) error
	Delete(ctx context.Context, id, WorkspaceId int64) error
	GetAll(ctx context.Context, FilterTask dto.TaskFilter) ([]*model.Task, error)
	MakeTaskOverDeue(ctx context.Context) error
}
