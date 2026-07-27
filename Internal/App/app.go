package app

import (
	database "TaskMangment/Internal/DataBase"
	middelware "TaskMangment/Internal/Middelware"
	repositry "TaskMangment/Internal/Repositry"
	service "TaskMangment/Internal/Service"
	"TaskMangment/Internal/handler"
	"TaskMangment/Internal/route"
	"database/sql"

	"github.com/gin-gonic/gin"
)

type App struct {
	DB     *sql.DB
	Router *gin.Engine
}

func New() (*App, error) {

	db, err := database.ConnectToDb()
	if err != nil {
		return nil, err
	}

	// Repositories
	userRepo := repositry.GetNewUserRepositry(db)
	workspaceRepo := repositry.GetNewWorkspaceRepository(db)
	WorkspaceMemberRepo := repositry.GetNewWorkspaceMemberRepository(db)
	ProjectRepository := repositry.GetNewProjectRepository(db)
	TaskRepositry := repositry.GetNewTaskRepository(db)

	// Services
	userService := service.NewUserService(userRepo)
	workspaceService := service.NewWorkspaceService(workspaceRepo)
	workspaceMemberService := service.NewWorkspaceMemberService(WorkspaceMemberRepo)
	ProjectService := service.NewProjectService(ProjectRepository)
	TaskService := service.NewTaskService(TaskRepositry)

	// Handlers
	userHandler := handler.NewUserHandler(*userService)
	workspaceHandler := handler.NewWorkspaceHandler(*workspaceService)
	workspaceMemberHandler := handler.NewWorkspaceMemberHandler(*workspaceMemberService)
	projectHandler := handler.NewProjectHandler(*ProjectService)
	taskHandler := handler.NewTaskHandler(*TaskService)

	// Router
	r := gin.Default()

	r.Use(middelware.ErrorMiddleware())

	authRoutes := r.Group("")
	authRoutes.Use(middelware.AuthMiddleeare())

	AdminRoutes := authRoutes.Group("")
	AdminRoutes.Use(middelware.RequireRole(workspaceRepo, "Admin"))

	adminAndMemperRoutes := authRoutes.Group("")
	adminAndMemperRoutes.Use(middelware.RequireRole(workspaceRepo, "Admin", "Member"))
	//Account Routes
	route.LoginUserRoutes(r, userHandler)

	route.RegisterUserRoutes(r, userHandler)
	//Workspace Routes
	route.CreateWorkspaceRoutes(authRoutes, workspaceHandler)

	route.GetAllWorkspaceRoutes(authRoutes, workspaceHandler)

	route.UpdateWorkspaceRoutes(AdminRoutes, workspaceHandler)

	route.DeleteWorkspaceRoutes(AdminRoutes, workspaceHandler)
	//Workspace Member Routes
	route.InviteMemberRoutes(AdminRoutes, workspaceMemberHandler)

	route.GetWorkspaceMembersRoutes(AdminRoutes, workspaceMemberHandler)

	route.UpdateMemberRoleRoutes(AdminRoutes, workspaceMemberHandler)

	route.DeleteMemberRoutes(AdminRoutes, workspaceMemberHandler)
	//project route

	route.CreateProject(adminAndMemperRoutes, projectHandler)

	route.GetById(authRoutes, projectHandler)

	route.Get(adminAndMemperRoutes, projectHandler)

	route.Update(adminAndMemperRoutes, projectHandler)

	route.Delete(adminAndMemperRoutes, projectHandler)
	//Task route

	route.CreateTaskRoute(adminAndMemperRoutes, taskHandler)

	route.EditTaskRoute(adminAndMemperRoutes, taskHandler)

	route.DeleteTaskRoute(adminAndMemperRoutes, taskHandler)
	return &App{
		DB:     db,
		Router: r,
	}, nil
}
