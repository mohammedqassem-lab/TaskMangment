package app

import (
	cashing "TaskMangment/Internal/Cashing"
	database "TaskMangment/Internal/DataBase"
	middelware "TaskMangment/Internal/Middelware"
	repositry "TaskMangment/Internal/Repositry"
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/handler"
	route "TaskMangment/Internal/route"
	"TaskMangment/Internal/worker"
	"context"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
	Cancel context.CancelFunc
}

func New() (*App, error) {
	db, err := database.ConnectToDb()
	if err != nil {
		return nil, err
	}
	cash := cashing.NewCache(100 * 1024 * 1024)
	// Repositories
	userRepo := repositry.GetNewUserRepositry(db)
	workspaceRepo := repositry.GetNewWorkspaceRepository(db, cash)
	workspaceMemberRepo := repositry.GetNewWorkspaceMemberRepository(db)
	projectRepo := repositry.GetNewProjectRepository(db, cash)
	taskRepo := repositry.GetNewTaskRepository(db, cash)

	// Services
	userService := service.NewUserService(userRepo)
	workspaceService := service.NewWorkspaceService(workspaceRepo)
	workspaceMemberService := service.NewWorkspaceMemberService(workspaceMemberRepo)
	projectService := service.NewProjectService(projectRepo)
	taskService := service.NewTaskService(taskRepo)

	// Handlers
	userHandler := handler.NewUserHandler(*userService)
	workspaceHandler := handler.NewWorkspaceHandler(*workspaceService)
	workspaceMemberHandler := handler.NewWorkspaceMemberHandler(*workspaceMemberService)
	projectHandler := handler.NewProjectHandler(*projectService)
	taskHandler := handler.NewTaskHandler(*taskService)

	// Router
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Use(middelware.ErrorMiddleware())

	authRoutes := r.Group("/")
	authRoutes.Use(middelware.AuthMiddleeare())

	adminRoutes := authRoutes.Group("/")
	adminRoutes.Use(middelware.RequireRole(workspaceRepo, "Admin"))

	adminAndMemberRoutes := authRoutes.Group("/")
	adminAndMemberRoutes.Use(middelware.RequireRole(workspaceRepo, "Admin", "Member"))

	AllRoute := authRoutes.Group("/", middelware.RequireRole(workspaceRepo, "Admin", "Member", "Viewer"))

	// User Routes
	route.LoginUserRoutes(r, userHandler)
	route.RegisterUserRoutes(r, userHandler)
	route.RefreshTokenRoute(r, userHandler)

	// Workspace Routes
	route.CreateWorkspaceRoutes(authRoutes, workspaceHandler)
	route.GetAllWorkspaceRoutes(authRoutes, workspaceHandler)
	route.UpdateWorkspaceRoutes(adminRoutes, workspaceHandler)
	route.DeleteWorkspaceRoutes(adminRoutes, workspaceHandler)

	// Workspace Member Routes
	route.InviteMemberRoutes(adminRoutes, workspaceMemberHandler)
	route.GetWorkspaceMembersRoutes(adminRoutes, workspaceMemberHandler)
	route.UpdateMemberRoleRoutes(adminRoutes, workspaceMemberHandler)
	route.DeleteMemberRoutes(adminRoutes, workspaceMemberHandler)

	// Project Routes
	route.CreateProject(adminAndMemberRoutes, projectHandler)
	route.GetById(AllRoute, projectHandler)
	route.Get(adminAndMemberRoutes, projectHandler)
	route.Update(adminAndMemberRoutes, projectHandler)
	route.Delete(adminAndMemberRoutes, projectHandler)

	// Task Routes
	route.CreateTaskRoute(adminAndMemberRoutes, taskHandler)
	route.EditTaskRoute(adminAndMemberRoutes, taskHandler)
	route.DeleteTaskRoute(adminAndMemberRoutes, taskHandler)
	route.GetAllRoute(AllRoute, taskHandler)

	// Worker
	ctx, cancel := context.WithCancel(context.Background())

	overdueWorker := worker.NewOverDueWorker(*taskService)
	userWorker := worker.NewUserWorker(*userService)
	go overdueWorker.Start(ctx)
	go userWorker.Start(ctx)

	return &App{
		DB:     db,
		Router: r,
		Cancel: cancel,
	}, nil
}
func (a *App) Shutdown() error {

	log.Println("Stopping background workers...")

	a.Cancel()

	log.Println("Closing database connection...")

	return a.DB.Close()
}
