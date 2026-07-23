package service

import (
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	bycrypt "golang.org/x/crypto/bcrypt"
	"context"
)

type UserService struct {
	repo repositry.IUserRepositry
}
func NewUserService(repo repositry.IUserRepositry)*UserService{
	return &UserService{
		repo: repo,
	}
}
func(s *UserService) Register(ctx context.Context,user model.User)error{
	hashpassword,err := bycrypt.GenerateFromPassword([]byte(user.Hashpassword), bycrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Hashpassword = string(hashpassword)
	return s.repo.Create(ctx,&user)
}