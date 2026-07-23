package main

import (
	database "TaskMangment/Internal/DataBase"
	repositry "TaskMangment/Internal/Repositry"
	service "TaskMangment/Internal/Service"
	handler "TaskMangment/Internal/handler"
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
	repo := repositry.GetNewUserRepositry(db)
	service := service.NewUserService(repo)
	handler := handler.NewUserHandler(*service)
	r := gin.Default()
	route.RegisterUserRoutes(r, handler)
	route.LoginUserRoutes(r, handler)
	r.Run(":8080")
}
