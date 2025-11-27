package router

import (
	"gofly/app/handler/admin"
	"gofly/middleware"

	"github.com/gin-gonic/gin"
)

/*  */

func initAdmin2(engine *gin.Engine) {
	adminOpenAPI := engine.Group("/admin-api/v1").Use(middleware.Cors.GinCors())
	{
		adminOpenAPI.POST("/login", admin.AuthHandler.Login)
		adminOpenAPI.POST("/logout", admin.AuthHandler.Logout)
	}

}
