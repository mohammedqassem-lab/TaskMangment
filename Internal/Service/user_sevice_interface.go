package service

import (
	model "TaskMangment/Internal/Model"
	"context"
)

type IUserService interface {
	Regester(ctx context.Context,user *model.User) error
}