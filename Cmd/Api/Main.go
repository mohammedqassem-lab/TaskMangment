package main

import (
	app "TaskMangment/Internal/App"
	"log"
)

func main() {
	app, err := app.New()
	if err != nil {
		log.Fatal(err)
	}

	defer app.Cancel()
	defer app.DB.Close()

	if err := app.Router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
