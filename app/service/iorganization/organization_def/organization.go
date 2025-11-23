package organization_def

import (
	"gofly/app/model"
	"gofly/lib/libdto"
)

/*  */

type OrganizationQueryForm struct {
	model.Organization
	libdto.PageForm
	libdto.OrderForm

	SearchKey string `json:"search_key" form:"search_key"`
	Ids       string `json:"ids" form:"ids"`
	IdList    []int  `json:"id_list" form:"id_list"`
}

type OrganizationQueryRes struct {
	Total int64                `json:"total" form:"total"`
	List  []model.Organization `json:"list" form:"list"`
}

type OrganizationExDTO struct {
	model.Organization

	// extend
}

type OrganizationQueryResEx struct {
	Total int64               `json:"total" form:"total"`
	List  []OrganizationExDTO `json:"list" form:"list"`
}
