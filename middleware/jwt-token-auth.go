package middleware

import (
	"gofly/lib/libutils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

const (
	// admin token
	AdminAuthKey   = "X-Admin-Authorization" // header 或者 cookie key
	AdminJwtSecret = "ADMIN_JWT_SECRET"      // jwt secret
	AdminUserKey   = "admin_id"              // auth 认证成功后，uid 存在gin.Context key

	// inner token
	InnerAuthKey   = "X-Inner-Authorization" // header 或者 cookie key
	InnerJwtSecret = "INNER_JWT_SECRET"      // jwt secret
	InnerSystemKey = "system_id"             // auth 认证成功后，uid 存在gin.Context的key

)

func GetAdminId(c *gin.Context) int {
	return c.GetInt(AdminUserKey)
}

func GetSystemId(c *gin.Context) int {
	return c.GetInt(InnerSystemKey)
}

func jwtTokenWare(tokenKey string, secret string, storeKey string) func(c *gin.Context) {
	// storeKey gin 和 jwt claims中的key
	// tokenKey header头 或者 cookie 包含jwt串的key
	// secret jwt解密密钥

	return func(c *gin.Context) {
		claims, err := libutils.Jwt.CheckToken(c, tokenKey, secret)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": ""})
			c.Abort()
		}

		value := claims[storeKey]
		c.Set(storeKey, cast.ToInt(value))
		c.Next()
	}
}

// InnerToken 内容服务接口认证
//func InnerToken(c *gin.Context) {
//	err := CheckToken(c, InnerJwtKey, InnerJwtSecret, InnerSystemId)
//	if err != nil {
//		c.JSON(http.StatusOK, gin.H{"code": 1, "msg": err.Error(), "data": ""})
//		c.Abort()
//	}
//	c.Next()
//}

// InnerToken 内部api token
func InnerToken() func(c *gin.Context) {
	tokenKey := InnerAuthKey // read from config
	secret := InnerJwtSecret //  os.Getenv("USER_JWT_SECRET")
	systemKey := InnerSystemKey
	return jwtTokenWare(tokenKey, secret, systemKey)
}

// AdminToken 避免每次调用AdminTokenWare中间件都要传入 tokenKet, secret, userKey, 三个参数， 封装一下， 自行改动。
func AdminToken() func(c *gin.Context) {
	return jwtTokenWare(AdminAuthKey, AdminJwtSecret, AdminUserKey)
}
