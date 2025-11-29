package operation_def

import (
	"donkey-admin/app/model"
	"donkey-admin/lib/libdto"
)

/*  */

type OperationQueryForm struct {
	model.Operation
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type OperationQueryRes struct {
	Total int64             `json:"total" form:"total"`
	List  []model.Operation `json:"list" form:"list"`
}

type OperationExDTO struct {
	model.Operation

	// extend
}

type OperationQueryResEx struct {
	Total int64            `json:"total" form:"total"`
	List  []OperationExDTO `json:"list" form:"list"`
}
