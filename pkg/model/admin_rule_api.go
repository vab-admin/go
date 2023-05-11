package model

type AdminRuleApi struct {
	RuleID uint64
	ApiID  uint64
}

func (*AdminRuleApi) TableName() string {
	return "admin_rule_apis"
}
