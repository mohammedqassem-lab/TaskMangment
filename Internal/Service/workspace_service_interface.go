package service

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IWorkspaceService interface {
	Create(ctx context.Context, workspace *model.Workspace) error
	GetAllWorkspace(ctx context.Context) ([]*model.Workspace, error)
	UpdateWorkspace(ctx context.Context, workspace *model.Workspace) error
	deleteWorkspace(ctx context.Context, workspaceID int64) error
}
