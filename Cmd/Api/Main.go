package main

import (
	database "TaskMangment/Internal/DataBase"
	repositry "TaskMangment/Internal/Repositry"
	service "TaskMangment/Internal/Service"
	handler "TaskMangment/Internal/handler"
	middelware "TaskMangment/Internal/middelware"
	"TaskMangment/Internal/route"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.ConnectToDb()
	if err != nil {
		log.Fatalf("error daring conact to db %v", err.Error())
	}
	defer db.Close()
	Userrepo := repositry.GetNewUserRepositry(db)
	Userservice := service.NewUserService(Userrepo)
	Userhandler := handler.NewUserHandler(*Userservice)
	workspaceRepo := repositry.GetNewWorkspaceRepository(db)
	workspaceService := service.NewWorkspaceService(workspaceRepo)
	workspaceHandler := handler.NewWorkspaceHandler(*workspaceService)
	r := gin.Default()
	r.Use(middelware.ErrorMiddleware())
	authRoutes := r.Group("", middelware.AuthMiddleeare())
	route.RegisterUserRoutes(authRoutes, Userhandler)
	route.LoginUserRoutes(r, Userhandler)
	route.CreateWorkspaceRoutes(authRoutes, workspaceHandler)
	r.Run(":8080")
}
