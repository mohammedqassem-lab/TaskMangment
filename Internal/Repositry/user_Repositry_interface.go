package repositry

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IUserRepositry interface {
	Create(ctx context.Context, user *model.User) error
}
