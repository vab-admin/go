package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

type AdminRoleRule struct{}

// Create
// @param ctx
// @param roleRule
// @date 2023-05-11 00:40:33
func (*AdminRoleRule) Create(ctx context.Context, roleRule ...*model.AdminRoleRule) error {
	tx := db.Instance(ctx).Create(roleRule)
	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// DeleteByRoleIdWithRuleId
// @param ctx
// @param roleId
// @param ruleId
// @date 2023-05-11 00:40:32
func (r *AdminRoleRule) DeleteByRoleIdWithRuleId(ctx context.Context, roleId uint64, ruleId ...uint64) error {
	tx := db.Instance(ctx).Where("role_id = @roleId AND rule_id IN @ruleId",
		sql.Named("roleId", roleId),
		sql.Named("ruleId", ruleId),
	).Delete(&model.AdminRoleRule{})

	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// RuleIdByRoleId
// @param ctx
// @param roleId
// @date 2023-05-11 00:40:31
func (r *AdminRoleRule) RuleIdByRoleId(ctx context.Context, roleId uint64) ([]uint64, error) {
	var ruleId []uint64
	tx := db.Instance(ctx).Model(&model.AdminRoleRule{}).Where("role_id = @roleId", sql.Named("roleId", roleId)).Pluck("rule_id", &ruleId)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据规则roleId获取有的RuleId失败")
		return nil, errors.ErrInternalServer
	}
	return ruleId, nil
}

func (r *AdminRoleRule) DeleteByRoleId(ctx context.Context, roleId uint64) error {
	return db.Session(ctx).Where("role_id = @roleId", sql.Named("roleId", roleId)).Delete(&model.AdminRoleRule{}).Error
}
