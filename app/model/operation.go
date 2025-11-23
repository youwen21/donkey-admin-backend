package model

import "time"

/*  */

type Operation struct {
	Id         int        `json:"id" form:"id" gorm:"autoIncrement"`                                   //
	SystemId   int        `json:"system_id" form:"system_id"`                                          // 子系统ID
	MenuId     int        `json:"menu_id" form:"menu_id"`                                              // 菜单ID
	Name       string     `json:"name" form:"name"`                                                    // 操作名称
	Code       string     `json:"code" form:"code"`                                                    // 操作编号,此编号对应页面class, 用来控制按钮是否显示，也对应后台接口操作编号，校验用户是否有此操作权限
	Status     int        `json:"status" form:"status"`                                                // 1:有效 2:禁用
	OrderNo    int        `json:"order_no" form:"order_no"`                                            // 排序
	CreateUid  int        `json:"create_uid" form:"create_uid"`                                        // 创建人id
	CreateTime *time.Time `json:"create_time" form:"create_time" gorm:"autoCreateTime"`                //
	UpdateUid  int        `json:"update_uid" form:"update_uid"`                                        // 更新用户id
	UpdateTime *time.Time `json:"update_time" form:"update_time" gorm:"autoCreateTime;autoUpdateTime"` //
}

// 自定义表名
func (m *Operation) TableName() string {
	return "t_operation"
}
