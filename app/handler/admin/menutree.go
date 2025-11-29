package admin

import (
	"donkey-admin/app/service/acl"
	"donkey-admin/app/service/tree/menutree"
	"donkey-admin/middleware"
	"donkey-admin/req-resp/appresp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type menuTreeHandler struct {
}

var (
	MenuTreeHandler = new(menuTreeHandler)
)

func (h *menuTreeHandler) AclMenuTree(c *gin.Context) {
	form := new(menutree.UserMenuForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	if form.SystemId == 0 {
		form.SystemId = 1
	}

	form.AdminId = middleware.GetAdminId(c)

	// 方便测试
	//if form.AdminId == 0 {
	//	form.AdminId = 1
	//}

	res, err := acl.MenuSrv.GetTreeMenu(form)
	c.JSON(http.StatusOK, appresp.Reps(res, err))
}

//func (h *menuTreeHandler) RbacMenuTree(c *gin.Context) {
//	form := new(menutree.UserMenuForm)
//	if err := c.ShouldBind(&form); err != nil {
//		c.JSON(http.StatusBadRequest, appresp.Err(err))
//		return
//	}
//	if form.SystemId == 0 {
//		form.SystemId = 1
//	}
//
//	form.AdminUserId = middleware.GetAdminId(c)
//
//	// 方便测试
//	if form.AdminUserId == 0 {
//		form.AdminUserId = 1
//	}
//
//	res, err := rbac.MenuSrv.GetTreeMenu(form)
//	c.JSON(http.StatusOK, appresp.Reps(res, err))
//}
