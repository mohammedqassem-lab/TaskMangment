package model

type WorkspaceMember struct {
	ID          int64  `json:"id"`
	WorkspaceID int64  `json:"workspace_id"`
	UserID      int64  `json:"user_id"`
	Role        string `json:"role"`
}
