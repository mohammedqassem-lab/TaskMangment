package service

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
)

type IprojectService interface {
	create(ctx context.Context, project *model.Project) error
	GetById(ctx context.Context, id int64) (*model.Project, error)
	Get(ctx context.Context, workspaceId int64) ([]*dto.ProjectDto, error)
	Update(ctx context.Context, project *dto.UpdateProject) error
	Delete(ctx context.Context, id int64) error
}
