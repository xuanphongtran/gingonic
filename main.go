package main

import (
	"fmt"
	"gingonic/controller"
	"gingonic/database"
	"gingonic/middleware"
	"gingonic/model"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    loadEnv()
    loadDatabase()
    serveApplication()
}

func loadDatabase() {
    database.Connect()
    database.Database.AutoMigrate(&model.User{})
    database.Database.AutoMigrate(&model.Entry{})
}

func loadEnv() {
    err := godotenv.Load(".env.local")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func serveApplication() {
    router := gin.Default()

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register", controller.Register)
    publicRoutes.POST("/login", controller.Login)

    protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.JWTAuthMiddleware())
    protectedRoutes.POST("/entry", controller.AddEntry)
    protectedRoutes.GET("/entry", controller.GetAllEntries)
    
    router.Run(":8080")
    fmt.Println("Server running on port 8080")
}
