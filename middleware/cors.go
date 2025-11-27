package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	methods = `POST, OPTIONS, GET, PUT, DELETE`
	headers = `Content-Type,Content-Length,Accept-Encoding,X-CSRF-Token,Authorization,Accept,Origin,Men,Cache-Control,X-Requested-With,Name,DNT,HOST,Pragma,Referer,Duo,Range,user-Agent,token`
)

type cors struct {
}

var Cors = &cors{}

func (c *cors) RawCors(next http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("X-ORIGIN")
		if origin == "" {
			origin = r.Header.Get("ORIGIN")
		}
		// 允许跨域的方法
		w.Header().Set("Access-Control-Allow-Methods", methods)

		// 允许跨域的 头 前端可以访问的
		w.Header().Set("Access-Control-Expose-Headers", "X-Auth-Token, X-Request-Id")
		// 允许跨域的 头 服务器端允许的
		w.Header().Set("Access-Control-Allow-Headers", headers)

		// 允许跨域 cookie
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// 允许跨域 cookie 的域名, 必须https
		w.Header().Set("Access-Control-Allow-Origin", origin) // 最好是配置指定，X-ORIGIN 不安全。

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (c *cors) GinCors() gin.HandlerFunc {
	// CORSMiddleware 跨域
	// @see https://stackoverflow.com/questions/29418478/go-gin-framework-cors
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("X-ORIGIN")
		if origin == "" {
			origin = c.Request.Header.Get("ORIGIN")
		}
		c.Writer.Header().Set("Access-Control-Allow-Methods", methods)

		c.Writer.Header().Set("Access-Control-Expose-Headers", "Access-Control-Allow-Origin")
		c.Writer.Header().Set("Access-Control-Allow-Headers", headers)

		c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
