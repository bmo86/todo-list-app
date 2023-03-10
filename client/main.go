package main

import (
	"log"
	"os"
	"todo-api/client/handlers"
	"todo-api/middleware"
	"todo-api/server"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env") // load file

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	DATABASE_URL := os.Getenv("DATABASE")
	REDIS_URL := os.Getenv("REDIS_URL")

	s, err := server.NewServer(&server.Config{
		PORT:       PORT,
		JWT_SECRET: JWT_SECRET,
		DATABASE:   DATABASE_URL,
		REDIS_URL:  REDIS_URL,
	})

	if err != nil {
		log.Fatal()
	}

	s.Strat(BindRoutes)

}

func BindRoutes(s server.Server, r *gin.Engine) {

	api := r.Group("/api/v1")

	//home
	r.GET("/home", handlers.HandlerHome())
	//usr
	api.Use(middleware.CheckAuthMiddleware(s))
	r.POST("/singup", handlers.HandlerSingUp(s))
	r.POST("/login", handlers.HandlerLogin(s))
	api.GET("/me", handlers.HandlerMe(s))
	//task
	api.POST("/task", handlers.HandlerCreateTask(s))
	api.DELETE("/task/:id", handlers.HandlerDeleteTask(s))
	api.PATCH("/task/:id", handlers.HandlerUpdateTask(s))
	api.GET("/task/:id", handlers.HandlerGetTask(s))
	api.GET("/tasks/:page", handlers.HandlerGetTasks(s))
	r.GET("/ws", handlers.HandlerWs(s))

}
