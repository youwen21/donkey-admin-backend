package middleware

import (
	"github.com/gin-gonic/gin"
	"gofly/app/service/iuser_permission/user_permission_admin"
	"gofly/app/service/iuser_permission/user_permission_def"
	"net/http"
)

func CheckMenuPerm(menuId int) func(c *gin.Context) {
	return func(c *gin.Context) {
		adminUId := GetAdminId(c)
		if adminUId == 0 {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "用户未登录", "data": ""})
			c.Abort()
		}

		form := new(user_permission_def.MyForm)
		form.OperatorUid = adminUId
		form.SystemId = 1

		ok, err := user_permission_admin.AdminSrv.CheckMenuPermission(adminUId, menuId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1002, "msg": err.Error(), "data": ""})
			c.Abort()
		}
		if !ok {
			c.JSON(http.StatusOK, gin.H{"code": 1001, "msg": "用户无权限", "data": ""})
			c.Abort()
		}
	}
}

// CheckOperaPerm 用户权限检查 中间件，必须在登录中间件之后调用
func CheckOperaPerm(operationId int) func(c *gin.Context) {
	return func(c *gin.Context) {
		adminUId := GetAdminId(c)
		if adminUId == 0 {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": "用户未登录", "data": ""})
			c.Abort()
		}
		form := new(user_permission_def.MyForm)
		form.OperatorUid = adminUId
		form.SystemId = 1

		ok, err := user_permission_admin.AdminSrv.CheckOperationPermission(adminUId, operationId)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1002, "msg": err.Error(), "data": ""})
			c.Abort()
		}
		if !ok {
			c.JSON(http.StatusOK, gin.H{"code": 1001, "msg": "用户无权限", "data": ""})
			c.Abort()
		}
	}
}
