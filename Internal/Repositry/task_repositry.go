package repositry

import (
	"TaskMangment/Internal/dto"
	"context"
	"database/sql"
	"fmt"
)

type TaskRepository struct {
	db *sql.DB
}

func GetNewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}
func (r *TaskRepository) CheckProject(ctx context.Context, projectId, workspaceId int64) error {
	query := `
	select Id,workspace_id from project where Id=$1`

	id := 0
	var workspace_id int64
	workspace_id = 0
	err := r.db.QueryRowContext(ctx, query, projectId).Scan(&id, &workspace_id)
	if err != nil {
		return err
	}
	if id == 0 {
		return fmt.Errorf("the project id is not valid")
	}
	if workspace_id == 0 || workspace_id != workspaceId {
		return fmt.Errorf("the project is not in this workspace")
	}
	return nil
}
func (r *TaskRepository) CheckUser(ctx context.Context, Parent_task_id int64, AssigneeId int64, workspace_id int64) error {
	query := `select * from workspaces_member
	where user_id = $1 and workspace_id=$2`
	row := r.db.QueryRowContext(ctx, query, AssigneeId, workspace_id)
	if err := row.Err(); err != nil {
		return err
	}
	if Parent_task_id != 0 {
		query = `select * from task
		where id=$1`
		row = r.db.QueryRowContext(ctx, query, Parent_task_id)
		if err := row.Err(); err != nil {
			return err
		}
	}
	return nil
}
func (r *TaskRepository) Create(ctx context.Context, Task *dto.AddTask) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO task (title, description, assignee_id,parent_task_id,due_date,project_id,created_by)
	VALUES ($1, $2, $3,$4,$5,$6,$7) RETURNING id;
	`
	var id int64
	id = 0
	err = tx.QueryRowContext(ctx, query, Task.Titel, Task.Description, Task.AssigneeId, Task.Parent_task_id, Task.DueDate, Task.ProjectId, Task.CreateUserId).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `INSERT INTO TaskHistory (task_id, action, changed_by)
	VALUES ($1, $2,$3);`
	_, err = tx.ExecContext(ctx, query, id, "create", Task.AssigneeId)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (r *TaskRepository) Update(ctx context.Context, Task *dto.EditTask) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	var task dto.EditTask
	query := `select id,title,description,status,priority from task
	WHERE id = $1
	`
	err = tx.QueryRowContext(ctx, query, Task.Id).Scan(&task.Id, &task.Titel, &task.Description, &task.Status, &task.Priority)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `
	UPDATE task
	SET title = $1, description = $2, 
	status = $3, priority = $4
	WHERE id = $5
	`
	if Task.Status == "" {
		Task.Status = task.Status
	}
	if Task.Priority == "" {
		Task.Priority = task.Priority
	}
	_, err = tx.ExecContext(ctx, query, Task.Titel, Task.Description, Task.Status, Task.Priority, Task.Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	if Task.Titel != "" {
		query = `INSERT INTO TaskHistory (task_id, action, changed_by, field_name, old_value, new_value)
	VALUES ($1, $2,$3,$4,$5,$6);`
		_, err = tx.ExecContext(ctx, query, Task.Id, "update", Task.UserId, "Title", task.Titel, Task.Titel)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if Task.Description != "" {
		query = `INSERT INTO TaskHistory (task_id, action, changed_by, field_name, old_value, new_value)
	VALUES ($1, $2,$3,$4,$5,$6);`
		_, err = tx.ExecContext(ctx, query, Task.Id, "update", Task.UserId, "Description", task.Description, Task.Description)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if Task.Status != "" && Task.Status != task.Status {
		query = `INSERT INTO TaskHistory (task_id, action, changed_by, field_name, old_value, new_value)
	VALUES ($1, $2,$3,$4,$5,$6);`
		_, err = tx.ExecContext(ctx, query, Task.Id, "update", Task.UserId, "Status", task.Status, Task.Status)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	if Task.Priority != "" && Task.Priority != task.Priority {
		query = `INSERT INTO TaskHistory (task_id, action, changed_by, field_name, old_value, new_value)
	VALUES ($1, $2,$3,$4,$5,$6);`
		_, err = tx.ExecContext(ctx, query, Task.Id, "update", Task.UserId, "Priority", task.Priority, Task.Priority)
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	tx.Commit()
	return nil
}
func (r *TaskRepository) Delete(ctx context.Context, Id int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := `
	DELETE FROM task
	WHERE id = $1
	`
	_, err = tx.ExecContext(ctx, query, Id)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `
	DELETE FROM task
	WHERE parent_task_id = $1
	`
	_, err = tx.ExecContext(ctx, query, Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
