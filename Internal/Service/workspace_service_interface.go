package service

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IWorkspaceService interface {
	Create(ctx context.Context, workspace *model.Workspace) error
	InviteMember(ctx context.Context, workspaceID int64, userID int64, role string) error
}
