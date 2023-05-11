package service

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/pkg/model"
)

type AdminGroupRule struct {
	AdminRuleRepo      *repository.AdminRule
	AdminGroupRuleRepo *repository.AdminGroupRule
}

// Create
// @param ctx
// @param groupId
// @param ruleId
// @date 2023-05-10 21:36:12
func (s *AdminGroupRule) Create(ctx context.Context, groupId uint64, ruleId ...uint64) (err error) {
	ruleId = pie.Unique(ruleId)

	ruleId, err = s.AdminRuleRepo.IdsByIds(ctx, ruleId...)
	if err != nil {
		return err
	}

	if len(ruleId) <= 0 {
		return nil
	}

	rules := pie.Map(ruleId, func(ruleId uint64) *model.AdminGroupRule {
		return &model.AdminGroupRule{GroupID: groupId, RuleID: ruleId}
	})

	return s.AdminGroupRuleRepo.Create(ctx, rules...)
}

// Update
// @param ctx
// @param groupId
// @param ruleId
// @date 2023-05-10 21:46:20
func (s *AdminGroupRule) Update(ctx context.Context, groupId uint64, ruleId ...uint64) error {
	ruleId = pie.Unique(ruleId)
	// 获取到此分组已有到权限
	existRuleId, err := s.AdminGroupRuleRepo.RuleIdByGroupId(ctx, groupId)
	if err != nil {
		return err
	}

	insertRuleId, deleteRuleId := pie.Diff(existRuleId, ruleId)

	if len(deleteRuleId) > 0 {
		if err = s.AdminGroupRuleRepo.DeleteByGroupIdWithRuleId(ctx, groupId, deleteRuleId...); err != nil {
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

// DeleteByGroupId
// @param ctx
// @param groupId
// @date 2023-05-11 21:16:09
func (s *AdminGroupRule) DeleteByGroupId(ctx context.Context, groupId uint64) error {
	return s.AdminGroupRuleRepo.DeleteByGroupId(ctx, groupId)
}
