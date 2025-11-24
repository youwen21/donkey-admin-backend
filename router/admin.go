package router

import (
	"gofly/app/handler/admin"
	"gofly/middleware"

	"github.com/gin-gonic/gin"
)

/*  */

func initAdmin(engine *gin.Engine) {
	// .Use(middleware.AdminToken())
	AdminGroup := engine.Group("/admin-api/v1").Use(middleware.Cors.GinCors())
	AdminGroup.GET("/organization/query", admin.OrganizationHdl.Query)      // query list
	AdminGroup.GET("/organization/get", admin.OrganizationHdl.Get)          // get one detail
	AdminGroup.POST("/organization/add", admin.OrganizationHdl.Add)         // insert
	AdminGroup.POST("/organization/update", admin.OrganizationHdl.Update)   // update
	AdminGroup.POST("/organization/del", admin.OrganizationHdl.Delete)      // delete
	AdminGroup.POST("/organization/setInfo", admin.OrganizationHdl.SetInfo) // setInfo

	AdminGroup.GET("/role/query", admin.RoleHdl.Query)      // query list
	AdminGroup.GET("/role/get", admin.RoleHdl.Get)          // get one detail
	AdminGroup.POST("/role/add", admin.RoleHdl.Add)         // insert
	AdminGroup.POST("/role/update", admin.RoleHdl.Update)   // update
	AdminGroup.POST("/role/del", admin.RoleHdl.Delete)      // delete
	AdminGroup.POST("/role/setInfo", admin.RoleHdl.SetInfo) // setInfo

	AdminGroup.GET("/role_menu/query", admin.RoleMenuHdl.Query)    // query list
	AdminGroup.GET("/role_menu/get", admin.RoleMenuHdl.Get)        // get one detail
	AdminGroup.POST("/role_menu/add", admin.RoleMenuHdl.Add)       // insert
	AdminGroup.POST("/role_menu/update", admin.RoleMenuHdl.Update) // update
	AdminGroup.POST("/role_menu/del", admin.RoleMenuHdl.Delete)    // delete

	AdminGroup.GET("/user/query", admin.UserHdl.Query)      // query list
	AdminGroup.GET("/user/get", admin.UserHdl.Get)          // get one detail
	AdminGroup.POST("/user/add", admin.UserHdl.Add)         // insert
	AdminGroup.POST("/user/update", admin.UserHdl.Update)   // update
	AdminGroup.POST("/user/del", admin.UserHdl.Delete)      // delete
	AdminGroup.POST("/user/setInfo", admin.UserHdl.SetInfo) // setInfo

	AdminGroup.GET("/user_role/query", admin.UserRoleHdl.Query)    // query list
	AdminGroup.GET("/user_role/get", admin.UserRoleHdl.Get)        // get one detail
	AdminGroup.POST("/user_role/add", admin.UserRoleHdl.Add)       // insert
	AdminGroup.POST("/user_role/update", admin.UserRoleHdl.Update) // update
	AdminGroup.POST("/user_role/del", admin.UserRoleHdl.Delete)    // delete

	//AdminGroup.GET("/user_permissions/query", admin.UserPermissionsHdl.Detail)    // query list
	//AdminGroup.GET("/user_permissions/get", admin.UserPermissionsHdl.Get)        // get one detail
	//AdminGroup.POST("/user_permissions/add", admin.UserPermissionsHdl.Add)       // insert
	//AdminGroup.POST("/user_permissions/update", admin.UserPermissionsHdl.Update) // update
	//AdminGroup.POST("/user_permissions/del", admin.UserPermissionsHdl.Delete)    // delete

	AdminGroup.GET("/menu/query", admin.MenuHdl.Query)      // query list
	AdminGroup.GET("/menu/get", admin.MenuHdl.Get)          // get one detail
	AdminGroup.POST("/menu/add", admin.MenuHdl.Add)         // insert
	AdminGroup.POST("/menu/update", admin.MenuHdl.Update)   // update
	AdminGroup.POST("/menu/del", admin.MenuHdl.Delete)      // delete
	AdminGroup.POST("/menu/setInfo", admin.MenuHdl.SetInfo) // setInfo

	AdminGroup.GET("/operation/query", admin.OperationHdl.Query)      // query list
	AdminGroup.GET("/operation/get", admin.OperationHdl.Get)          // get one detail
	AdminGroup.POST("/operation/add", admin.OperationHdl.Add)         // insert
	AdminGroup.POST("/operation/update", admin.OperationHdl.Update)   // update
	AdminGroup.POST("/operation/del", admin.OperationHdl.Delete)      // delete
	AdminGroup.POST("/operation/setInfo", admin.OperationHdl.SetInfo) // setInfo

	AdminGroup.GET("/subsystem/query", admin.SubsystemHdl.Query)      // query list
	AdminGroup.GET("/subsystem/get", admin.SubsystemHdl.Get)          // get one detail
	AdminGroup.POST("/subsystem/add", admin.SubsystemHdl.Add)         // insert
	AdminGroup.POST("/subsystem/update", admin.SubsystemHdl.Update)   // update
	AdminGroup.POST("/subsystem/del", admin.SubsystemHdl.Delete)      // delete
	AdminGroup.POST("/subsystem/setInfo", admin.SubsystemHdl.SetInfo) // setInfo

	AdminGroup.GET("/user_permission/detail", admin.UserPermissionsHdl.Detail)                // get detail
	AdminGroup.POST("/user_permission/setPermission", admin.UserPermissionsHdl.SetPermission) // save

}
