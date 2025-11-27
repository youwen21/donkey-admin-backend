package admin

import (
	"gofly/app/service/iuser_permission/user_permission_admin"
	"gofly/app/service/iuser_permission/user_permission_def"
	"gofly/middleware/middle_auth"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type userPermissionHdl struct{}

var (
	UserPermissionHdl = &userPermissionHdl{}
)

func (hdl *userPermissionHdl) My(c *gin.Context) {
	form := new(user_permission_def.MyForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	form.OperatorUid = middle_auth.GetAdminId(c)
	if form.SystemId == 0 {
		form.SystemId = 1
	}

	ret, err := user_permission_admin.AdminSrv.My(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userPermissionHdl) Config(c *gin.Context) {
	form := new(user_permission_def.ConfigForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_permission_admin.AdminSrv.Config(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userPermissionHdl) Save(c *gin.Context) {
	form := new(user_permission_def.SetPermissionForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	form.OperatorUid = middle_auth.GetAdminId(c)
	form.Unique()

	affected, err := user_permission_admin.AdminSrv.Save(form)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
