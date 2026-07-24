package service

import (
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	"context"
)

type WorkspaceService struct {
	repo repositry.IWorkspaceRepository
}

func NewWorkspaceService(repo repositry.IWorkspaceRepository) *WorkspaceService {
	return &WorkspaceService{
		repo: repo,
	}
}
func (s *WorkspaceService) Create(ctx context.Context, workspace *model.Workspace) error {
	return s.repo.Create(ctx, workspace)
}
