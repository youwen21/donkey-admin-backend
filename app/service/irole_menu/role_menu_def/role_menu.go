package role_menu_def

import (
	"gofly/app/model"
	"gofly/lib/libdto"
)

/*  */

type RoleMenuQueryForm struct {
	model.RoleMenu
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`

	RoleIdList []int
}

type RoleMenuQueryRes struct {
	Total int64            `json:"total" form:"total"`
	List  []model.RoleMenu `json:"list" form:"list"`
}

type RoleMenuExDTO struct {
	model.RoleMenu

	// extend
}

type RoleMenuQueryResEx struct {
	Total int64           `json:"total" form:"total"`
	List  []RoleMenuExDTO `json:"list" form:"list"`
}
