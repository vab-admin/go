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

type AdminRole struct {
	EnforcerService      *Enforcer
	AdminRoleRuleService *AdminRoleRule
	AdminUserRoleService *AdminUserRole
	AdminRoleRepo        *repository.AdminRole
}

// Query
// @param ctx
// @param req
// @date 2023-05-10 20:41:32
func (l *AdminRole) Query(ctx context.Context, req *schema.AdminRoleQueryRequest) ([]*model.AdminRole, error) {

	return l.AdminRoleRepo.Query(ctx, req)
}

// Create
// @param ctx
// @param req
// @date 2023-05-10 20:41:31
func (l *AdminRole) Create(ctx context.Context, req *schema.AdminRoleCreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	now := tea.Int64(time.Now().Unix())

	m := &model.AdminRole{
		Name:      tea.String(req.Name),
		CreatedAt: now,
		UpdatedAt: now,
	}

	return db.Transaction(ctx, func(ctx context.Context) error {

		if err := l.AdminRoleRepo.Create(ctx, m); err != nil {
			return err
		}

		if err := l.AdminRoleRuleService.Create(ctx, m.GetId(), req.Rules...); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
	})
}

// Edit
// @param ctx
// @param req
// @date 2023-05-10 20:41:31
func (l *AdminRole) Edit(ctx context.Context, req *schema.AdminRoleEditRequest) (*model.AdminRole, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return l.AdminRoleRepo.Edit(ctx, req.Id)
}

// Update
// @param ctx
// @param req
// @date 2023-05-10 20:41:30
func (l *AdminRole) Update(ctx context.Context, req *schema.AdminRoleUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	now := tea.Int64(time.Now().Unix())

	m := &model.AdminRole{
		Name:      tea.String(req.Name),
		UpdatedAt: now,
	}

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := l.AdminRoleRepo.Update(ctx, req.Id, m); err != nil {
			return err
		}

		if err := l.AdminRoleRuleService.Update(ctx, req.Id, req.Rules...); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
	})
}

// Delete
// @param ctx
// @param req
// @date 2023-05-10 20:41:29
func (l *AdminRole) Delete(ctx context.Context, req *schema.AdminRoleDeleteRequest) error {

	if err := req.Validate(); err != nil {
		return err
	}

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := l.AdminRoleRepo.Delete(ctx, req.Id); err != nil {
			return err
		}

		if err := l.AdminRoleRuleService.DeleteByRoleId(ctx, req.Id); err != nil {
			return err
		}

		if err := l.AdminUserRoleService.DeleteByRoleId(ctx, req.Id); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
	})
}
