package service

import (
	"TaskMangment/Internal/dto"
	"context"
)

type IWorkspaceMemberService interface {
	InviteMember(ctx context.Context, workspaceID int64, userID int64, role string) error
	GetWorkspaceMembers(ctx context.Context, workspaceID int64) ([]*dto.GetWorkspaceMembersDto, error)
	UpdateMemberRole(ctx context.Context, workspaceID int64, member dto.UpdateMemberRoleDto) error
	DeleteMember(ctx context.Context, workspaceID, userID int64) error
}
