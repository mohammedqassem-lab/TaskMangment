package repositry

import (
	model "TaskMangment/Internal/Model"
	"context"
	"database/sql"
)

type WorkspaceRepository struct {
	db *sql.DB
}

func GetNewWorkspaceRepository(db *sql.DB) *WorkspaceRepository {
	return &WorkspaceRepository{
		db: db,
	}
}
func (r *WorkspaceRepository) Create(ctx context.Context, workspace *model.Workspace) error {

	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := `
	INSERT INTO workspaces (name, description, owner_id)
	VALUES ($1, $2, $3)
	RETURNING id;
	`

	err = tx.QueryRowContext(
		ctx,
		query,
		workspace.Name,
		workspace.Description,
		workspace.OwnerID,
	).Scan(&workspace.ID)

	if err != nil {
		tx.Rollback()
		return err
	}

	memberQuery := `
	INSERT INTO workspaces_member (workspace_id, user_id, role)
	VALUES ($1, $2, $3);
	`

	_, err = tx.ExecContext(
		ctx,
		memberQuery,
		workspace.ID,
		workspace.OwnerID,
		"Admin",
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
func (r *WorkspaceRepository) GetRole(ctx context.Context, workspaceID, userID int64) (string, error) {
	query := `
	SELECT role FROM workspaces_member
	WHERE workspace_id = $1 AND user_id = $2;
	`
	var role string
	err := r.db.QueryRowContext(ctx, query, workspaceID, userID).Scan(&role)
	if err != nil {
		return "", err
	}
	return role, nil
}

func (r *WorkspaceRepository) GetWorkspaceByUserID(ctx context.Context, UserId int64) (*model.Workspace, error) {
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
func (r *WorkspaceRepository) AddMember(ctx context.Context, workspaceID, userID int64, role string) error {
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
