package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"vab-admin/go/pkg/validate"
)

type (
	AdminRoleQueryRequest struct {
		Name string `query:"name"`
	}

	AdminRoleCreateRequest struct {
		Name  string   `json:"name"`
		Rules []uint64 `json:"rules"`
	}

	AdminRoleEditRequest struct {
		Id uint64 `param:"id"`
	}

	AdminRoleUpdateRequest struct {
		AdminRoleEditRequest
		AdminRoleCreateRequest
	}

	AdminRoleDeleteRequest struct {
		Id uint64 `param:"id"`
	}
)

var (
	adminRoleNameRequired = validation.Required.Error("管理员角色名称不得为空")
	adminRoleNameLength   = validation.RuneLength(1, 20).Error("管理员角色名称长度应该在1-20位")
)

func (v *AdminRoleCreateRequest) Validate() error {
	return validate.Check(validate.Field(&v.Name, adminRoleNameRequired, adminRoleNameLength))
}

func (v *AdminRoleEditRequest) Validate() error { return nil }

func (v *AdminRoleUpdateRequest) Validate() error {
	return validate.Check(validate.Field(&v.Name, adminRoleNameRequired, adminRoleNameLength))
}

func (v *AdminRoleDeleteRequest) Validate() error { return nil }
