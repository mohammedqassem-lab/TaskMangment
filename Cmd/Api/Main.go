package main

import (
	app "TaskMangment/Internal/App"
	"fmt"
	"log"
)

func main() {
	application, err := app.New()
	if err != nil {
		log.Fatalf("error during application startup: %v", err)
	}

	defer application.DB.Close()

	if err := application.Router.Run(":8080"); err != nil {
		fmt.Print(err.Error())
	}
}
