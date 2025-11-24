package admin

import (
	"github.com/gin-gonic/gin"
	"gofly/app/service/rbac"
	"gofly/app/service/tree/menutree"
	"gofly/middleware/middle_auth"
	"gofly/req-resp/appresp"
	"net/http"
)

type rbacHandler struct {
}

var (
	RbacHandler = new(rbacHandler)
)

func (h *rbacHandler) GetMenuTree(c *gin.Context) {
	form := new(menutree.UserMenuForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	if form.SystemId == 0 {
		form.SystemId = 1
	}

	form.AdminUserId = middle_auth.GetAdminId(c)

	// 方便测试
	if form.AdminUserId == 0 {
		form.AdminUserId = 1
	}

	res, err := rbac.MenuSrv.GetTreeMenu(form)
	c.JSON(http.StatusOK, appresp.Reps(res, err))
}
