package service

import (
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	"context"
	"fmt"
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
func (s *WorkspaceService) InviteMember(ctx context.Context, workspaceID int64, userID int64, role string) error {
	workspace, err := s.repo.GetWorkspaceByUserID(ctx, userID)
	if err == nil {
		return fmt.Errorf("user is already a member of the workspace")
	}
	if workspace != nil {
		return fmt.Errorf("user is already a member of the workspace")
	}
	return s.repo.AddMember(ctx, workspaceID, userID, role)
}
