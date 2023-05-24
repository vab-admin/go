package model

type AdminRoleRule struct {
	RoleID uint64 `gorm:"comment:分组ID" json:"roleId,omitempty"`
	RuleID uint64 `gorm:"comment:规则ID" json:"ruleId,omitempty"`
}
