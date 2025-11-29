package role_def

import (
	"donkey-admin/app/model"
	"donkey-admin/lib/libdto"
)

/*  */

type RoleQueryForm struct {
	model.Role
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type RoleQueryRes struct {
	Total int64        `json:"total" form:"total"`
	List  []model.Role `json:"list" form:"list"`
}

type RoleExDTO struct {
	model.Role

	// extend
}

type RoleQueryResEx struct {
	Total int64       `json:"total" form:"total"`
	List  []RoleExDTO `json:"list" form:"list"`
}
