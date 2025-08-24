package main

import (
	todo "awesomeProject"
	"awesomeProject/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)

	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("server run error: %v", err.Error())
	}
}
