package service

import (
	"context"
	"erp/app/admin/repository"
	"erp/app/admin/schema"
	"erp/pkg/model"
	"github.com/alibabacloud-go/tea/tea"
	"time"
)

type AdminRuleAction struct {
	AdminRuleActionRepo *repository.AdminRuleAction
}

// Query
// @param ctx
// @param req
// @date 2023-05-10 02:23:44
func (l *AdminRuleAction) Query(ctx context.Context, req *schema.AdminRuleActionQueryRequest) ([]*model.AdminRuleAction, error) {

	return l.AdminRuleActionRepo.Query(ctx, req)
}

// Create
// @param ctx
// @param req
// @date 2023-05-10 02:23:42
func (l *AdminRuleAction) Create(ctx context.Context, req *schema.AdminRuleActionCreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	m := req.ToAdminRuleActionModel()
	m.CreatedAt = tea.Int64(time.Now().Unix())
	return l.AdminRuleActionRepo.Create(ctx, m)
}

// Edit
// @param ctx
// @param req
// @date 2023-05-10 02:23:41
func (l *AdminRuleAction) Edit(ctx context.Context, req *schema.AdminRuleActionEditRequest) (*model.AdminRuleAction, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return l.AdminRuleActionRepo.Edit(ctx, req.Id)
}

// Update
// @param ctx
// @param req
// @date 2023-05-10 02:23:39
func (l *AdminRuleAction) Update(ctx context.Context, req *schema.AdminRuleActionUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	m := req.ToAdminRuleActionModel()
	return l.AdminRuleActionRepo.Update(ctx, req.Id, m)
}

// Delete
// @param ctx
// @param req
// @date 2023-05-10 02:23:38
func (l *AdminRuleAction) Delete(ctx context.Context, req *schema.AdminRuleActionDeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return l.AdminRuleActionRepo.Delete(ctx, req.Id)
}
