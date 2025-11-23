package admin

import (
	"gofly/app/model"
	"gofly/app/service/irole_menu"
	"gofly/app/service/irole_menu/role_menu_admin"
	"gofly/app/service/irole_menu/role_menu_def"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type roleMenuHdl struct{}

var (
	RoleMenuHdl = &roleMenuHdl{}
)

func (hdl *roleMenuHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *roleMenuHdl) Query(c *gin.Context) {
	form := new(role_menu_def.RoleMenuQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := role_menu_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *roleMenuHdl) GetList(c *gin.Context) {
	form := new(role_menu_def.RoleMenuQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := role_menu_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *roleMenuHdl) Get(c *gin.Context) {
	info := new(model.RoleMenu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := role_menu_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *roleMenuHdl) Add(c *gin.Context) {
	info := new(model.RoleMenu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	err := irole_menu.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *roleMenuHdl) Update(c *gin.Context) {
	info := new(model.RoleMenu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	_, err := irole_menu.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *roleMenuHdl) Delete(c *gin.Context) {
	info := new(model.RoleMenu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := irole_menu.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}
