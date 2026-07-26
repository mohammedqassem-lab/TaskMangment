package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
)

type IProjectRepository interface {
	Create(ctx context.Context, project *model.Project) error
	GetById(ctx context.Context, id int64) (*dto.ProjectDto, error)
	Get(ctx context.Context, workspaceId int64) ([]*dto.ProjectDto, error)
	Update(ctx context.Context, project *dto.UpdateProject) error
	Delete(ctx context.Context, id int64) error
}
