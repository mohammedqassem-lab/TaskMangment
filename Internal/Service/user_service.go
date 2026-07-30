package service

import (
	auth "TaskMangment/Internal/Auth"
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	"TaskMangment/Internal/dto"
	"context"
	"fmt"
	"time"

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
func (s *UserService) Login(ctx context.Context, loginDto model.User) (dto.TokenResponce, error) {
	user, err := s.repo.GetByEmail(ctx, loginDto.Email)
	var respose dto.TokenResponce
	if err != nil {
		return respose, err
	}
	err = bycrypt.CompareHashAndPassword([]byte(user.Hashpassword), []byte(loginDto.Hashpassword))
	if err != nil {
		return respose, err
	}
	respose.Accsestoken, err = auth.CreateToken(int(user.Id))
	if err != nil {
		return respose, err
	}
	respose.RefreshToken, err = auth.GenerateRefreshToken()
	if err != nil {
		return respose, err
	}
	refreshToken := model.RefreshToken{
		Token:     respose.RefreshToken,
		UserId:    user.Id,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
	err = s.repo.SaveRefreshToken(ctx, refreshToken)
	return respose, nil
}
func (s *UserService) RefreshToken(ctx context.Context, token string) (dto.TokenResponce, error) {
	var t dto.TokenResponce
	refreshToken, err := s.repo.ValidateRefreshToken(ctx, token)
	if err != nil {
		return t, err
	}
	if refreshToken.ExpiresAt.Before(time.Now()) {
		return t, fmt.Errorf("the token is expired")
	}

	t.Accsestoken, err = auth.CreateToken(int(refreshToken.UserId))
	t.RefreshToken, err = auth.GenerateRefreshToken()
	model := model.RefreshToken{
		UserId:    refreshToken.UserId,
		ExpiresAt: time.Now().Add(24 * time.Hour),
		Token:     t.RefreshToken,
	}
	err = s.repo.SaveRefreshToken(ctx, model)
	return t, err
}
func (s *UserService) MakeRefreshtokenRevoked(ctx context.Context) error {
	ids, err := s.repo.GetRevokedToken(ctx)
	if err != nil {
		return err
	}
	for _, id := range ids {
		if err := s.repo.MakerevokedTrue(ctx, *id); err != nil {
			return err
		}
	}
	return nil
}
