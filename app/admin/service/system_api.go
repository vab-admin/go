package service

import (
	"context"
	"github.com/alibabacloud-go/tea/tea"
	"time"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/pagination"
)

type SystemApi struct {
	SystemApiRepo *repository.SystemApi
}

// Query
// @param ctx
// @param req
// @date 2023-05-10 02:23:44
func (l *SystemApi) Query(ctx context.Context, req *schema.SystemApiQueryRequest) (*pagination.Paginator[[]*model.SystemApi], error) {

	return l.SystemApiRepo.Query(ctx, req)
}

// Create
// @param ctx
// @param req
// @date 2023-05-10 02:23:42
func (l *SystemApi) Create(ctx context.Context, req *schema.SystemApiCreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	m := req.ToSystemApiModel()
	m.CreatedAt = tea.Int64(time.Now().Unix())
	return l.SystemApiRepo.Create(ctx, m)
}

// Edit
// @param ctx
// @param req
// @date 2023-05-10 02:23:41
func (l *SystemApi) Edit(ctx context.Context, req *schema.SystemApiEditRequest) (*model.SystemApi, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return l.SystemApiRepo.Edit(ctx, req.Id)
}

// Update
// @param ctx
// @param req
// @date 2023-05-10 02:23:39
func (l *SystemApi) Update(ctx context.Context, req *schema.SystemApiUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	m := req.ToSystemApiModel()
	return l.SystemApiRepo.Update(ctx, req.Id, m)
}

// Delete
// @param ctx
// @param req
// @date 2023-05-10 02:23:38
func (l *SystemApi) Delete(ctx context.Context, req *schema.SystemApiDeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return l.SystemApiRepo.Delete(ctx, req.Id)
}

func (l *SystemApi) All(ctx context.Context) ([]*model.SystemApi, error) {
	return l.SystemApiRepo.All(ctx)
}
