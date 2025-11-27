package acl

import (
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/iuser"
	"gofly/app/service/iuser_permission"
	"gofly/app/service/iuser_permission/user_permission_def"
	"gofly/app/service/tree/menutree"
)

type menuSrv struct{}

var (
	MenuSrv = &menuSrv{}
)

func (s *menuSrv) GetTreeMenu(f *menutree.UserMenuForm) ([]*menutree.TreeMenu, error) {
	menuList, err := s.GetMenuList(f)
	if err != nil {
		return nil, err
	}

	tree := menutree.ToTree(menuList)
	return tree, nil
}

func (s *menuSrv) GetMenuList(f *menutree.UserMenuForm) ([]model.Menu, error) {
	adminInfo, err := iuser.Srv.Get(f.AdminId)
	if err != nil {
		return nil, err
	}

	// root , 根据system获取菜单
	if adminInfo.IsRoot == 1 {
		return imenu.Srv.GetListBySystemId(f.SystemId)
	}

	// 非root, 获取用户授权的菜单
	permQForm := new(user_permission_def.UserPermissionQueryForm)
	permQForm.SystemId = f.SystemId
	permQForm.UserId = f.AdminId
	permQForm.Page = 1
	permQForm.PageSize = 1000
	permissions, err := iuser_permission.Srv.GetList(permQForm)
	if err != nil {
		return nil, err
	}
	if len(permissions) < 1 {
		return nil, nil
	}
	menuIdList := make([]int, 0)
	for _, v := range permissions {
		menuIdList = append(menuIdList, v.MenuId)
	}
	return imenu.Srv.GetLisByPkList(menuIdList)
}
