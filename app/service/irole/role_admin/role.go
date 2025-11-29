package role_admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/irole"
	"donkey-admin/app/service/irole/role_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *role_def.RoleQueryForm) (*role_def.RoleQueryResEx, error) {
	res, err := irole.Srv.Query(f)
	// biz process
	result := new(role_def.RoleQueryResEx)
	result.Total = res.Total
	result.List = make([]role_def.RoleExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.Role) role_def.RoleExDTO {
	return role_def.RoleExDTO{
		// TODO
		Role: v,
	}
}

func (aSrv *adminSrv) GetList(f *role_def.RoleQueryForm) ([]model.Role, error) {
	res, err := irole.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.Role, error) {
	res, err := irole.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.Role, error) {
	res, err := irole.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.Role, error) {
	res, err := irole.Srv.GetMulti(pkList)
	return res, err
}
