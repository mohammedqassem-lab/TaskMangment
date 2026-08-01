package dto

type RefreshTokenDto struct {
	Token string `json:"token" binding:"required"`
}
