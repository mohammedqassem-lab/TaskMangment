package repositry

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IWorkspaceRepository interface {
	Create(ctx context.Context, workspace *model.Workspace) error
	GetRole(ctx context.Context, workspaceID, userID int64) (string, error)
	GetWorkspaceByUserID(ctx context.Context, UserId int64) (*model.Workspace, error)
	GetAllWorkspace(ctx context.Context) ([]*model.Workspace, error)
	UpdateWorkspace(ctx context.Context, workspace *model.Workspace) error
	DeleteWorkspace(ctx context.Context, workspaceID int64) error
}
