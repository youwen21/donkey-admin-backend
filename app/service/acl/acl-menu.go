package acl

import (
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/iuser"
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
	adminInfo, err := iuser.Srv.Get(f.AdminUserId)
	if err != nil {
		return nil, err
	}

	// root , 根据system获取菜单
	if adminInfo.IsRoot == 1 {
		return imenu.Srv.GetListBySystemId(f.SystemId)
	}

	// 非root, 获取用户授权的菜单
	return nil, nil
	//roleIds, _ := s.GetRoleIds(f.AdminUserId)
	//ruleMenus, err := irole_menu.Srv.GetListBy(f.SystemId, roleIds)
	//
	//menuIdList := make([]int, 0)
	//for _, v := range ruleMenus {
	//	if !slices.Contains(menuIdList, v.MenuId) {
	//		menuIdList = append(menuIdList, v.MenuId)
	//	}
	//}
	//
	//return imenu.Srv.GetLisByPkList(menuIdList)
}
