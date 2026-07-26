package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
	"database/sql"
)

type WorkspaceMemberRepository struct {
	db *sql.DB
}

func GetNewWorkspaceMemberRepository(db *sql.DB) *WorkspaceMemberRepository {
	return &WorkspaceMemberRepository{
		db: db,
	}
}

func (r *WorkspaceMemberRepository) AddMember(ctx context.Context, workspaceID, userID int64, role string) error {
	query := `
	INSERT INTO workspaces_member (workspace_id, user_id, role)
	VALUES ($1, $2, $3);
	`
	_, err := r.db.ExecContext(ctx, query, workspaceID, userID, role)
	if err != nil {
		return err
	}
	return nil
}
func (r *WorkspaceMemberRepository) GetWorkspaceMembers(ctx context.Context, workspaceID int64) ([]*dto.GetWorkspaceMembersDto, error) {
	query := `SELECT u.name, w.workspace_id, w.role
	FROM workspaces_member w
	join users u on u.id = w.user_id
	WHERE workspace_id = $1;`
	rows, err := r.db.QueryContext(ctx, query, workspaceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*dto.GetWorkspaceMembersDto
	for rows.Next() {
		var member dto.GetWorkspaceMembersDto
		err := rows.Scan(&member.UserName, &member.WorkspaceID, &member.Role)
		if err != nil {
			return nil, err
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		members = append(members, &member)
	}

	return members, nil
}
func (r *WorkspaceMemberRepository) UpdateMemberRole(ctx context.Context, workspaceID, userID int64, role string) error {
	query := `
	UPDATE workspaces_member
	SET role = $1
	WHERE workspace_id = $2 AND user_id = $3;
	`
	_, err := r.db.ExecContext(ctx, query, role, workspaceID, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *WorkspaceMemberRepository) DeleteMember(ctx context.Context, workspaceID, userID int64) error {
	query := `
	DELETE FROM workspaces_member
	WHERE workspace_id = $1 AND user_id = $2;
	`
	_, err := r.db.ExecContext(ctx, query, workspaceID, userID)
	if err != nil {
		return err
	}
	return nil
}
func (r *WorkspaceMemberRepository) GetWorkspaceByUserID(ctx context.Context, UserId int64) (*model.Workspace, error) {
	query := `
	SELECT w.id, w.name, w.description, w.owner_id
	FROM workspaces w
	JOIN workspaces_member wm ON w.id = wm.workspace_id
	WHERE wm.user_id = $1;
	`
	var workspace model.Workspace
	err := r.db.QueryRowContext(ctx, query, UserId).Scan(&workspace.ID, &workspace.Name, &workspace.Description, &workspace.OwnerID)
	if err != nil {
		return nil, err
	}
	return &workspace, nil
}
