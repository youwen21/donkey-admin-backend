package model

import "time"

/*  */

type Organization struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                                   //
	ParentId   int        `json:"parent_id" form:"parent_id"`                                          // 上级组织ID
	Name       string     `json:"name" form:"name"`                                                    // 组织名称
	Level      int        `json:"level" form:"level"`                                                  // 组织级别
	NodePath   string     `json:"node_path" form:"node_path"`                                          // 组织节点路径
	Status     int        `json:"status" form:"status"`                                                // 状态 1启用 2禁用 0删除
	OrderNo    int        `json:"order_no" form:"order_no"`                                            // 序号
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        //
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        //
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *Organization) TableName() string {
	return "t_organization"
}
