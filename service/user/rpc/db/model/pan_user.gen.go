// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNamePanUser = "pan_user"

// PanUser mapped from table <pan_user>
type PanUser struct {
	ID         int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true" json:"id"`
	Email      string    `gorm:"column:email;type:varchar(32);not null" json:"email"`        // 用户邮箱
	Username   string    `gorm:"column:username;type:varchar(20);not null" json:"username"`  // 用户名称
	Password   string    `gorm:"column:password;type:varchar(32);not null" json:"password"`  // 用户密码
	Status     int32     `gorm:"column:status;type:tinyint unsigned;not null" json:"status"` // 用户状态
	Roles      string    `gorm:"column:roles;type:varchar(64);not null" json:"roles"`        // 用户身份
	Ticket     string    `gorm:"column:ticket;type:varchar(6);not null" json:"ticket"`
	LastActive time.Time `gorm:"column:last_active;type:timestamp" json:"last_active"`    // 最后活跃时间戳
	Created    time.Time `gorm:"column:created;type:datetime(3);not null" json:"created"` // 创建时间
	Updated    time.Time `gorm:"column:updated;type:datetime(3);not null" json:"updated"` // 更新时间
	Deleted    time.Time `gorm:"column:deleted;type:datetime(3)" json:"deleted"`          // 删除时间
}

// TableName PanUser's table name
func (*PanUser) TableName() string {
	return TableNamePanUser
}