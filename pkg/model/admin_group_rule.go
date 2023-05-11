package model

type AdminGroupRule struct {
	GroupID uint64 `gorm:"comment:分组ID" json:"groupId,omitempty"`
	RuleID  uint64 `gorm:"comment:规则ID" json:"ruleId,omitempty"`
}
