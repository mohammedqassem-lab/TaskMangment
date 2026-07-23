package service

import (
	auth "TaskMangment/Internal/Auth"
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	"context"

	bycrypt "golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo repositry.IUserRepositry
}

func NewUserService(repo repositry.IUserRepositry) *UserService {
	return &UserService{
		repo: repo,
	}
}
func (s *UserService) Register(ctx context.Context, user model.User) error {
	hashpassword, err := bycrypt.GenerateFromPassword([]byte(user.Hashpassword), bycrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Hashpassword = string(hashpassword)
	return s.repo.Create(ctx, &user)
}
func (s *UserService) Login(ctx context.Context, loginDto model.User) (string, error) {
	user, err := s.repo.GetByEmail(ctx, loginDto.Email)
	if err != nil {
		return "", err
	}
	err = bycrypt.CompareHashAndPassword([]byte(user.Hashpassword), []byte(loginDto.Hashpassword))
	if err != nil {
		return "", err
	}
	jwtToken, err := auth.CreateToken(int(user.Id))
	if err != nil {
		return "", err
	}
	return jwtToken, nil
}
