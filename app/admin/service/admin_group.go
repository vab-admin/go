package service

import (
	"context"
	"github.com/alibabacloud-go/tea/tea"
	"time"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/model"
)

type AdminGroup struct {
	EnforcerService       *Enforcer
	AdminGroupRuleService *AdminGroupRule
	AdminGroupRepo        *repository.AdminGroup
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

		if err := l.AdminGroupRuleService.Create(ctx, m.GetId(), req.Rules...); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
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

		if err := l.AdminGroupRuleService.Update(ctx, req.Id, req.Rules...); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
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

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := l.AdminGroupRepo.Delete(ctx, req.Id); err != nil {
			return err
		}

		if err := l.AdminGroupRuleService.DeleteByGroupId(ctx, req.Id); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
	})
}
