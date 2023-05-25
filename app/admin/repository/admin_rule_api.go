package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

type AdminRuleApi struct{}

// Create
// @param ctx
// @param groupRule
// @date 2023-05-11 22:37:34
func (*AdminRuleApi) Create(ctx context.Context, groupRule ...*model.AdminRuleApi) error {
	tx := db.Instance(ctx).Create(groupRule)
	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// ApiIdByRuleId
// @param ctx
// @param ruleId
// @date 2023-05-11 22:37:36
func (r *AdminRuleApi) ApiIdByRuleId(ctx context.Context, ruleId uint64) ([]uint64, error) {
	var apiId []uint64
	tx := db.Instance(ctx).Model(&model.AdminRuleApi{}).Where("rule_id = @ruleId", sql.Named("ruleId", ruleId)).Pluck("api_id", &apiId)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("获取规则绑定的接口id失败")
		return nil, errors.ErrInternalServer
	}
	return apiId, nil
}

func (r *AdminRuleApi) DeleteByRuleIdWithApiId(ctx context.Context, ruleId uint64, apiId ...uint64) error {
	tx := db.Instance(ctx).Where("rule_id = @ruleId AND api_id IN @apiId",
		sql.Named("ruleId", ruleId),
		sql.Named("apiId", apiId),
	).Delete(&model.AdminRuleApi{})

	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// DeleteByRuleId
// @param ctx
// @param ruleId
// @date 2023-05-25 23:09:59
func (r *AdminRuleApi) DeleteByRuleId(ctx context.Context, ruleId uint64) error {
	return db.Session(ctx).Where("rule_id = @ruleId", sql.Named("ruleId", ruleId)).Delete(&model.AdminRuleApi{}).Error
}

// DeleteByApiId
// @param ctx
// @param apiId
// @date 2023-05-25 23:09:58
func (r *AdminRuleApi) DeleteByApiId(ctx context.Context, apiId uint64) error {
	return db.Session(ctx).Where("api_id = @apiId", sql.Named("apiId", apiId)).Delete(&model.AdminRuleApi{}).Error
}
