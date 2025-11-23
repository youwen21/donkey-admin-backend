package model

import "time"

/*  */

type UserRole struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                    //
	UserId     int        `json:"user_id" form:"user_id"`                               // 角色id
	RoleId     int        `json:"role_id" form:"role_id"`                               // 所属组织
	CreateUid  int        `json:"create_uid" form:"create_uid"`                         //
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"` //
}

// 自定义表名
func (m *UserRole) TableName() string {
	return "t_user_role"
}
