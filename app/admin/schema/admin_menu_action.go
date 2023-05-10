package schema

import (
	"erp/pkg/model"
	"github.com/alibabacloud-go/tea/tea"
)

type (
	AdminRuleActionQueryRequest struct {
		RuleId uint64 `param:"ruleId"`
	}

	AdminRuleActionCreateRequest struct {
		RuleID uint64 `param:"ruleId"`
		Method string `json:"method"`
		Path   string `json:"path"`
	}

	AdminRuleActionEditRequest struct {
		Id uint64 `param:"id"`
	}

	AdminRuleActionUpdateRequest struct {
		AdminRuleActionEditRequest
		AdminRuleActionCreateRequest
	}

	AdminRuleActionDeleteRequest struct {
		Id uint64 `param:"id"`
	}
)

func (v *AdminRuleActionCreateRequest) Validate() error { return nil }
func (v *AdminRuleActionCreateRequest) ToAdminRuleActionModel() *model.AdminRuleAction {
	return &model.AdminRuleAction{
		RuleID: tea.Uint64(v.RuleID),
		Method: tea.String(v.Method),
		Path:   tea.String(v.Path),
	}
}

func (v *AdminRuleActionEditRequest) Validate() error { return nil }

func (v *AdminRuleActionUpdateRequest) Validate() error { return nil }

func (v *AdminRuleActionDeleteRequest) Validate() error { return nil }
