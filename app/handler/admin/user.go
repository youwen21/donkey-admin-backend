package admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iuser"
	"donkey-admin/app/service/iuser/user_admin"
	"donkey-admin/app/service/iuser/user_def"
	"donkey-admin/middleware"
	"donkey-admin/req-resp/appresp"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

/*  */

type userHdl struct{}

var (
	UserHdl = &userHdl{}
)

func (hdl *userHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *userHdl) Query(c *gin.Context) {
	form := new(user_def.UserQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	if form.SearchKey != "" {
		form.SearchKey = strings.TrimSpace(form.SearchKey)
	}

	ret, err := user_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userHdl) GetList(c *gin.Context) {
	form := new(user_def.UserQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userHdl) Get(c *gin.Context) {
	info := new(model.User)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := user_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *userHdl) Add(c *gin.Context) {
	info := new(model.User)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.CreateUid = middleware.GetAdminId(c)

	_, err := user_admin.AdminSrv.Add(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userHdl) Update(c *gin.Context) {
	info := new(model.User)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.UpdateUid = middleware.GetAdminId(c)

	_, err := user_admin.AdminSrv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *userHdl) Delete(c *gin.Context) {
	info := new(model.User)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := iuser.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// SetInfo 没有 update 安全，优先使用update
// 弥补 int=0, string="" 时，update 不更新的问题
func (hdl *userHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := iuser.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
