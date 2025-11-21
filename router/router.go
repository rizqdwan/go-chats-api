package router

import (
	"time"

	"github.com/rizqdwan/go-chats-api/internal/user"
	"github.com/rizqdwan/go-chats-api/internal/ws"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	router = gin.Default()

	// setup cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST"},
		AllowHeaders: []string{"Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func (origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	router.POST("/api/users/signup", userHandler.CreateUser)
	router.POST("/api/users/signin", userHandler.Login)
	router.GET("/api/users/logout", userHandler.Logout)

	router.POST("/ws/create-room", wsHandler.CreateRoom)
	router.GET("/ws/join-room/:roomId", wsHandler.JoinRoom)
	router.GET("/ws/get-rooms", wsHandler.GetRooms)
	router.GET("/ws/get-clients/:roomId", wsHandler.GetClients)
}

func Start(addr string)	error {
	return router.Run(addr)
}