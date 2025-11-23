package admin

import (
	"gofly/app/model"
	"gofly/app/service/isubsystem"
	"gofly/app/service/isubsystem/subsystem_admin"
	"gofly/app/service/isubsystem/subsystem_def"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type subsystemHdl struct{}

var (
	SubsystemHdl = &subsystemHdl{}
)

func (hdl *subsystemHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *subsystemHdl) Query(c *gin.Context) {
	form := new(subsystem_def.SubsystemQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := subsystem_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *subsystemHdl) GetList(c *gin.Context) {
	form := new(subsystem_def.SubsystemQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := subsystem_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *subsystemHdl) Get(c *gin.Context) {
	info := new(model.Subsystem)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := subsystem_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *subsystemHdl) Add(c *gin.Context) {
	info := new(model.Subsystem)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	err := isubsystem.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *subsystemHdl) Update(c *gin.Context) {
	info := new(model.Subsystem)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	_, err := isubsystem.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *subsystemHdl) Delete(c *gin.Context) {
	info := new(model.Subsystem)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := isubsystem.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// SetInfo 没有 update 安全，优先使用update
// 弥补 int=0, string="" 时，update 不更新的问题
func (hdl *subsystemHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := isubsystem.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
