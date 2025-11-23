package menu_permission

import (
	"errors"
	"gofly/app/model"
	"gofly/app/service/imenu/menu_def"
	"gofly/app/service/ioperation/operation_def"
	"gofly/app/service/iuser_permissions/user_permissions_def"
)

type Form struct {
	SystemId int `json:"system_id" form:"system_id"`
	UserId   int `json:"user_id" form:"user_id"`
}

func (f *Form) ToMenuForm() (*menu_def.MenuQueryForm, error) {
	if f.SystemId == 0 {
		return nil, errors.New("system_id is required")
	}

	menuForm := new(menu_def.MenuQueryForm)
	menuForm.SystemId = f.SystemId
	menuForm.Page = 1
	menuForm.PageSize = 1000
	return menuForm, nil
}

func (f *Form) ToOperationForm() (*operation_def.OperationQueryForm, error) {
	if f.SystemId == 0 {
		return nil, errors.New("system_id is required")
	}
	operaForm := new(operation_def.OperationQueryForm)
	operaForm.SystemId = f.SystemId
	operaForm.Page = 1
	operaForm.PageSize = 1000
	return operaForm, nil
}

func (f *Form) ToUserPermissionsForm() (*user_permissions_def.UserPermissionsQueryForm, error) {
	if f.SystemId == 0 {
		return nil, errors.New("system_id is required")
	}

	if f.UserId == 0 {
		return nil, errors.New("user_id is required")
	}

	UPForm := new(user_permissions_def.UserPermissionsQueryForm)
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

type MenuPermission struct {
	Form            *Form            `json:"form"`
	IsAdmin         bool             `json:"is_admin"`
	SystemMenu      []MenuOperation  `json:"system_menu"`
	UserPermissions *UserPermissions `json:"user_permissions"`
}
