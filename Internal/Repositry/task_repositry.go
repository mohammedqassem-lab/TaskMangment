package repositry

import (
	model "TaskMangment/Internal/Model"
	"TaskMangment/Internal/dto"
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
)

type TaskRepository struct {
	db *sql.DB
}

func GetNewTaskRepository(db *sql.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}
func BuildQuery(taskFilter dto.TaskFilter) (string, []any) {
	query := `
SELECT
	id,
	project_id,
	title,
	description,
	status,
	priority,
	parent_task_id,
	assignee_id,
	created_by,
	due_date,
	created_at,
	updated_at,
	version
FROM task
WHERE 1 = 1`

	args := []any{}
	index := 1

	if taskFilter.ProjectId != 0 {
		query += " AND project_id = $" + strconv.Itoa(index)
		args = append(args, taskFilter.ProjectId)
		index++
	}

	if taskFilter.AssigneeId != 0 {
		query += " AND assignee_id = $" + strconv.Itoa(index)
		args = append(args, taskFilter.AssigneeId)
		index++
	}

	if taskFilter.Priorty != "" {
		query += " AND priority = $" + strconv.Itoa(index)
		args = append(args, taskFilter.Priorty)
		index++
	}

	if taskFilter.Status != "" {
		query += " AND status =$" + strconv.Itoa(index)
		args = append(args, taskFilter.Status)
		index++
	}

	if taskFilter.Serch != "" {
		query += " AND (title ILIKE $" + strconv.Itoa(index) +
			" OR description ILIKE $" + strconv.Itoa(index) + ")"

		args = append(args, "%"+taskFilter.Serch+"%")
		index++
	}

	// Sorting
	sortColumn := "created_at"

	switch taskFilter.SortBy {
	case "priority":
		sortColumn = "priority"
	case "title":
		sortColumn = "title"
	case "due_date":
		sortColumn = "due_date"
	}

	order := "ASC"
	if strings.EqualFold(taskFilter.Order, "DESC") {
		order = "DESC"
	}

	query += " ORDER BY " + sortColumn + " " + order

	// Pagination
	if taskFilter.Limit > 0 {
		query += " LIMIT $" + strconv.Itoa(index)
		args = append(args, taskFilter.Limit)
		index++
	}

	if taskFilter.Offset > 0 {
		query += " OFFSET $" + strconv.Itoa(index)
		args = append(args, taskFilter.Offset)
		index++
	}

	return query, args
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
func (r *TaskRepository) Create(ctx context.Context, Task *model.Task) error {
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
	err = tx.QueryRowContext(ctx, query, Task.Titel, Task.Description, Task.AssigneeId, Task.Parent_task_id, Task.Due_date, Task.ProjectId, Task.CreatedBy).Scan(&id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `INSERT INTO TaskHistory (task_id, action, changed_by)
	VALUES ($1, $2,$3);`
	_, err = tx.ExecContext(ctx, query, id, "create", Task.CreatedBy)
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
	query := `select id,title,description,status,priority,version from task
	WHERE id = $1 AND version=$2
	`
	err = tx.QueryRowContext(ctx, query, Task.Id, Task.Version).Scan(&task.Id, &task.Titel, &task.Description, &task.Status, &task.Priority, &task.Version)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = `
	UPDATE task
	SET title = $1, description = $2, 
	status = $3, priority = $4,version=version+1
	WHERE id = $5 AND version=$6
	`
	if Task.Status == "" {
		Task.Status = task.Status
	}
	if Task.Priority == "" {
		Task.Priority = task.Priority
	}
	result, err := tx.ExecContext(ctx, query, Task.Titel, Task.Description, Task.Status, Task.Priority, Task.Id, Task.Version)
	if err != nil {
		tx.Rollback()
		return err
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		tx.Rollback()
		return fmt.Errorf("tha data was changed")
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
func (r *TaskRepository) GetAll(ctx context.Context, TaskFilter dto.TaskFilter) ([]*model.Task, error) {
	qeury, args := BuildQuery(TaskFilter)
	rows, err := r.db.QueryContext(ctx, qeury, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.Id, &task.ProjectId, &task.Titel, &task.Description, &task.Status, &task.Priority, &task.Parent_task_id,
			&task.AssigneeId, &task.CreatedBy, &task.Due_date, &task.CreatedAt, &task.UpdatedAt, &task.Version)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
func (r *TaskRepository) GetOverDueTasks(ctx context.Context) ([]*int64, error) {
	query := `SELECT id from task
	where due_date < now()
	AND status!='Done'
	AND status!='Overdue'`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []*int64
	for rows.Next() {
		var task int64
		err := rows.Scan(&task)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
func (r *TaskRepository) MakeTaskOverDeue(ctx context.Context, id *int64) error {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	query := `UPDATE task
	SET status = 'Overdue',version=version+1,updated_at=now()
	WHERE id = $1`
	_, err = tx.ExecContext(ctx, query, id)
	if err != nil {
		tx.Rollback()
		return err
	}
	query = `INSERT INTO TaskHistory (task_id, action, changed_by, field_name, old_value, new_value)
	VALUES ($1, $2,$3,$4,$5,$6);`
	_, err = tx.ExecContext(ctx, query, id, "update", 1, "status", "Todo or InProgress", "Overdue")
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
