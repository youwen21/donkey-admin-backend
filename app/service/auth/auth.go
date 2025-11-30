package auth

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/auth/auth_def"
	"donkey-admin/app/service/iuser"
	"donkey-admin/lib/libutils"
	"donkey-admin/middleware"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Login(f *auth_def.LoginForm) (string, *model.UserBaseInfo, error) {
	adminData, err := iuser.Srv.GetByUsername(f.Username)

	if nil == adminData || err != nil {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 密码校验
	genPwd := libutils.EncryptWord(f.Password)
	if adminData.Password != genPwd {
		return "", nil, errors.New("用户名或密码错误2")
	}

	tokenString, err := libutils.Jwt.GenToken(middleware.AdminJwtSecret, jwt.MapClaims{middleware.AdminCtxKey: adminData.Id, "exp": time.Now().Unix() + 86400*7}) // 记住登录7天，或者7天不关闭网页，登录最长有效期7天
	if err != nil {
		return "", nil, err
	}

	//adminInfo := &auth_def.AdminInfo{
	//	Id:     adminData.Id,
	//	Name:   adminData.Name,
	//	Avatar: adminData.Avatar,
	//	IsRoot: adminData.IsRoot,
	//	RoleId: adminData.RoleId,
	//	OrgId:  adminData.OrgId,
	//}
	return tokenString, &adminData.UserBaseInfo, nil
}
