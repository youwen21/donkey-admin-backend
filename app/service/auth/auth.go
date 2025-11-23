package auth

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"gofly/app/service/auth/auth_def"
	"gofly/app/service/iuser"
	"gofly/lib/libutils"
	"gofly/middleware/middle_auth"
	"time"
)

type srv struct{}

var (
	Srv = &srv{}
)

func (s *srv) Login(f *auth_def.AuthForm) (string, *auth_def.AdminInfo, error) {
	adminData, err := iuser.Srv.GetByUsername(f.Username)

	if nil == adminData || err != nil {
		return "", nil, errors.New("用户名或密码错误")
	}

	// 密码校验
	genPwd := libutils.EncryptWord(f.Password)
	if adminData.Password != genPwd {
		return "", nil, errors.New("用户名或密码错误2")
	}

	tokenString, err := libutils.Jwt.GenToken(middle_auth.AdminJwtSecret, jwt.MapClaims{middle_auth.AdminUserKey: adminData.Id, "exp": time.Now().Unix() + 86400*30})
	if err != nil {
		return "", nil, err
	}

	adminInfo := &auth_def.AdminInfo{
		Id:     adminData.Id,
		Name:   adminData.Name,
		Avatar: adminData.Avatar,
		IsRoot: adminData.IsRoot,
		RoleId: adminData.RoleId,
		OrgId:  adminData.OrgId,
	}
	return tokenString, adminInfo, nil
}
