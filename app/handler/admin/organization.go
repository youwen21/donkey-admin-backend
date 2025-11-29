package admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iorganization"
	"donkey-admin/app/service/iorganization/organization_admin"
	"donkey-admin/app/service/iorganization/organization_def"
	"donkey-admin/middleware"
	"donkey-admin/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type organizationHdl struct{}

var (
	OrganizationHdl = &organizationHdl{}
)

func (hdl *organizationHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *organizationHdl) Query(c *gin.Context) {
	form := new(organization_def.OrganizationQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := organization_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *organizationHdl) GetList(c *gin.Context) {
	form := new(organization_def.OrganizationQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := organization_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *organizationHdl) Get(c *gin.Context) {
	info := new(model.Organization)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := organization_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *organizationHdl) Add(c *gin.Context) {
	info := new(model.Organization)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.CreateUid = middleware.GetAdminId(c)

	_, err := organization_admin.AdminSrv.Add(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *organizationHdl) Update(c *gin.Context) {
	info := new(model.Organization)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.UpdateUid = middleware.GetAdminId(c)

	_, err := organization_admin.AdminSrv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *organizationHdl) Delete(c *gin.Context) {
	info := new(model.Organization)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := iorganization.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// 优先使用update
// SetInfo 弥补 int=0, string="" update 不更新问题
func (hdl *organizationHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := iorganization.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
