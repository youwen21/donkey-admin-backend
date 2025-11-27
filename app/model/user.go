package model

import "time"

/*  */

type StaffInfo struct {
	Id     int    `json:"id" form:"id" gorm:"autoIncrement"` //
	Name   string `json:"name" form:"name"`                  // 登陆名
	Avatar string `json:"avatar" form:"avatar"`              // 用户头像

	RealName string `json:"real_name" form:"real_name"` // 真实名字
	Email    string `json:"email" form:"email"`         // 员工邮箱
	Phone    string `json:"phone" form:"phone"`         // 员工手机号

	RoleId int `json:"role_id" form:"role_id"` // 角色id
	OrgId  int `json:"org_id" form:"org_id"`   // 所属组织

	IsRoot  int8 `json:"is_root" form:"is_root"`   // 是否root用户
	IsStaff int8 `json:"is_staff" form:"is_staff"` // 是否内部员工
	StaffNo int  `json:"staff_no" form:"staff_no"` // 员工号
	Status  int  `json:"status" form:"status"`     // 是否在职，1:在职，0:离职
}

type User struct {
	StaffInfo
	Password   string     `json:"password" form:"password"`                                            // 密码
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        //
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        //
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *User) TableName() string {
	return "t_user"
}
