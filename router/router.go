package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"time"
	v1 "vue-typescript-go-gin/api/v1"
	"vue-typescript-go-gin/controllers"
)

func Router(dbConn *gorm.DB) {
	todoHandler := controllers.TodoHandler{
		Db: dbConn,
	}
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "DELETE", "POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "*"
		},
		MaxAge: 12 * time.Hour,
	}))
	r.LoadHTMLGlob("templates/*")
	r.GET("/todo", todoHandler.GetAll)
	r.POST("/todo", todoHandler.CreateTask)
	r.GET("/todo/:id", todoHandler.EditTask)
	r.POST("//todo/edit/:id", todoHandler.UpdateTask)
	r.POST("/todo/delete/:id", todoHandler.DeleteTask)

	apiV1 := r.Group("/api/v1")
	{
		apiTodoHandler := v1.TodoHandler{
			Db: dbConn,
		}
		apiV1.GET("/todo", apiTodoHandler.GetAll)
		apiV1.POST("/todo", apiTodoHandler.CreateTask)
		apiV1.GET("/todo/:id", apiTodoHandler.EditTask)
		apiV1.PUT("/todo/:id", apiTodoHandler.UpdateTask)
		apiV1.DELETE("/todo/:id", apiTodoHandler.DeleteTask)
	}
	r.Run(":8080")
}
