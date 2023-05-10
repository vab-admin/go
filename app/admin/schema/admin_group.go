package schema

import (
	"erp/pkg/validate"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	AdminGroupQueryRequest struct {
	}

	AdminGroupCreateRequest struct {
		Name  string   `json:"name"`
		Rules []uint64 `json:"rules"`
	}

	AdminGroupEditRequest struct {
		Id uint64 `param:"id"`
	}

	AdminGroupUpdateRequest struct {
		AdminGroupEditRequest
		AdminGroupCreateRequest
	}

	AdminGroupDeleteRequest struct {
		Id uint64 `param:"id"`
	}
)

var (
	adminGroupNameRequired = validation.Required.Error("管理员分组名称不得为空")
	adminGroupNameLength   = validation.RuneLength(1, 20).Error("管理员分组名称长度应该在1-20位")
)

func (v *AdminGroupCreateRequest) Validate() error {
	return validate.Check(validate.Field(&v.Name, adminGroupNameRequired, adminGroupNameLength))
}

func (v *AdminGroupEditRequest) Validate() error { return nil }

func (v *AdminGroupUpdateRequest) Validate() error {
	return validate.Check(validate.Field(&v.Name, adminGroupNameRequired, adminGroupNameLength))
}

func (v *AdminGroupDeleteRequest) Validate() error { return nil }
