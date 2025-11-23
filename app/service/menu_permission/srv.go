package menu_permission

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/ioperation"
	"gofly/app/service/iuser"
	"gofly/app/service/iuser_permissions"
	"gofly/app/service/tree/menutree"
	"gofly/lib/libutils"
)

type srv struct {
}

var SrvInstance = srv{}

func (s *srv) Detail(f *Form) (*MenuPermission, error) {
	result := new(MenuPermission)
	result.Form = f

	menuOperations, err := s.MenuOperations(f)
	if err != nil {
		return nil, err
	}
	result.SystemMenu = menuOperations

	userInfo, _ := iuser.Srv.Get(f.UserId)
	if userInfo == nil {
		return nil, errors.New("user not found")
	}

	// 管理员
	if userInfo.IsRoot == 1 {
		result.IsAdmin = true
		return result, nil
	}

	// 用户权限
	userPermissions, err := s.GetAuthorizedPermissions(f)
	if err != nil {
		return nil, err
	}
	result.UserPermissions = userPermissions

	return result, nil
}

func (s *srv) MenuOperations(f *Form) ([]MenuOperation, error) {
	menuForm, err := f.ToMenuForm()
	if err != nil {
		return nil, err
	}

	menuList, err := imenu.Srv.GetList(menuForm)
	if err != nil {
		return nil, err
	}

	tree := menutree.ToTree(menuList)
	sortedMenuList := menutree.TreeToMenuListBFS(tree)

	operationForm, _ := f.ToOperationForm()
	operationList, err := ioperation.Srv.GetList(operationForm)
	if err != nil {
		return nil, err
	}

	result := make([]MenuOperation, 0)
	for _, menu := range sortedMenuList {
		menuOperation := new(MenuOperation)
		menuOperation.Menu = menu
		menuOperation.Operations = s.filterOperations(menu.Id, operationList)
		result = append(result, *menuOperation)
	}

	return result, nil
}

func (s *srv) filterOperations(menuId int, operationList []model.Operation) []model.Operation {
	result := make([]model.Operation, 0)
	for _, operation := range operationList {
		if operation.MenuId == menuId {
			result = append(result, operation)
		}
	}
	return result
}

func (s *srv) GetAuthorizedPermissions(f *Form) (*UserPermissions, error) {
	UPForm, err := f.ToUserPermissionsForm()
	if err != nil {
		return nil, err
	}

	res, err := iuser_permissions.Srv.Query(UPForm)
	if err != nil {
		return nil, err
	}

	menuIdList := make([]int, 0)
	operationIdList := make([]int, 0)
	for _, v := range res.List {
		menuIdList = append(menuIdList, v.MenuId)

		if v.OperaIds != "" {
			menuOperaIdList := libutils.SplitToIntList(v.OperaIds, ",")
			operationIdList = append(operationIdList, menuOperaIdList...)
		}
	}

	result := new(UserPermissions)
	result.MenuIdList = menuIdList
	result.OperationIdList = operationIdList
	return result, nil
}
