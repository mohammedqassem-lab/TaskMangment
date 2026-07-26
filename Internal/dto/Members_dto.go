package dto

type GetWorkspaceMembersDto struct {
	UserName    string `json:"user_name"`
	Role        string `json:"role"`
	WorkspaceID int64  `json:"workspace_id"`
}
