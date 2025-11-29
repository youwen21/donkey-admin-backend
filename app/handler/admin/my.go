package admin

import (
	"donkey-admin/app/service/iuser"
	"donkey-admin/middleware"
	"donkey-admin/req-resp/appresp"
	"github.com/gin-gonic/gin"
	"net/http"
)

type myHandler struct {
}

var (
	MyHandler = new(myHandler)
)

func (h *myHandler) StaffInfo(c *gin.Context) {
	adminId := middleware.GetAdminId(c)
	info, err := iuser.Srv.Get(adminId)
	if err != nil {
		c.JSON(http.StatusOK, appresp.Err(err))
		return
	}
	c.JSON(http.StatusOK, appresp.Reps(info.StaffInfo, nil))
}
