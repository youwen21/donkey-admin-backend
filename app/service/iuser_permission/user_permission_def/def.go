package user_permission_def

import (
	"donkey-admin/app/model"
	"donkey-admin/app/service/imenu/menu_def"
	"donkey-admin/app/service/ioperation/operation_def"
	"errors"
)

type MyForm struct {
	SystemId    int `json:"system_id" form:"system_id"`
	OperatorUid int // 登录的用户ID
}

type ConfigForm struct {
	SystemId int `json:"system_id" form:"system_id"`
	UserId   int `json:"user_id" form:"user_id"` // 页面中 指定的用户ID
}

func (f *ConfigForm) ToMenuForm() (*menu_def.MenuQueryForm, error) {
	if f.SystemId == 0 {
		return nil, errors.New("system_id is required")
	}

	menuForm := new(menu_def.MenuQueryForm)
	menuForm.SystemId = f.SystemId
	menuForm.Page = 1
	menuForm.PageSize = 2000
	return menuForm, nil
}

func (f *ConfigForm) ToOperationForm() (*operation_def.OperationQueryForm, error) {
	if f.SystemId == 0 {
		return nil, errors.New("system_id is required")
	}
	operaForm := new(operation_def.OperationQueryForm)
	operaForm.SystemId = f.SystemId
	operaForm.Page = 1
	operaForm.PageSize = 2000
	return operaForm, nil
}

func (f *ConfigForm) ToUserPermissionsForm() (*UserPermissionQueryForm, error) {
	if f.SystemId == 0 {
		return nil, errors.New("system_id is required")
	}

	if f.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	UPForm := new(UserPermissionQueryForm)
	UPForm.SystemId = f.SystemId
	UPForm.UserId = f.UserId
	UPForm.Page = 1
	UPForm.PageSize = 1000
	return UPForm, nil
}

type MenuOperation struct {
	model.Menu

	Operations []model.Operation `json:"operations"`
}

type UserPermissions struct {
	MenuIdList      []int `json:"menu_id_list"`
	OperationIdList []int `json:"operation_id_list"`
}

// Unique 过滤MenuIdList和OperationIdList中的重复值，确保列表中每个ID都是唯一的
func (up *UserPermissions) Unique() {
	up.MenuIdList = uniqueIntSlice(up.MenuIdList)
	up.OperationIdList = uniqueIntSlice(up.OperationIdList)
}

// uniqueIntSlice 过滤整数切片中的重复值
func uniqueIntSlice(slice []int) []int {
	if len(slice) == 0 {
		return slice
	}

	// 使用map记录已存在的元素
	seen := make(map[int]bool)
	result := make([]int, 0, len(slice))

	for _, item := range slice {
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

type MenuPermission struct {
	Form            *ConfigForm      `json:"form"`
	IsRoot          bool             `json:"is_root"`
	SystemMenu      []MenuOperation  `json:"system_menu"`
	UserPermissions *UserPermissions `json:"user_permissions"`
}

type SetPermissionForm struct {
	ConfigForm
	UserPermissions

	OperatorUid int
}

type MenuActions struct {
	Id       int      `json:"id"`  // menuId
	Url      string   `json:"url"` // 菜单的URL
	Name     string   `json:"name"`
	NodePath string   `json:"node_path"`
	Actions  []string `json:"actions"`
}

type MyPermissions struct {
	IsRoot      bool          `json:"isRoot"`
	MenuActions []MenuActions `json:"menuActions"`
}
