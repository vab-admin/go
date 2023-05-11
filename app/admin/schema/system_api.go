package schema

import (
	"github.com/alibabacloud-go/tea/tea"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/pagination"
)

type (
	SystemApiQueryRequest struct {
		pagination.Param
		Name string `json:"name"`
	}

	SystemApiCreateRequest struct {
		Method string `json:"method"`
		Name   string `json:"name"`
		Path   string `json:"path"`
	}

	SystemApiEditRequest struct {
		Id uint64 `param:"id"`
	}

	SystemApiUpdateRequest struct {
		SystemApiEditRequest
		SystemApiCreateRequest
	}

	SystemApiDeleteRequest struct {
		Id uint64 `param:"id"`
	}
)

func (v *SystemApiCreateRequest) Validate() error { return nil }

func (v *SystemApiCreateRequest) ToSystemApiModel() *model.SystemApi {
	return &model.SystemApi{
		Name:   tea.String(v.Name),
		Method: tea.String(v.Method),
		Path:   tea.String(v.Path),
	}
}

func (v *SystemApiEditRequest) Validate() error { return nil }

func (v *SystemApiUpdateRequest) Validate() error { return nil }

func (v *SystemApiDeleteRequest) Validate() error { return nil }
