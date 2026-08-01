package service

import (
	model "TaskMangment/Internal/Model"
	repositry "TaskMangment/Internal/Repositry"
	"TaskMangment/Internal/dto"
	"context"
	"fmt"
	"time"
)

type TaskService struct {
	repo repositry.ITaskRepositry
}

func NewTaskService(repo repositry.ITaskRepositry) *TaskService {
	return &TaskService{
		repo: repo,
	}
}

func (t *TaskService) Create(ctx context.Context, Task *dto.AddTask) error {
	if Task.DueDate.Before(time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local)) {
		return fmt.Errorf("duedate can tot befoer today")
	}
	model := model.Task{
		Titel:          Task.Titel,
		Description:    Task.Description,
		ProjectId:      Task.ProjectId,
		AssigneeId:     Task.AssigneeId,
		Parent_task_id: Task.Parent_task_id,
		Due_date:       Task.DueDate,
		CreatedBy:      Task.CreateUserId,
	}
	err := t.repo.CheckProject(ctx, Task.ProjectId, Task.WorkSpaceId)
	if err != nil {
		return err
	}
	err = t.repo.CheckUser(ctx, Task.Parent_task_id, Task.AssigneeId, Task.WorkSpaceId)
	if err != nil {
		return err
	}
	err = t.repo.Create(ctx, &model)
	if err != nil {
		return err
	}
	return nil
}
func (t *TaskService) Edit(ctx context.Context, Task *dto.EditTask) error {
	return t.repo.Update(ctx, Task)
}
func (t *TaskService) Delete(ctx context.Context, id, WorkspaceId int64) error {
	return t.repo.Delete(ctx, id, WorkspaceId)
}
func (t *TaskService) GetAll(ctx context.Context, FilterTask dto.TaskFilter) ([]*model.Task, error) {
	return t.repo.GetAll(ctx, FilterTask)
}
func (t *TaskService) MakeTaskOverDeue(ctx context.Context) error {
	ids, err := t.repo.GetOverDueTasks(ctx)
	if err != nil {
		return err
	}

	for _, id := range ids {
		if err := t.repo.MakeTaskOverDeue(ctx, id); err != nil {
			return err
		}
	}

	return nil
}
