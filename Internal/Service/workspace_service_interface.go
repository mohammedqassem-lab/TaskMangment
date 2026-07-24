package service

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IWorkspaceService interface {
	Create(ctx context.Context, workspace *model.Workspace) error
}
