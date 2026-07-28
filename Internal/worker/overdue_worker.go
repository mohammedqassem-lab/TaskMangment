package worker

import (
	service "TaskMangment/Internal/Service"
	"context"
	"fmt"
	"log"
	"time"
)

type OverDueWorker struct {
	taskService service.TaskService
}

func NewOverDueWorker(taskservice service.TaskService) *OverDueWorker {
	return &OverDueWorker{
		taskService: taskservice,
	}
}
func (w *OverDueWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	fmt.Println("worker starts")
	for {
		select {
		case <-ctx.Done():
			log.Println("worker stopped")
			return
		case <-ticker.C:
			if err := w.taskService.MakeTaskOverDeue(ctx); err != nil {
				log.Printf("failed to update overdue tasks: %v", err)
				continue
			}
		}
	}
}
