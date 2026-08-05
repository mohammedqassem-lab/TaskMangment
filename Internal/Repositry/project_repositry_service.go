package repositry

import (
	cashing "TaskMangment/Internal/Cashing"
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
)

type ProjectRepository struct {
	db   *sql.DB
	cash *cashing.Cache
}

func GetNewProjectRepository(db *sql.DB, cash *cashing.Cache) *ProjectRepository {
	return &ProjectRepository{
		db:   db,
		cash: cash,
	}
}

func (r *ProjectRepository) Create(ctx context.Context, project *model.Project) error {
	query := `INSERT INTO Project (name, description, created_by, workspace_id)
	VALUES ($1, $2, $3, $4) RETURNING id
	`
	err := r.db.QueryRowContext(ctx, query, project.Name, project.Description, project.CreatedBy, project.WorkspaceId).Scan(&project.Id)
	if err != nil {
		return err
	}
	val, err := json.Marshal(project)
	if err == nil {
		r.cash.Set("Project"+strconv.FormatInt(project.Id, 10), val, 60)
	}
	return nil
}
func (r *ProjectRepository) GetById(ctx context.Context, id int64, workspaceId int64) (*dto.ProjectDto, error) {
	val, err := r.cash.Get("Project" + strconv.FormatInt(id, 10))
	if err == nil {
		var model dto.ProjectDto
		json.Unmarshal(val, &model)
		return &model, nil
	}
	query := `
	SELECT p.id,p.name,w.name,p.description,u.name FROM Project p
	join workspaces w on w.id=p.workspace_id
	join users u on u.id=p.created_by
	WHERE p.id = $1 And p.workspace_id = $2
	`
	var model dto.ProjectDto
	row := r.db.QueryRowContext(ctx, query, id, workspaceId)
	err = row.Scan(&model.Id, &model.Name, &model.WorkspaceName, &model.Description, &model.UserName)
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
	description=$2,
	version=version+1
	where id = $3
	AND workspace_id=$4
	AND version=$5
	`
	result, err := r.db.ExecContext(ctx, query, project.Name, project.Description, project.Id, project.WorkspaceId, project.Version)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return fmt.Errorf("the data was changed")
	}
	r.cash.Delete("Project" + strconv.FormatInt(project.Id, 10))
	return nil
}
func (r *ProjectRepository) Delete(ctx context.Context, id int64, workspaceId int64) error {
	query := `DELETE FROM Project
	WHERE id = $1 And workspace_id=$2`
	result, err := r.db.ExecContext(ctx, query, id, workspaceId)
	if err != nil {
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return fmt.Errorf("Project not found or already deleted")
	}
	r.cash.Delete("Project" + strconv.FormatInt(id, 10))
	return nil
}
