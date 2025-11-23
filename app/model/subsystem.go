package model

import "time"

/*  */

type Subsystem struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                                   //
	Name       string     `json:"name" form:"name"`                                                    // 名称
	Domain     string     `json:"domain" form:"domain"`                                                // 域名
	Syskey     string     `json:"syskey" form:"syskey"`                                                // key
	Secret     string     `json:"secret" form:"secret"`                                                // secret
	Status     int        `json:"status" form:"status"`                                                // 1:有效 2:禁用
	OrderNo    int        `json:"order_no" form:"order_no"`                                            // 排序
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        // 创建人id
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        // 更新用户id
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *Subsystem) TableName() string {
	return "t_subsystem"
}
