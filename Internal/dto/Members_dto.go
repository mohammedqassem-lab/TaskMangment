package dto

type GetWorkspaceMembersDto struct {
	UserName    string `json:"user_name" binding:"required"`
	Role        string `json:"role" binding:"required"`
	WorkspaceID int64  `json:"workspace_id" binding:"required"`
}
