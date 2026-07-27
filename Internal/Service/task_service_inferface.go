package service

import (
	"TaskMangment/Internal/dto"
	"context"
)

type ITaskService interface {
	Create(ctx context.Context, Task *dto.AddTask) error
	Edit(ctx context.Context, Task *dto.EditTask) error
	Delete(ctx context.Context, id int64) error
}
