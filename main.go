package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/jinzhu/gorm/dialects/postgres"
	"task-manager/controllers"
	"task-manager/models"
	"task-manager/middlewares"
)

func main() {
	// Database connection
	db, err := gorm.Open("postgres", "host=db port=5432 user=user dbname=task_db sslmode=disable password=password")
	if err != nil {
		panic("Failed to connect to database!")
	}
	defer db.Close()

	db.AutoMigrate(&models.User{}, &models.Task{})

	r := gin.Default()

	// Public routes
	r.POST("/api/register", func(c *gin.Context) { controllers.Register(c, db) })
	r.POST("/api/login", func(c *gin.Context) { controllers.Login(c, db) })

	// Protected routes
	protected := r.Group("/api")
	protected.Use(middlewares.AuthRequired())
	{
		protected.POST("/tasks", func(c *gin.Context) { controllers.CreateTask(c, db) })
		protected.GET("/tasks", func(c *gin.Context) { controllers.GetTasks(c, db) })
		protected.PUT("/tasks/:id", func(c *gin.Context) { controllers.UpdateTask(c, db) })
		protected.DELETE("/tasks/:id", func(c *gin.Context) { controllers.DeleteTask(c, db) })
		protected.GET("/tasks/search", func(c *gin.Context) { controllers.SearchTasks(c, db) })
	}

	r.Run(":8080")
}
