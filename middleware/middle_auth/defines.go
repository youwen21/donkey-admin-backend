package middle_auth

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
