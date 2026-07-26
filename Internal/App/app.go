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

	// Services
	userService := service.NewUserService(userRepo)
	workspaceService := service.NewWorkspaceService(workspaceRepo)
	workspaceMemberService := service.NewWorkspaceMemberService(WorkspaceMemberRepo)

	// Handlers
	userHandler := handler.NewUserHandler(*userService)
	workspaceHandler := handler.NewWorkspaceHandler(*workspaceService)
	workspaceMemberHandler := handler.NewWorkspaceMemberHandler(*workspaceMemberService)

	// Router
	r := gin.Default()

	r.Use(middelware.ErrorMiddleware())

	authRoutes := r.Group("")
	authRoutes.Use(middelware.AuthMiddleeare())

	AdminRoutes := authRoutes.Group("")
	AdminRoutes.Use(middelware.RequireRole(workspaceRepo, "Admin"))
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

	return &App{
		DB:     db,
		Router: r,
	}, nil
}
