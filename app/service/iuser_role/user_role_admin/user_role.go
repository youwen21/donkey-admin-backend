package user_role_admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/iuser_role"
	"donkey-admin/app/service/iuser_role/user_role_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *user_role_def.UserRoleQueryForm) (*user_role_def.UserRoleQueryResEx, error) {
	res, err := iuser_role.Srv.Query(f)
	// biz process
	result := new(user_role_def.UserRoleQueryResEx)
	result.Total = res.Total
	result.List = make([]user_role_def.UserRoleExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.UserRole) user_role_def.UserRoleExDTO {
	return user_role_def.UserRoleExDTO{
		// TODO
		UserRole: v,
	}
}

func (aSrv *adminSrv) GetList(f *user_role_def.UserRoleQueryForm) ([]model.UserRole, error) {
	res, err := iuser_role.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.UserRole, error) {
	res, err := iuser_role.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.UserRole, error) {
	res, err := iuser_role.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.UserRole, error) {
	res, err := iuser_role.Srv.GetMulti(pkList)
	return res, err
}
