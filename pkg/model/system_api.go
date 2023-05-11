package model

import "github.com/alibabacloud-go/tea/tea"

type (
	SystemApi struct {
		ID        *uint64 `gorm:"size:11;primaryKey;comment:主键;column:id" json:"id,omitempty"`
		Name      *string `gorm:"size:20" json:"name,omitempty"`
		Method    *string `gorm:"size:50;" json:"method,omitempty"` // 资源请求方式(支持正则)
		Path      *string `gorm:"size:255;" json:"path,omitempty"`  // 资源请求路径（支持/:id匹配）
		CreatedAt *int64  `json:"createdAt,omitempty"`
	}
)

func (*SystemApi) TableName() string   { return "system_api" }
func (m *SystemApi) GetMethod() string { return tea.StringValue(m.Method) }
func (m *SystemApi) GetPath() string   { return tea.StringValue(m.Path) }
