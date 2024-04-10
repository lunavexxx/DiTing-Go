// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameGroupMember = "group_member"

// GroupMember 群成员表
type GroupMember struct {
	ID         int64     `gorm:"column:id;primaryKey;autoIncrement:true;comment:id" json:"id"`                             // id
	GroupID    int64     `gorm:"column:group_id;not null;comment:群主id" json:"group_id"`                                    // 群主id
	UID        int64     `gorm:"column:uid;not null;comment:成员uid" json:"uid"`                                             // 成员uid
	Role       int32     `gorm:"column:role;not null;comment:成员角色 1群主 2管理员 3普通成员" json:"role"`                             // 成员角色 1群主 2管理员 3普通成员
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP(3);comment:创建时间" json:"create_time"` // 创建时间
	UpdateTime time.Time `gorm:"column:update_time;not null;default:CURRENT_TIMESTAMP(3);comment:修改时间" json:"update_time"` // 修改时间
}

// TableName GroupMember's table name
func (*GroupMember) TableName() string {
	return TableNameGroupMember
}