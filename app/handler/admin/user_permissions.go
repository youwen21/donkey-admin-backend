package admin

import (
	"github.com/gin-gonic/gin"
	"gofly/app/service/iuser_permissions/user_permissions_admin"
	"gofly/app/service/iuser_permissions/user_permissions_def"
	"gofly/middleware/middle_auth"
	"gofly/req-resp/appresp"
	"net/http"
)

/*  */

type userPermissionsHdl struct{}

var (
	UserPermissionsHdl = &userPermissionsHdl{}
)

func (hdl *userPermissionsHdl) Detail(c *gin.Context) {
	form := new(user_permissions_def.Form)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_permissions_admin.AdminSrv.Detail(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userPermissionsHdl) SetPermission(c *gin.Context) {
	form := new(user_permissions_def.SetPermissionForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	form.OperatorUid = middle_auth.GetAdminId(c)
	form.Unique()

	affected, err := user_permissions_admin.AdminSrv.SetPermission(form)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
