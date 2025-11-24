package admin

import (
	"gofly/app/model"
	"gofly/app/service/iuser_role"
	"gofly/app/service/iuser_role/user_role_admin"
	"gofly/app/service/iuser_role/user_role_def"
	"gofly/middleware/middle_auth"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type userRoleHdl struct{}

var (
	UserRoleHdl = &userRoleHdl{}
)

func (hdl *userRoleHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *userRoleHdl) Query(c *gin.Context) {
	form := new(user_role_def.UserRoleQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := user_role_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userRoleHdl) GetList(c *gin.Context) {
	form := new(user_role_def.UserRoleQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_role_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userRoleHdl) Get(c *gin.Context) {
	info := new(model.UserRole)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_role_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userRoleHdl) Add(c *gin.Context) {
	info := new(model.UserRole)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.CreateUid = middle_auth.GetAdminId(c)

	err := iuser_role.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userRoleHdl) Update(c *gin.Context) {
	info := new(model.UserRole)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	_, err := iuser_role.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userRoleHdl) Delete(c *gin.Context) {
	info := new(model.UserRole)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := iuser_role.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}
