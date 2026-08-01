package dto

type AddMemberDto struct {
	UserID int64  `json:"user_id" binding:"required"`
	Role   string `json:"role" binding:"required,oneof=Admin Member Viewer"`
}
