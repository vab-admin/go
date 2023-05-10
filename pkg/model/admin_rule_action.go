package model

type (
	AdminRuleAction struct {
		ID        *uint64 `gorm:"size:11;primaryKey;comment:主键;column:id" json:"id,omitempty"`
		RuleID    *uint64 `gorm:"notNull" json:"ruleId,omitempty"`
		Method    *string `gorm:"size:50;" json:"method"` // 资源请求方式(支持正则)
		Path      *string `gorm:"size:255;" json:"path"`  // 资源请求路径（支持/:id匹配）
		CreatedAt *int64  `json:"createdAt,omitempty"`
	}
)

func (*AdminRuleAction) TableName() string { return "admin_rule_action" }
