package user_permissions_def

import (
	"gofly/app/model"
	"gofly/lib/libdto"
)

/*  */

type UserPermissionsQueryForm struct {
	model.UserPermissions
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type UserPermissionsQueryRes struct {
	Total int64                   `json:"total" form:"total"`
	List  []model.UserPermissions `json:"list" form:"list"`
}

type UserPermissionsExDTO struct {
	model.UserPermissions

	// extend
}

type UserPermissionsQueryResEx struct {
	Total int64                  `json:"total" form:"total"`
	List  []UserPermissionsExDTO `json:"list" form:"list"`
}
