package role_menu_admin

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/irole_menu"
	"donkey-admin/app/service/irole_menu/role_menu_def"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Query(f *role_menu_def.RoleMenuQueryForm) (*role_menu_def.RoleMenuQueryResEx, error) {
	res, err := irole_menu.Srv.Query(f)
	// biz process
	result := new(role_menu_def.RoleMenuQueryResEx)
	result.Total = res.Total
	result.List = make([]role_menu_def.RoleMenuExDTO, len(res.List))
	for i, v := range res.List {
		result.List[i] = aSrv.extendToDTO(v)
	}

	return result, err
}

func (aSrv *adminSrv) extendToDTO(v model.RoleMenu) role_menu_def.RoleMenuExDTO {
	return role_menu_def.RoleMenuExDTO{
		// TODO
		RoleMenu: v,
	}
}

func (aSrv *adminSrv) GetList(f *role_menu_def.RoleMenuQueryForm) ([]model.RoleMenu, error) {
	res, err := irole_menu.Srv.GetList(f)
	return res, err
}

func (aSrv *adminSrv) GetAll() ([]model.RoleMenu, error) {
	res, err := irole_menu.Srv.GetAll()
	return res, err
}

func (aSrv *adminSrv) Get(pk int) (*model.RoleMenu, error) {
	res, err := irole_menu.Srv.Get(pk)
	return res, err
}

func (aSrv *adminSrv) GetMulti(pkList []int) (map[int]model.RoleMenu, error) {
	res, err := irole_menu.Srv.GetMulti(pkList)
	return res, err
}
