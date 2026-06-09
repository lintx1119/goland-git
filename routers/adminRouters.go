package routers

import (
	"github.com/gin-gonic/gin"
	"go-study/controllers/admin"
	"go-study/middlewares"
)

func AdminRoutersInit(r *gin.Engine) {
	userController := &admin.UserController{}
	adminRouters := r.Group("/admin") // r.Group("/admin", middlewares.Init)
	adminRouters.Use(middlewares.Init)
	{
		adminRouters.GET("/get", userController.GetUser)
		adminRouters.GET("/add", userController.AddUser)
	}
}
