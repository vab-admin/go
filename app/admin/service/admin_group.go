package service

import (
	"context"
	"erp/app/admin/repository"
	"erp/app/admin/schema"
	"erp/pkg/db"
	"erp/pkg/model"
	"github.com/alibabacloud-go/tea/tea"
	"time"
)

type AdminGroup struct {
	AdminGroupRepo        *repository.AdminGroup
	AdminGroupRuleService *AdminGroupRule
}

// Query
// @param ctx
// @param req
// @date 2023-05-10 20:41:32
func (l *AdminGroup) Query(ctx context.Context, req *schema.AdminGroupQueryRequest) ([]*model.AdminGroup, error) {

	return l.AdminGroupRepo.Query(ctx, req)
}

// Create
// @param ctx
// @param req
// @date 2023-05-10 20:41:31
func (l *AdminGroup) Create(ctx context.Context, req *schema.AdminGroupCreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	now := tea.Int64(time.Now().Unix())

	m := &model.AdminGroup{
		Name:      tea.String(req.Name),
		CreatedAt: now,
		UpdatedAt: now,
	}

	return db.Transaction(ctx, func(ctx context.Context) error {

		if err := l.AdminGroupRepo.Create(ctx, m); err != nil {
			return err
		}

		return l.AdminGroupRuleService.Create(ctx, m.GetId(), req.Rules...)
	})
}

// Edit
// @param ctx
// @param req
// @date 2023-05-10 20:41:31
func (l *AdminGroup) Edit(ctx context.Context, req *schema.AdminGroupEditRequest) (*model.AdminGroup, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return l.AdminGroupRepo.Edit(ctx, req.Id)
}

// Update
// @param ctx
// @param req
// @date 2023-05-10 20:41:30
func (l *AdminGroup) Update(ctx context.Context, req *schema.AdminGroupUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	now := tea.Int64(time.Now().Unix())

	m := &model.AdminGroup{
		Name:      tea.String(req.Name),
		UpdatedAt: now,
	}

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := l.AdminGroupRepo.Update(ctx, req.Id, m); err != nil {
			return err
		}

		return l.AdminGroupRuleService.Update(ctx, req.Id, req.Rules...)
	})
}

// Delete
// @param ctx
// @param req
// @date 2023-05-10 20:41:29
func (l *AdminGroup) Delete(ctx context.Context, req *schema.AdminGroupDeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}
	return l.AdminGroupRepo.Delete(ctx, req.Id)
}
