package user_permission_admin

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/imenu"
	"gofly/app/service/ioperation"
	"gofly/app/service/iuser"
	"gofly/app/service/iuser_permission"
	"gofly/app/service/iuser_permission/user_permission_def"
	"gofly/app/service/tree/menutree"
	"gofly/lib/libutils"
	"slices"
)

/*  */

type adminSrv struct{}

var (
	AdminSrv = &adminSrv{}
)

func (aSrv *adminSrv) My(f *user_permission_def.MyForm) (*user_permission_def.MyPermissions, error) {
	adminInfo, err := iuser.Srv.Get(f.OperatorUid)
	if err != nil {
		return nil, err
	}

	result := new(user_permission_def.MyPermissions)
	result.MenuActions = make([]user_permission_def.MenuActions, 0)

	if adminInfo.IsRoot == 1 {
		result.IsRoot = true
		return result, nil
	}

	qForm := new(user_permission_def.UserPermissionQueryForm)
	qForm.Page = 1
	qForm.PageSize = 1000
	qForm.SystemId = f.SystemId
	qForm.UserId = f.OperatorUid
	myPerms, err := iuser_permission.Srv.GetList(qForm)
	if err != nil {
		return nil, err
	}

	menuIdList := make([]int, 0)
	actionIdList := make([]int, 0)
	for _, v := range myPerms {
		menuIdList = append(menuIdList, v.MenuId)
		actionIdList = append(actionIdList, libutils.SplitIntList(v.OperaIds, ",")...)
	}
	menuMap, _ := imenu.Srv.GetMulti(menuIdList)
	actionMap, _ := ioperation.Srv.GetMulti(actionIdList)
	for _, v := range myPerms {
		menu := menuMap[v.MenuId]

		var menuActions user_permission_def.MenuActions
		menuActions.Id = menu.Id
		menuActions.Url = menu.Url
		menuActions.Name = menu.Name
		menuActions.NodePath = menu.NodePath

		if v.OperaIds != "" {
			operaIds := libutils.SplitIntList(v.OperaIds, ",")
			menuActions.Actions = make([]string, 0)
			for _, v := range operaIds {
				action := actionMap[v]
				menuActions.Actions = append(menuActions.Actions, action.Code)
			}
		}
		result.MenuActions = append(result.MenuActions, menuActions)
	}

	return result, nil
}
func (aSrv *adminSrv) Config(f *user_permission_def.ConfigForm) (*user_permission_def.MenuPermission, error) {
	result := new(user_permission_def.MenuPermission)
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

func (aSrv *adminSrv) MenuOperations(f *user_permission_def.ConfigForm) ([]user_permission_def.MenuOperation, error) {
	menuForm, err := f.ToMenuForm()
	if err != nil {
		return nil, err
	}

	menuList, err := imenu.Srv.GetList(menuForm)
	if err != nil {
		return nil, err
	}

	tree := menutree.ToTree(menuList)
	sortedMenuList := menutree.TreeToMenuListDFS(tree)
	//sortedMenuList = slices.Clip(sortedMenuList)

	operationForm, _ := f.ToOperationForm()
	operationList, err := ioperation.Srv.GetList(operationForm)
	if err != nil {
		return nil, err
	}

	result := make([]user_permission_def.MenuOperation, 0)
	for _, menu := range sortedMenuList {
		menuOperation := new(user_permission_def.MenuOperation)
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

func (aSrv *adminSrv) GetAuthorizedPermissions(f *user_permission_def.ConfigForm) (*user_permission_def.UserPermissions, error) {
	UPForm, err := f.ToUserPermissionsForm()
	if err != nil {
		return nil, err
	}

	res, err := iuser_permission.Srv.Query(UPForm)
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

	result := new(user_permission_def.UserPermissions)
	result.MenuIdList = menuIdList
	result.OperationIdList = operationIdList
	return result, nil
}

func (aSrv *adminSrv) Save(f *user_permission_def.SetPermissionForm) (int64, error) {
	menuOperations, err := aSrv.MenuOperations(&f.ConfigForm)
	if err != nil {
		return 0, err
	}

	moMap := make(map[int]user_permission_def.MenuOperation)
	for _, v := range menuOperations {
		moMap[v.Menu.Id] = v
	}

	userPermissions := make([]*model.UserPermission, 0)
	for _, v := range f.MenuIdList {
		mo := moMap[v]

		menuOperationIds := make([]int, 0)
		for _, op := range mo.Operations {
			menuOperationIds = append(menuOperationIds, op.Id)
		}

		row := new(model.UserPermission)
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

	_, _ = iuser_permission.Srv.ClearUserPermission(f.UserId, f.SystemId)
	return iuser_permission.Srv.BatchInsert(userPermissions, 100)
}

func (aSrv *adminSrv) CheckMenuPermission(userId int, menuId int) (bool, error) {
	userInfo, err := iuser.Srv.Get(userId)
	if err != nil {
		return false, err
	}
	if userInfo.IsRoot == 1 {
		return true, nil
	}

	pForm := new(user_permission_def.ConfigForm)
	pForm.UserId = userId
	pForm.SystemId = 1
	permissions, err := aSrv.GetAuthorizedPermissions(pForm)
	if err != nil {
		return false, err
	}
	return slices.Contains(permissions.MenuIdList, menuId), nil
}
func (aSrv *adminSrv) CheckOperationPermission(userId int, operationId int) (bool, error) {
	userInfo, err := iuser.Srv.Get(userId)
	if err != nil {
		return false, err
	}
	if userInfo.IsRoot == 1 {
		return true, nil
	}

	pForm := new(user_permission_def.ConfigForm)
	pForm.UserId = userId
	pForm.SystemId = 1
	permissions, err := aSrv.GetAuthorizedPermissions(pForm)
	if err != nil {
		return false, err
	}
	return slices.Contains(permissions.OperationIdList, operationId), nil
}
