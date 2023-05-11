package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

type AdminGroupRule struct{}

// Create
// @param ctx
// @param groupRule
// @date 2023-05-11 00:40:33
func (*AdminGroupRule) Create(ctx context.Context, groupRule ...*model.AdminGroupRule) error {
	tx := db.Instance(ctx).Create(groupRule)
	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// DeleteByGroupIdWithRuleId
// @param ctx
// @param groupId
// @param ruleId
// @date 2023-05-11 00:40:32
func (r *AdminGroupRule) DeleteByGroupIdWithRuleId(ctx context.Context, groupId uint64, ruleId ...uint64) error {
	tx := db.Instance(ctx).Where("group_id = @groupId AND rule_id IN @ruleId",
		sql.Named("groupId", groupId),
		sql.Named("ruleId", ruleId),
	).Delete(&model.AdminGroupRule{})

	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// RuleIdByGroupId
// @param ctx
// @param groupId
// @date 2023-05-11 00:40:31
func (r *AdminGroupRule) RuleIdByGroupId(ctx context.Context, groupId uint64) ([]uint64, error) {
	var ruleId []uint64
	tx := db.Instance(ctx).Model(&model.AdminGroupRule{}).Where("group_id = @groupId", sql.Named("groupId", groupId)).Pluck("rule_id", &ruleId)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据规则groupId获取有的RuleId失败")
		return nil, errors.ErrInternalServer
	}
	return ruleId, nil
}

func (r *AdminGroupRule) DeleteByGroupId(ctx context.Context, groupId uint64) error {
	return db.Session(ctx).Where("group_id = @groupId", sql.Named("groupId", groupId)).Delete(&model.AdminGroupRule{}).Error
}
