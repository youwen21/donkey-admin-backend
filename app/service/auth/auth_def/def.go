package auth_def

type AuthForm struct {
	Username  string `json:"username" form:"username"`
	Password  string `json:"password" form:"password"`
	KeepLogin string `json:"keep_login" form:"keep_login"`
	JumpURL   string `json:"jump_url" form:"jump_url"`
}

type AdminInfo struct {
	Id     int    `json:"id" form:"id"`           //
	Name   string `json:"name" form:"name"`       // 登陆名
	IsRoot int8   `json:"is_root" form:"is_root"` // 是否root用户
	Avatar string `json:"avatar" form:"avatar"`   // 用户头像
	RoleId int    `json:"role_id" form:"role_id"` // 角色id
	OrgId  int    `json:"org_id" form:"org_id"`   // 所属组织
}
