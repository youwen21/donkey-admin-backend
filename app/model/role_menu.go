package model

/*  */

type RoleMenu struct {
	Id       int    `json:"id" form:"id" gorm:"autoIncrement"` //
	RoleId   int    `json:"role_id" form:"role_id"`            // 角色ID
	SystemId int    `json:"system_id" form:"system_id"`        // 系统id
	MenuId   int    `json:"menu_id" form:"menu_id"`            // 菜单id
	OperaIds string `json:"opera_ids" form:"opera_ids"`        // 菜单下的可用操作
}

// 自定义表名
func (m *RoleMenu) TableName() string {
	return "t_role_menu"
}
