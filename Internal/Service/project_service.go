package service

import (
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	"TaskMangment/Internal/dto"
	"context"
)

type ProjectService struct {
	repo repositry.IProjectRepository
}

func NewProjectService(repo repositry.IProjectRepository) *ProjectService {
	return &ProjectService{
		repo: repo,
	}
}

func (p *ProjectService) Create(ctx context.Context, project *model.Project) error {
	return p.repo.Create(ctx, project)
}
func (p *ProjectService) GetById(ctx context.Context, id int64) (*dto.ProjectDto, error) {
	return p.repo.GetById(ctx, id)
}
func (p *ProjectService) Get(ctx context.Context, workspaceId int64) ([]*dto.ProjectDto, error) {
	return p.repo.Get(ctx, workspaceId)
}
func (p *ProjectService) Update(ctx context.Context, project *dto.UpdateProject) error {
	return p.repo.Update(ctx, project)
}
func (p *ProjectService) Delete(ctx context.Context, id int64) error {
	return p.repo.Delete(ctx, id)
}
