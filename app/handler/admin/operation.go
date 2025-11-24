package admin

import (
	"gofly/app/model"
	"gofly/app/service/ioperation"
	"gofly/app/service/ioperation/operation_admin"
	"gofly/app/service/ioperation/operation_def"
	"gofly/middleware/middle_auth"
	"gofly/req-resp/appresp"
	"net/http"

	"github.com/gin-gonic/gin"
)

/*  */

type operationHdl struct{}

var (
	OperationHdl = &operationHdl{}
)

func (hdl *operationHdl) Search(c *gin.Context) {
	hdl.Query(c)
}

func (hdl *operationHdl) Query(c *gin.Context) {
	form := new(operation_def.OperationQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	// if form.Ids != "" {
	//     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	// }

	form.OrderBy = append(form.OrderBy, "id desc")

	ret, err := operation_admin.AdminSrv.Query(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *operationHdl) GetList(c *gin.Context) {
	form := new(operation_def.OperationQueryForm)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := operation_admin.AdminSrv.GetList(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *operationHdl) Get(c *gin.Context) {
	info := new(model.Operation)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := operation_admin.AdminSrv.Get(info.Id)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *operationHdl) Add(c *gin.Context) {
	info := new(model.Operation)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.CreateUid = middle_auth.GetAdminId(c)

	err := ioperation.Srv.Insert(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *operationHdl) Update(c *gin.Context) {
	info := new(model.Operation)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}
	info.UpdateUid = middle_auth.GetAdminId(c)

	_, err := ioperation.Srv.Update(info)
	c.JSON(http.StatusOK, appresp.Reps(info, err))
}

func (hdl *operationHdl) Delete(c *gin.Context) {
	info := new(model.Operation)
	if err := c.ShouldBind(info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	err := ioperation.Srv.Delete(info.Id)

	c.JSON(http.StatusOK, appresp.Reps("", err))
}

// SetInfo 没有 update 安全，优先使用update
// 弥补 int=0, string="" 时，update 不更新的问题
func (hdl *operationHdl) SetInfo(c *gin.Context) {
	info := make(map[string]any)
	if err := c.ShouldBind(&info); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	affected, err := ioperation.Srv.SetInfo(info)
	c.JSON(http.StatusOK, appresp.Reps(affected, err))
}
