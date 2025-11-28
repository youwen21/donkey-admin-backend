package admin

import (
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/imenu/menu_admin"
	"gofly/app/service/imenu/menu_def"
	"gofly/middleware"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type menuHdl struct{}

var (
	MenuHdl = &menuHdl{}
)

func (hdl *menuHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *menuHdl) Query(c *gin.Context) {
	form := new(menu_def.MenuQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := menu_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *menuHdl) GetList(c *gin.Context) {
	form := new(menu_def.MenuQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := menu_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *menuHdl) Get(c *gin.Context) {
	info := new(model.Menu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := menu_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *menuHdl) Add(c *gin.Context) {
	info := new(model.Menu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.CreateUid = middleware.GetAdminId(c)

	_, err := menu_admin.AdminSrv.Add(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *menuHdl) Update(c *gin.Context) {
	info := new(model.Menu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.UpdateUid = middleware.GetAdminId(c)

	_, err := menu_admin.AdminSrv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *menuHdl) Delete(c *gin.Context) {
	info := new(model.Menu)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := imenu.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// SetInfo 没有 update 安全，优先使用update
// 弥补 int=0, string="" 时，update 不更新的问题
func (hdl *menuHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := imenu.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
