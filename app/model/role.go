package model

import "time"

/*  */

type Role struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                                   //
	Name       string     `json:"name" form:"name"`                                                    // 角色名称
	Status     int        `json:"status" form:"status"`                                                // 角色状态：1、启用，2、禁用
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        //
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        //
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *Role) TableName() string {
	return "t_role"
}
