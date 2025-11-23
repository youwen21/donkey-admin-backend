package admin

import (
	"github.com/gin-gonic/gin"
	"gofly/app/service/auth"
	"gofly/app/service/auth/auth_def"
	"gofly/middleware/middle_auth"
	"gofly/req-resp/appresp"
	"net/http"
)

type authHandler struct {
}

var (
	AuthHandler = new(authHandler)
)

type login struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	KeepLogin string `json:"keep_login" form:"keep_login"`
	JumpURL   string `json:"jump_url" form:"jump_url"`
}

func (h *authHandler) Login(c *gin.Context) {
	loginParams := &auth_def.AuthForm{}
	err := c.ShouldBind(&loginParams)
	if nil != err {
		c.JSON(http.StatusBadRequest, appresp.Err(err))
	}

	tokenString, adminInfo, err := auth.Srv.Login(loginParams)
	if err != nil {
		c.JSON(http.StatusOK, appresp.Err(err))
		return
	}

	// 设置cookie
	c.SetCookie(middle_auth.AdminJwtKey, tokenString, 86400*30, "/", c.Request.Host, false, true)
	c.JSON(http.StatusOK, appresp.Reps(gin.H{"token": tokenString, "info": adminInfo, "jump_url": loginParams.JumpURL}, nil))
}

func (h *authHandler) Logout(c *gin.Context) {
	// 设置cookie
	c.SetCookie(middle_auth.AdminJwtKey, "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusOK, appresp.Reps("success", nil))
}
