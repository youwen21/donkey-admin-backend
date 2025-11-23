package menutree

import "gofly/app/model"

type UserMenuForm struct {
	SystemId    int `json:"system_id,default=1" form:"system_id,default=1"`
	AdminUserId int
}

type TreeMenu struct {
	model.Menu
	Children []*TreeMenu `json:"children"`
}
