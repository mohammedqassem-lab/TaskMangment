package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
)

type IWorkspaceMemberRepository interface {
	AddMember(ctx context.Context, workspaceID, userID int64, role string) error
	GetWorkspaceMembers(ctx context.Context, workspaceID int64) ([]*dto.GetWorkspaceMembersDto, error)
	UpdateMemberRole(ctx context.Context, workspaceID, userID int64, role string) error
	DeleteMember(ctx context.Context, workspaceID, userID int64) error
	GetWorkspaceByUserID(ctx context.Context, UserId int64) (*model.Workspace, error)
}
