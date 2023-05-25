package service

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/pkg/model"
)

type AdminRuleApi struct {
	SystemApiRepo    *repository.SystemApi
	AdminRuleApiRepo *repository.AdminRuleApi
}

// Create
// @param ctx
// @param ruleId
// @param apiId
// @date 2023-05-24 21:17:41
func (s *AdminRuleApi) Create(ctx context.Context, ruleId uint64, apiId ...uint64) (err error) {
	apiId = pie.Unique(apiId)

	apiId, err = s.SystemApiRepo.IdsByIds(ctx, apiId...)
	if err != nil {
		return err
	}

	if len(apiId) <= 0 {
		return nil
	}

	ruleApis := pie.Map(apiId, func(apiId uint64) *model.AdminRuleApi {
		return &model.AdminRuleApi{RuleID: ruleId, ApiID: apiId}
	})

	return s.AdminRuleApiRepo.Create(ctx, ruleApis...)
}

// Update
// @param ctx
// @param ruleId
// @param apiId
// @date 2023-05-24 21:17:38
func (s *AdminRuleApi) Update(ctx context.Context, ruleId uint64, apiId ...uint64) error {
	apiId = pie.Unique(apiId)

	existApiId, err := s.AdminRuleApiRepo.ApiIdByRuleId(ctx, ruleId)
	if err != nil {
		return err
	}

	insertApiId, deleteApiId := pie.Diff(existApiId, apiId)

	if len(deleteApiId) > 0 {
		if err = s.AdminRuleApiRepo.DeleteByRuleIdWithApiId(ctx, ruleId, deleteApiId...); err != nil {
			return err
		}
	}

	if len(insertApiId) > 0 {
		if err = s.Create(ctx, ruleId, insertApiId...); err != nil {
			return err
		}
	}

	return nil
}

// Delete
// @param ctx
// @param ruleId
// @date 2023-05-24 21:19:53
func (s *AdminRuleApi) Delete(ctx context.Context, ruleId uint64) error {
	return s.AdminRuleApiRepo.DeleteByRuleId(ctx, ruleId)
}
