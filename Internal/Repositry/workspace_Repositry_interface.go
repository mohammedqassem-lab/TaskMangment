package repositry

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IWorkspaceRepository interface {
	Create(ctx context.Context, workspace *model.Workspace) error
}