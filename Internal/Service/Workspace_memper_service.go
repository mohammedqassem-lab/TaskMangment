package service

import (
	repositry "TaskMangment/Internal/Repositry"
	"TaskMangment/Internal/dto"
	"context"
	"fmt"
)

type WorkspaceMemberService struct {
	repo repositry.IWorkspaceMemberRepository
}

func NewWorkspaceMemberService(repo repositry.IWorkspaceMemberRepository) *WorkspaceMemberService {
	return &WorkspaceMemberService{
		repo: repo,
	}
}
func (s *WorkspaceMemberService) InviteMember(ctx context.Context, workspaceID int64, userID int64, role string) error {
	workspace, err := s.repo.GetWorkspaceByUserID(ctx, userID)
	if err == nil {
		return fmt.Errorf("user is already a member of the workspace")
	}
	if workspace != nil {
		return fmt.Errorf("user is already a member of the workspace")
	}
	return s.repo.AddMember(ctx, workspaceID, userID, role)
}
func (s *WorkspaceMemberService) GetWorkspaceMembers(ctx context.Context, workspaceID int64) ([]*dto.GetWorkspaceMembersDto, error) {
	return s.repo.GetWorkspaceMembers(ctx, workspaceID)
}
func (s *WorkspaceMemberService) UpdateMemberRole(ctx context.Context, workspaceID, userID int64, role string) error {
	return s.repo.UpdateMemberRole(ctx, workspaceID, userID, role)
}
func (s *WorkspaceMemberService) DeleteMember(ctx context.Context, workspaceID, userID int64) error {
	return s.repo.DeleteMember(ctx, workspaceID, userID)
}
