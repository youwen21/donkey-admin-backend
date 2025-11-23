package admin

import (
	"gofly/app/model"
	"gofly/app/service/iuser_permissions"
	"gofly/app/service/iuser_permissions/user_permissions_admin"
	"gofly/app/service/iuser_permissions/user_permissions_def"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type userPermissionsHdl struct{}

var (
	UserPermissionsHdl = &userPermissionsHdl{}
)

func (hdl *userPermissionsHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *userPermissionsHdl) Query(c *gin.Context) {
	form := new(user_permissions_def.UserPermissionsQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := user_permissions_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userPermissionsHdl) GetList(c *gin.Context) {
	form := new(user_permissions_def.UserPermissionsQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_permissions_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userPermissionsHdl) Get(c *gin.Context) {
	info := new(model.UserPermissions)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_permissions_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userPermissionsHdl) Add(c *gin.Context) {
	info := new(model.UserPermissions)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	err := iuser_permissions.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userPermissionsHdl) Update(c *gin.Context) {
	info := new(model.UserPermissions)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	_, err := iuser_permissions.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userPermissionsHdl) Delete(c *gin.Context) {
	info := new(model.UserPermissions)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := iuser_permissions.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// 优先使用update
// SetInfo 弥补 int=0, string="" update 不更新问题
func (hdl *userPermissionsHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := iuser_permissions.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
