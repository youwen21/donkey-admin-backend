package menu_def

import (
	"gofly/app/model"
	"gofly/lib/libdto"
)

/*  */

type MenuQueryForm struct {
	model.Menu
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type MenuQueryRes struct {
	Total int64        `json:"total" form:"total"`
	List  []model.Menu `json:"list" form:"list"`
}

type MenuExDTO struct {
	model.Menu

	// extend
}

type MenuQueryResEx struct {
	Total int64       `json:"total" form:"total"`
	List  []MenuExDTO `json:"list" form:"list"`
}
