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

// Login 详细说明 ：
// 跨域基础条件 1
// 当AJAX请求的协议、域名或端口与目标服务器不一致时，就会触发跨域限制1。
//
// secure属性的作用 1
// Cookie.secure=true表示该Cookie仅通过HTTPS协议传输，这是跨域请求中携带Cookie的必要条件之一1。
// 需与SameSite=None属性配合使用，才能在跨域请求中发送Cookie1。
//
// 其他关键配置 1
// 前端设置 ：AJAX请求需设置withCredentials=true（XMLHttpRequest）或credentials: 'include'（Fetch API）1。
// 后端配置 ：服务器需在CORS响应中设置Access-Control-Allow-Credentials=true，并明确指定Access-Control-Allow-Origin为具体源（非*）1。
//
// 注意事项 2
// SameSite=None和Secure属性需同时设置，否则Cookie不会在跨域请求中发送1。
// JSONP已被淘汰，不支持跨域请求携带Cookie1。
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

	// 设置cookie, 支持ajax
	c.SetSameSite(http.SameSiteNoneMode)
	// todo Host 白名单
	c.SetCookie(middle_auth.AdminJwtKey, tokenString, 86400*30, "/", c.Request.Host, true, true)
	c.JSON(http.StatusOK, appresp.Reps(gin.H{"token": tokenString, "info": adminInfo, "jump_url": loginParams.JumpURL}, nil))
}

func (h *authHandler) Logout(c *gin.Context) {
	// 设置cookie
	c.SetCookie(middle_auth.AdminJwtKey, "", -1, "/", c.Request.Host, false, true)
	c.JSON(http.StatusOK, appresp.Reps("success", nil))
}
