package main

import (
	"log"

	"github.com/rizqdwan/go-chats-api/config"
	"github.com/rizqdwan/go-chats-api/internal/repositories"
	"github.com/rizqdwan/go-chats-api/internal/services"
	"github.com/rizqdwan/go-chats-api/internal/user"
	"github.com/rizqdwan/go-chats-api/internal/ws"
	"github.com/rizqdwan/go-chats-api/router"
)

func main() {
	dbConn, err := config.NewDatabase()
	if err != nil {
		log.Fatalf("could not initialize database connection: %s", err)
	}

	userRepo := repositories.NewUserRepository(dbConn.GetDB())
	userSvc := services.NewUserService(userRepo)
	userHandler := user.NewHandler(userSvc)

	hub := ws.NewHub()
	wsHandler := ws.NewHandler(hub)
	go hub.Run()

	router.InitRouter(userHandler, wsHandler)
	if err := router.Start("0.0.0.0:8080"); err != nil {
		log.Fatalf("could not start server: %s", err)
	}
}