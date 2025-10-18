package routes

import (
	"github.com/gin-gonic/gin"
	"ngodinginaja-be/controllers"
	"ngodinginaja-be/middleware"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(200, gin.H{"user": user})
		})


		auth.GET("/course", controllers.GetCourse)
		auth.GET("/courses/:id/modules", controllers.GetModule)


		auth.POST("/course/create", controllers.CreateCourse)
		auth.POST("/course/module/create", controllers.CreateModule)
	}
}
