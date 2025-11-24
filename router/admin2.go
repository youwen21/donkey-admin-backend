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

	cacheAPI := engine.Group("/admin-api/v1").Use(middleware.Cors.GinCors())
	//cacheAPI.Use(middleware.AdminToken())
	//cacheAPI.Use(middleware.BrowserCacheMiddleware)
	{
		// 菜单树
		cacheAPI.GET("/menu/tree", admin.RbacHandler.GetMenuTree)
	}

	// 权限相关
	//{
	//	adminOpenAPI.GET("/role_menu_perm", admin.RoleMenuPermHandler.Detail)        // 获取指定系统 指定角色拥有的 菜单和按钮权限
	//	adminOpenAPI.GET("/role_menu_perm/:role_id", admin.RoleMenuPermHandler.Get) // 获取指定系统 指定角色拥有的 菜单和按钮权限
	//	adminOpenAPI.POST("/role_menu_perm", admin.RoleMenuPermHandler.Edit)        // 更新指定系统 指定角色拥有的权限，t_role_menu
	//
	//	adminOpenAPI.GET("/menu_operation_perm", admin.MenuOperaPermHandler.Get) // 获取指定系统 的 菜单 包含按钮
	//
	//	adminOpenAPI.GET("/user_opera_perm", admin.UserPermHandler.GetUserOperationIdList) // 获取 指定用户的按钮列表
	//
	//	adminOpenAPI.GET("/get_sub_user_id_list", admin.OrgPermHandler.GetSubUserIdList) // 获取 下级用户Id列表
	//	adminOpenAPI.GET("/get_sub_user_list", admin.OrgPermHandler.GetSubUserList)      // 获取 下级用户Id列表
	//}
}
