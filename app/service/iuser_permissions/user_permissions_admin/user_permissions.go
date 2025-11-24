package user_permissions_admin

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/ioperation"
	"gofly/app/service/iuser"
	"gofly/app/service/iuser_permissions"
	"gofly/app/service/iuser_permissions/user_permissions_def"
	"gofly/app/service/tree/menutree"
	"gofly/lib/libutils"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) Detail(f *user_permissions_def.Form) (*user_permissions_def.MenuPermission, error) {
	result := new(user_permissions_def.MenuPermission)
	result.Form = f

	menuOperations, err := aSrv.MenuOperations(f)
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
		result.IsRoot = true
		return result, nil
	}

	// 用户权限
	userPermissions, err := aSrv.GetAuthorizedPermissions(f)
	if err != nil {
		return nil, err
	}
	result.UserPermissions = userPermissions

	return result, nil
}

func (aSrv *adminSrv) MenuOperations(f *user_permissions_def.Form) ([]user_permissions_def.MenuOperation, error) {
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

	result := make([]user_permissions_def.MenuOperation, 0)
	for _, menu := range sortedMenuList {
		menuOperation := new(user_permissions_def.MenuOperation)
		menuOperation.Menu = menu
		menuOperation.Operations = aSrv.filterOperations(menu.Id, operationList)
		result = append(result, *menuOperation)
	}

	return result, nil
}

func (aSrv *adminSrv) filterOperations(menuId int, operationList []model.Operation) []model.Operation {
	result := make([]model.Operation, 0)
	for _, operation := range operationList {
		if operation.MenuId == menuId {
			result = append(result, operation)
		}
	}
	return result
}

func (aSrv *adminSrv) GetAuthorizedPermissions(f *user_permissions_def.Form) (*user_permissions_def.UserPermissions, error) {
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
			menuOperaIdList := libutils.SplitIntList(v.OperaIds, ",")
			operationIdList = append(operationIdList, menuOperaIdList...)
		}
	}

	result := new(user_permissions_def.UserPermissions)
	result.MenuIdList = menuIdList
	result.OperationIdList = operationIdList
	return result, nil
}

func (aSrv *adminSrv) SetPermission(f *user_permissions_def.SetPermissionForm) (int64, error) {
	menuOperations, err := aSrv.MenuOperations(&f.Form)
	if err != nil {
		return 0, err
	}

	moMap := make(map[int]user_permissions_def.MenuOperation)
	for _, v := range menuOperations {
		moMap[v.Menu.Id] = v
	}

	userPermissions := make([]*model.UserPermissions, 0)
	for _, v := range f.MenuIdList {
		mo := moMap[v]

		menuOperationIds := make([]int, 0)
		for _, op := range mo.Operations {
			menuOperationIds = append(menuOperationIds, op.Id)
		}

		row := new(model.UserPermissions)
		row.UserId = f.UserId
		row.SystemId = f.SystemId
		row.MenuId = v
		row.CreateUid = f.OperatorUid

		if len(menuOperationIds) > 0 && len(f.OperationIdList) > 0 {
			// 授权的operations
			intersection := libutils.SlicesUtil.IntersectionIntSlice(menuOperationIds, f.OperationIdList)
			row.OperaIds = libutils.JoinIntList(intersection, ",")
		}

		userPermissions = append(userPermissions, row)
	}

	_, _ = iuser_permissions.Srv.ClearUserPermissions(f.UserId, f.SystemId)
	return iuser_permissions.Srv.BatchInsert(userPermissions, 100)
}
