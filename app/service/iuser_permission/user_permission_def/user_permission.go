package user_permission_def

import (
	"gofly/app/model"
	"gofly/lib/libdto"
)

/*  */

type UserPermissionQueryForm struct {
	model.UserPermission
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type UserPermissionQueryRes struct {
	Total int64                  `json:"total" form:"total"`
	List  []model.UserPermission `json:"list" form:"list"`
}

type UserPermissionExDTO struct {
	model.UserPermission

	// extend
}

type UserPermissionQueryResEx struct {
	Total int64                 `json:"total" form:"total"`
	List  []UserPermissionExDTO `json:"list" form:"list"`
}
