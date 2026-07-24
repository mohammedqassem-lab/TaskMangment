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

	// Services
	userService := service.NewUserService(userRepo)
	workspaceService := service.NewWorkspaceService(workspaceRepo)

	// Handlers
	userHandler := handler.NewUserHandler(*userService)
	workspaceHandler := handler.NewWorkspaceHandler(*workspaceService)

	// Router
	r := gin.Default()

	r.Use(middelware.ErrorMiddleware())

	authRoutes := r.Group("")
	authRoutes.Use(middelware.AuthMiddleeare())

	AdminRoutes := authRoutes.Group("")
	AdminRoutes.Use(middelware.RequireRole(workspaceRepo, "Admin"))

	route.LoginUserRoutes(r, userHandler)

	route.RegisterUserRoutes(r, userHandler)

	route.CreateWorkspaceRoutes(authRoutes, workspaceHandler)

	route.InviteMemberRoutes(AdminRoutes, workspaceHandler)

	return &App{
		DB:     db,
		Router: r,
	}, nil
}
