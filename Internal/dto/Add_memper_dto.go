package dto

type AddMemberDto struct {
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}
