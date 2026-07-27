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
	err := t.repo.CheckProject(ctx, Task.ProjectId, Task.WorkSpaceId)
	if err != nil {
		return err
	}
	err = t.repo.CheckUser(ctx, Task.Parent_task_id, Task.AssigneeId, Task.WorkSpaceId)
	if err != nil {
		return err
	}
	err = t.repo.Create(ctx, Task)
	if err != nil {
		return err
	}
	return nil
}
func (t *TaskService) Edit(ctx context.Context, Task *dto.EditTask) error {
	return t.repo.Update(ctx, Task)
}
func (t *TaskService) Delete(ctx context.Context, id int64) error {
	return t.repo.Delete(ctx, id)
}
func (t *TaskService) GetAll(ctx context.Context, FilterTask dto.TaskFilter) ([]*model.Task, error) {
	return t.repo.GetAll(ctx, FilterTask)
}
