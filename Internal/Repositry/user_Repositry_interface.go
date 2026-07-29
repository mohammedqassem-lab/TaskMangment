package repositry

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IUserRepositry interface {
	Create(ctx context.Context, user *model.User) error
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	SaveRefreshToken(ctx context.Context, token model.RefreshToken) error
	ValidateRefreshToken(ctx context.Context, token string) (model.RefreshToken, error)
}
