package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
	"database/sql"
)

type ProjectRepository struct {
	db *sql.DB
}

func GetNewProjectRepository(db *sql.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (r *ProjectRepository) Create(ctx context.Context, project *model.Project) error {
	query := `INSERT INTO Project (name, description, created_by, workspace_id)
	VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.ExecContext(ctx, query, project.Name, project.Description, project.CreatedBy, project.WorkspaceId)
	if err != nil {
		return err
	}
	return nil
}
func (r *ProjectRepository) GetById(ctx context.Context, id int64) (*dto.ProjectDto, error) {
	query := `
	SELECT p.id,p.name,w.name,p.description,u.name FROM Project p
	join workspaces w on w.id=p.workspace_id
	join users u on u.id=p.created_by
	WHERE p.id = $1
	`
	var model dto.ProjectDto
	row := r.db.QueryRowContext(ctx, query, id)
	err := row.Scan(&model.Id, &model.Name, &model.WorkspaceName, &model.Description, &model.UserName)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (r *ProjectRepository) Get(ctx context.Context, workspaceId int64) ([]*dto.ProjectDto, error) {
	query := `
	SELECT p.id,p.name,w.name,p.description,u.name FROM Project p
	join workspaces w on w.id=p.workspace_id
	join users u on u.id=p.created_by
	WHERE p.workspace_id = $1
	`
	var model []*dto.ProjectDto
	rows, err := r.db.QueryContext(ctx, query, workspaceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var project dto.ProjectDto
		err := rows.Scan(&project.Id, &project.Name, &project.WorkspaceName, &project.Description, &project.UserName)
		if err != nil {
			return nil, err
		}
		if err := rows.Err(); err != nil {
			return nil, err
		}
		model = append(model, &project)
	}

	return model, nil
}
func (r *ProjectRepository) Update(ctx context.Context, project *dto.UpdateProject) error {
	query := `UPDATE Project set
	name =$1,
	description=$2
	where id = $3
	`
	_, err := r.db.ExecContext(ctx, query, project.Name, project.Description, project.Id)
	if err != nil {
		return err
	}
	return nil
}
func (r *ProjectRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM Project
	WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	return nil
}
