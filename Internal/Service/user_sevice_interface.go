package service

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IUserService interface {
	Regester(ctx context.Context, user *model.User) error
	login(ctx context.Context, user *model.User) (string, error)
	RefreshToken(ctx context.Context, token string) (model.RefreshToken, error)
}
