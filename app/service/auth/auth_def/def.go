package auth_def

type LoginForm struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
	Remember bool   `json:"remember" form:"remember"`
}
