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
func (s *WorkspaceService) GetAllWorkspace(ctx context.Context) ([]*model.Workspace, error) {
	return s.repo.GetAllWorkspace(ctx)
}
func (s *WorkspaceService) UpdateWorkspace(ctx context.Context, workspace *model.Workspace) error {
	return s.repo.UpdateWorkspace(ctx, workspace)
}
func (s *WorkspaceService) DeleteWorkspace(ctx context.Context, workspaceID int64) error {
	return s.repo.DeleteWorkspace(ctx, workspaceID)
}
