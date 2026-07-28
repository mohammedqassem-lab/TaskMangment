package dto

type UpdateMemberRoleDto struct {
	UserID  int64  `json:"user_id"`
	Role    string `json:"role"`
	Version int64  `json:"version"`
}
