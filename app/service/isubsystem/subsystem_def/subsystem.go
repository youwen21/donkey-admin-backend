package subsystem_def

import (
	"donkey-admin/app/model"
	"donkey-admin/lib/libdto"
)

/*  */

type SubsystemQueryForm struct {
	model.Subsystem
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type SubsystemQueryRes struct {
	Total int64             `json:"total" form:"total"`
	List  []model.Subsystem `json:"list" form:"list"`
}

type SubsystemExDTO struct {
	model.Subsystem

	// extend
}

type SubsystemQueryResEx struct {
	Total int64            `json:"total" form:"total"`
	List  []SubsystemExDTO `json:"list" form:"list"`
}
