package model

import "github.com/alibabacloud-go/tea/tea"

type AdminGroup struct {
	ID        *uint64 `gorm:"size:11;primaryKey;comment:主键;column:id" json:"id,omitempty"`
	Name      *string `gorm:"size:20;notNull;comment:分组名称" json:"name,omitempty"`
	CreatedAt *int64  `gorm:"comment:创建时间" json:"createdAt,omitempty"`
	UpdatedAt *int64  `gorm:"comment:更新时间" json:"updatedAt,omitempty"`

	Rules []*AdminRule `gorm:"many2many:AdminGroupRule;foreignKey:ID;joinForeignKey:GroupID;References:ID;joinReferences:RuleID" json:"rules,omitempty"`
}

func (*AdminGroup) TableName() string { return "admin_groups" }

func (g *AdminGroup) GetId() uint64 { return tea.Uint64Value(g.ID) }
