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
