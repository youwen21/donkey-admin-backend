package rbac

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/imenu"
	"donkey-admin/app/service/irole_menu"
	"donkey-admin/app/service/irole_menu/role_menu_def"
	"donkey-admin/app/service/iuser"
	"donkey-admin/app/service/iuser_role"
	"donkey-admin/app/service/tree/menutree"
	"donkey-admin/lib/libutils"
	"slices"
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

	// 非root, 获取用户角色
	// 获取角色菜单
	roleIds, _ := s.GetRoleIds(f.AdminId)
	ruleMenus, err := irole_menu.Srv.GetListBy(f.SystemId, roleIds)

	menuIdList := make([]int, 0)
	for _, v := range ruleMenus {
		if !slices.Contains(menuIdList, v.MenuId) {
			menuIdList = append(menuIdList, v.MenuId)
		}
	}

	return imenu.Srv.GetLisByPkList(menuIdList)
}

func (s *menuSrv) GetRoleIds(userId int) ([]int, error) {
	l, err := iuser_role.Srv.GetListByUserId(userId)
	if err != nil {
		return nil, err
	}
	roleIdList := make([]int, 0)
	for _, v := range l {
		if !slices.Contains(roleIdList, v.RoleId) {
			roleIdList = append(roleIdList, v.RoleId)
		}
	}
	return roleIdList, nil
}

func (s *menuSrv) GetAuthedMenuOperation(systemId int, roleIdList []int) (*role_menu_def.RolesAuthed, error) {
	res, err := irole_menu.Srv.GetListBy(systemId, roleIdList)
	if err != nil {
		return nil, err
	}

	AuthMap := make(map[int]role_menu_def.AuthMenu)
	for _, v := range res {
		operations := libutils.SplitIntList(v.OperaIds, ".")
		authed, ok := AuthMap[v.MenuId]
		if !ok {
			AuthMap[v.MenuId] = role_menu_def.AuthMenu{
				MenuId:          systemId,
				OperationIdList: operations,
			}
		} else {
			for _, op := range operations {
				if !slices.Contains(authed.OperationIdList, op) {
					authed.OperationIdList = append(authed.OperationIdList, op)
				}
			}
			AuthMap[v.MenuId] = authed
		}

	}

	authMenuList := make([]role_menu_def.AuthMenu, 0)
	for _, v := range AuthMap {
		authMenuList = append(authMenuList, v)
	}

	rsAuthed := new(role_menu_def.RolesAuthed)
	rsAuthed.SystemId = systemId
	rsAuthed.RoleIds = roleIdList
	rsAuthed.Authed = authMenuList
	return rsAuthed, nil
}
