package model

import (
	"github.com/alibabacloud-go/tea/tea"
)

type AdminUser struct {
	ID        *uint64 `gorm:"size:11;primaryKey;comment:主键;column:id" json:"id,omitempty"`
	Mobile    *string `gorm:"size:20;notNull;comment:手机号码;uniqueIndex" json:"mobile,omitempty"`
	Account   *string `gorm:"size:20;notNull;comment:账号;uniqueIndex" json:"username,omitempty"`
	Nickname  *string `gorm:"size:20;notNull;comment:昵称;default:''" json:"nickname,omitempty"`
	Password  *string `gorm:"size:200;notNull;comment:登录密码" json:"password,omitempty"`
	CreatedAt *int64  `gorm:"comment:创建时间" json:"createdAt,omitempty"`
	UpdatedAt *int64  `gorm:"comment:更新时间" json:"updatedAt,omitempty"`

	Roles []*AdminRole `gorm:"many2many:AdminUserRole;foreignKey:ID;joinForeignKey:UserID;References:ID;joinReferences:RoleID" json:"roles,omitempty"`
}

func (*AdminUser) TableName() string { return "admin_users" }

func (u *AdminUser) GetId() uint64 { return tea.Uint64Value(u.ID) }

func (u *AdminUser) GetPassword() string { return tea.StringValue(u.Password) }
