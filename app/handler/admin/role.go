package admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/irole"
	"donkey-admin/app/service/irole/role_admin"
	"donkey-admin/app/service/irole/role_def"
	"donkey-admin/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type roleHdl struct{}

var (
	RoleHdl = &roleHdl{}
)

func (hdl *roleHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *roleHdl) Query(c *gin.Context) {
	form := new(role_def.RoleQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := role_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *roleHdl) GetList(c *gin.Context) {
	form := new(role_def.RoleQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := role_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *roleHdl) Get(c *gin.Context) {
	info := new(model.Role)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := role_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *roleHdl) Add(c *gin.Context) {
	info := new(model.Role)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	err := irole.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *roleHdl) Update(c *gin.Context) {
	info := new(model.Role)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	_, err := irole.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *roleHdl) Delete(c *gin.Context) {
	info := new(model.Role)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := irole.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// SetInfo 没有 update 安全，优先使用update
// 弥补 int=0, string="" 时，update 不更新的问题
func (hdl *roleHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := irole.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
