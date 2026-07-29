package dto

type UpdateMemberRoleDto struct {
	UserID  int64  `json:"user_id" binding:"required"`
	Role    string `json:"role" binding:"required"`
	Version int64  `json:"version" binding:"required"`
}
