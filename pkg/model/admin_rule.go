package model

import (
	"github.com/alibabacloud-go/tea/tea"
)

const (
	AdminRuleTypeMenu   uint8 = 1
	AdminRuleTypeAction uint8 = 2
)

const (
	AdminRuleStatusEnable  uint8 = 1
	AdminRuleStatusDisable uint8 = 2
)

type (
	AdminRule struct {
		ID          *uint64      `gorm:"size:11;primaryKey;comment:主键;column:id" json:"id,omitempty"`
		ParentID    *uint64      `gorm:"comment:上级id;default:0" json:"parentId,omitempty"`
		Path        *string      `gorm:"size:200;comment:路由" json:"path,omitempty"`
		Name        *string      `gorm:"size:100;uniqueIndex" json:"name,omitempty"`
		Component   *string      `gorm:"size:100;default:null" json:"component,omitempty"`
		Redirect    *string      `gorm:"default:null;comment:重定向到子路由" json:"redirect,omitempty"`
		Title       *string      `gorm:"size:50" json:"title,omitempty"`
		Hidden      *bool        `gorm:"default:false" json:"hidden,omitempty"`
		LevelHidden *bool        `gorm:"default:false" json:"levelHidden,omitempty"`
		Icon        *string      `gorm:"size:20" json:"icon,omitempty"`
		NoKeepAlive *bool        `gorm:"default:false" json:"noKeepAlive,omitempty"`
		NoClosable  *bool        `gorm:"default:false" json:"noClosable,omitempty"`
		NoColumn    *bool        `gorm:"default:false" json:"noColumn,omitempty"`
		Badge       *string      `gorm:"" json:"badge,omitempty"`
		TabHidden   *bool        `gorm:"default:false" json:"tabHidden,omitempty"`
		Target      *string      `gorm:"" json:"target,omitempty"`
		Dot         *bool        `gorm:"default:false" json:"dot,omitempty"`
		Sort        *int64       `gorm:"default:0" json:"sort,omitempty"`
		Status      *uint8       `gorm:"default:1;comment:是否启用" json:"status,omitempty"`
		Type        *uint8       `gorm:"default:1;comment:菜单类型,0:公开,1:鉴权" json:"type,omitempty"`
		Children    []*AdminRule `gorm:"-" json:"children,omitempty"`

		Apis []*SystemApi `gorm:"many2many:AdminRuleApi;foreignKey:ID;joinForeignKey:RuleID;References:ID;joinReferences:ApiID" json:"apis,omitempty"`
	}
)

func (*AdminRule) TableName() string     { return "admin_rules" }
func (m *AdminRule) GetParentId() uint64 { return tea.Uint64Value(m.ParentID) }
func (m *AdminRule) GetId() uint64       { return tea.Uint64Value(m.ID) }
func (m *AdminRule) GetType() uint8      { return tea.Uint8Value(m.Type) }
