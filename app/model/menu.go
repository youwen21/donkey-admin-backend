package model

import "time"

/*  */

type Menu struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                                   //
	SystemId   int        `json:"system_id" form:"system_id"`                                          // 子系统ID
	ParentId   int        `json:"parent_id" form:"parent_id"`                                          // 上级菜单ID
	Name       string     `json:"name" form:"name"`                                                    // 菜单名
	Level      int        `json:"level" form:"level"`                                                  // 菜单级别
	NodePath   string     `json:"node_path" form:"node_path"`                                          // 菜单路径
	Url        string     `json:"url" form:"url"`                                                      // url
	Status     int        `json:"status" form:"status"`                                                // 1:有效 2:禁用
	OrderNo    int        `json:"order_no" form:"order_no"`                                            // 菜单排序
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        //
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        //
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *Menu) TableName() string {
	return "t_menu"
}
