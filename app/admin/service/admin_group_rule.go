package service

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/pkg/model"
)

type AdminRoleRule struct {
	AdminRuleRepo     *repository.AdminRule
	AdminRoleRuleRepo *repository.AdminRoleRule
}

// Create
// @param ctx
// @param groupId
// @param ruleId
// @date 2023-05-10 21:36:12
func (s *AdminRoleRule) Create(ctx context.Context, groupId uint64, ruleId ...uint64) (err error) {
	ruleId = pie.Unique(ruleId)

	ruleId, err = s.AdminRuleRepo.IdsByIds(ctx, ruleId...)
	if err != nil {
		return err
	}

	if len(ruleId) <= 0 {
		return nil
	}

	rules := pie.Map(ruleId, func(ruleId uint64) *model.AdminRoleRule {
		return &model.AdminRoleRule{RoleID: groupId, RuleID: ruleId}
	})

	return s.AdminRoleRuleRepo.Create(ctx, rules...)
}

// Update
// @param ctx
// @param groupId
// @param ruleId
// @date 2023-05-10 21:46:20
func (s *AdminRoleRule) Update(ctx context.Context, groupId uint64, ruleId ...uint64) error {
	ruleId = pie.Unique(ruleId)
	// 获取到此分组已有到权限
	existRuleId, err := s.AdminRoleRuleRepo.RuleIdByRoleId(ctx, groupId)
	if err != nil {
		return err
	}

	insertRuleId, deleteRuleId := pie.Diff(existRuleId, ruleId)

	if len(deleteRuleId) > 0 {
		if err = s.AdminRoleRuleRepo.DeleteByRoleIdWithRuleId(ctx, groupId, deleteRuleId...); err != nil {
			return err
		}
	}

	if len(insertRuleId) > 0 {
		if err = s.Create(ctx, groupId, insertRuleId...); err != nil {
			return err
		}
	}

	return nil
}

// DeleteByRoleId
// @param ctx
// @param groupId
// @date 2023-05-11 21:16:09
func (s *AdminRoleRule) DeleteByRoleId(ctx context.Context, groupId uint64) error {
	return s.AdminRoleRuleRepo.DeleteByRoleId(ctx, groupId)
}
