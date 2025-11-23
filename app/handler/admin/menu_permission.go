package admin

import (
	"github.com/gin-gonic/gin"
	"gofly/app/service/menu_permission"
	"gofly/req-resp/appresp"
	"net/http"
)

/*  */

type menuPermissionHdl struct{}

var (
	MenuPermissionHdl = &menuPermissionHdl{}
)

//func (hdl *menuPermissionHdl) Search(c *gin.Context) {
//	hdl.Detail(c)
//}

//func (hdl *menuPermissionHdl) Query(c *gin.Context) {
//	form := new(menu_permission.Form)
//	if err := c.ShouldBind(&form); err != nil {
//		c.JSON(http.StatusBadRequest, appresp.Err(err))
//		return
//	}
//
//	ret, err := menu_permission.SrvInstance.Detail(form)
//	c.JSON(http.StatusOK, appresp.Reps(ret, err))
//}

func (hdl *menuPermissionHdl) Detail(c *gin.Context) {
	form := new(menu_permission.Form)
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
		return
	}

	ret, err := menu_permission.SrvInstance.Detail(form)
	c.JSON(http.StatusOK, appresp.Reps(ret, err))
}

func (hdl *menuPermissionHdl) Save(c *gin.Context) {
	//form := new(user_permissions_def.UserPermissionsQueryForm)
	//if err := c.ShouldBind(&form); err != nil {
	//	c.JSON(http.StatusBadRequest, appresp.Err(err))
	//	return
	//}
	//// if form.Ids != "" {
	////     form.IdList = libutils.SplitToIntList(form.Ids, ",")
	//// }
	//
	//form.OrderBy = append(form.OrderBy, "id desc")
	//
	//ret, err := user_permissions_admin.AdminSrv.Detail(form)
	//c.JSON(http.StatusOK, appresp.Reps(ret, err))
}
