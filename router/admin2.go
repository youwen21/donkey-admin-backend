package router

import (
	"donkey-admin/app/handler/admin"
	"donkey-admin/middleware"

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
