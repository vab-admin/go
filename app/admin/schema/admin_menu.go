package schema

import (
	"github.com/alibabacloud-go/tea/tea"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/validate"
)

type (
	AdminRuleCreateRequest struct {
		Type        uint8    `json:"type"`
		Status      uint8    `json:"status"`
		ParentId    uint64   `json:"parentId"`
		Name        string   `json:"name"`
		Path        string   `json:"path"`
		Component   string   `json:"component"`
		Redirect    string   `json:"redirect"`
		Sort        int64    `json:"sort"`
		Title       string   `json:"title"`
		Hidden      bool     `json:"hidden"`
		LevelHidden bool     `json:"levelHidden"`
		Icon        string   `json:"icon"`
		NoKeepAlive bool     `json:"noKeepAlive"`
		NoClosable  bool     `json:"noClosable"`
		NoColumn    bool     `json:"noColumn"`
		Badge       string   `gorm:"" json:"badge"`
		TabHidden   bool     `json:"tabHidden"`
		Dot         bool     `json:"dot"`
		Target      string   `json:"target"`
		Apis        []uint64 `json:"apis"`
	}

	AdminRuleEditRequest struct {
		Id uint64 `param:"id"`
	}

	AdminRuleUpdateRequest struct {
		AdminRuleEditRequest
		AdminRuleCreateRequest
	}

	AdminRuleUpdateFieldRequest struct {
		AdminRuleEditRequest
		Field string `json:"field"`
		Value any    `json:"value"`
	}

	AdminRuleDeleteRequest struct {
		Id uint64 `param:"id"`
	}
)

var (
	adminRuleTypeRequired = validation.Required.Error("菜单类型不得为空")
	adminRuleTypeIn       = validation.In(model.AdminRuleTypeMenu, model.AdminRuleTypeAction)

	adminRuleStatusRequired = validation.Required.Error("菜单状态不得为空")
	adminRuleStatusIn       = validation.In(model.AdminRuleStatusEnable, model.AdminRuleStatusDisable)

	adminRuleNameRequired = validation.Required.Error("菜单名称不得为空")
	adminRuleNameLength   = validation.RuneLength(1, 50).Error("菜单名称不得大于50个字符")
)

// ToAdminRuleModel
// @date 2023-05-10 00:54:33
func (v *AdminRuleCreateRequest) ToAdminRuleModel() *model.AdminRule {
	return &model.AdminRule{
		Title:       tea.String(v.Title),
		Hidden:      tea.Bool(v.Hidden),
		LevelHidden: tea.Bool(v.LevelHidden),
		Icon:        tea.String(v.Icon),
		NoKeepAlive: tea.Bool(v.NoKeepAlive),
		NoClosable:  tea.Bool(v.NoClosable),
		NoColumn:    tea.Bool(v.NoColumn),
		Badge:       tea.String(v.Badge),
		TabHidden:   tea.Bool(v.TabHidden),
		Target:      tea.String(v.Target),
		Dot:         tea.Bool(v.Dot),
		ParentID:    tea.Uint64(v.ParentId),
		Path:        tea.String(v.Path),
		Name:        tea.String(v.Name),
		Component:   tea.String(v.Component),
		Redirect:    tea.String(v.Redirect),
		Sort:        tea.Int64(v.Sort),
		Status:      tea.Uint8(v.Status),
		Type:        tea.Uint8(v.Type),
	}
}

// Validate
// @date 2023-05-10 00:41:39
func (v *AdminRuleCreateRequest) Validate() error {
	return validate.Check(
		validate.Field(&v.Type, adminRuleTypeRequired, adminRuleTypeIn),
		validate.Field(&v.Status, adminRuleStatusRequired, adminRuleStatusIn),
		validate.Field(&v.Name, adminRuleNameRequired, adminRuleNameLength),
	)
}

// Validate
// @date 2023-05-10 00:41:35
func (v *AdminRuleEditRequest) Validate() error { return nil }

// Validate
// @date 2023-05-10 00:41:34
func (v *AdminRuleUpdateRequest) Validate() error {
	return validate.Check(
		validate.Field(&v.Type, adminRuleTypeRequired, adminRuleTypeIn),
		validate.Field(&v.Status, adminRuleStatusRequired, adminRuleStatusIn),
		validate.Field(&v.Name, adminRuleNameRequired, adminRuleNameLength),
	)
}

// Validate
// @date 2023-05-10 00:41:41
func (v *AdminRuleDeleteRequest) Validate() error { return nil }
