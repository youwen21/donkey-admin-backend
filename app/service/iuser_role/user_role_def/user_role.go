package user_role_def

import (
	"donkey-admin/app/model"
	"donkey-admin/lib/libdto"
)

/*  */

type UserRoleQueryForm struct {
	model.UserRole
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type UserRoleQueryRes struct {
	Total int64            `json:"total" form:"total"`
	List  []model.UserRole `json:"list" form:"list"`
}

type UserRoleExDTO struct {
	model.UserRole

	// extend
}

type UserRoleQueryResEx struct {
	Total int64           `json:"total" form:"total"`
	List  []UserRoleExDTO `json:"list" form:"list"`
}
