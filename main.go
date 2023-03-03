package main

import (
	"InternBCC/Handler"
	"InternBCC/database"
	"InternBCC/middleware"
	"InternBCC/model"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	db := database.InitDB()
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to Migrate")
	}
	if err != nil {
		log.Fatalln("failed to load env file")
	}
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"ping": "pong",
		})
	})

	model.TagDummy()
	model.GDummy()
	v0 := r.Group("/v0")
	v1 := r.Group("/v1")
	v0.POST("/register", Handler.Register)
	v1.POST("/login", Handler.LogIn)
	v1.GET("/logout", Handler.Logoout)
	v1.GET("/validate", middleware.Auth, Handler.Validate)
	//v1.POST("/changePass", Handler.ChangePass)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
