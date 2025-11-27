package model

import "time"

/*  */

type UserPermission struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                                   //
	UserId     int        `json:"user_id" form:"user_id"`                                              // 角色ID
	SystemId   int        `json:"system_id" form:"system_id"`                                          // 系统id
	MenuId     int        `json:"menu_id" form:"menu_id"`                                              // 菜单id
	OperaIds   string     `json:"opera_ids" form:"opera_ids"`                                          // 菜单下的可用操作
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        // 创建人id
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        // 更新用户id
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *UserPermission) TableName() string {
	return "t_user_permission"
}
