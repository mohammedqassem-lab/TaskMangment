package worker

import (
	service "TaskMangment/Internal/Service"
	"context"
	"fmt"
	"log"
	"time"
)

type TokenWorker struct {
	UserService service.UserService
}

func NewUserWorker(Userservice service.UserService) *TokenWorker {
	return &TokenWorker{
		UserService: Userservice,
	}
}
func (w *TokenWorker) Start(ctx context.Context) {
	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()
	fmt.Println("token worker starts")
	for {
		select {
		case <-ctx.Done():
			log.Println("worker stopped")
			return
		case <-ticker.C:
			if err := w.UserService.MakeRefreshtokenRevoked(ctx); err != nil {
				log.Printf("failed to update overdue tasks: %v", err)
				continue
			}
		}
	}
}
