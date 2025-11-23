package user_permissions_admin

import (
	"gofly/app/model"
	"gofly/app/service/iuser_permissions"
	"gofly/app/service/iuser_permissions/user_permissions_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *user_permissions_def.UserPermissionsQueryForm) (*user_permissions_def.UserPermissionsQueryResEx, error) {
	res, err := iuser_permissions.Srv.Query(f)
	// biz process
	result := new(user_permissions_def.UserPermissionsQueryResEx)
	result.Total = res.Total
	result.List = make([]user_permissions_def.UserPermissionsExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.UserPermissions) user_permissions_def.UserPermissionsExDTO {
	return user_permissions_def.UserPermissionsExDTO{
		// TODO
		UserPermissions: v,
	}
}

func (aSrv *adminSrv) GetList(f *user_permissions_def.UserPermissionsQueryForm) ([]model.UserPermissions, error) {
	res, err := iuser_permissions.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.UserPermissions, error) {
	res, err := iuser_permissions.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.UserPermissions, error) {
	res, err := iuser_permissions.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.UserPermissions, error) {
	res, err := iuser_permissions.Srv.GetMulti(pkList)
	return res, err
}

func (aSrv *adminSrv) Add(v *model.UserPermissions) (*model.UserPermissions, error) {
	err := iuser_permissions.Srv.Insert(v)
	return v, err
}

func (aSrv *adminSrv) Update(v *model.UserPermissions) (int64, error) {
	affected, err := iuser_permissions.Srv.Update(v)
	return affected, err
}
