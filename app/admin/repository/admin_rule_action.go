package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

var (
	ErrAdminRuleActionNotFound = errors.New("不存在的规则操作")
)

type AdminRuleAction struct{}

// Query
// @date: 2023/05/10 02:16:31
func (r *AdminRuleAction) Query(ctx context.Context, req *schema.AdminRuleActionQueryRequest) ([]*model.AdminRuleAction, error) {
	var rows []*model.AdminRuleAction
	tx := db.Session(ctx).Model(&model.AdminRuleAction{}).Where("rule_id = @ruleId", sql.Named("ruleId", req.RuleId)).Order("created_at DESC").Find(&rows)

	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	return rows, nil
}

// Edit
// @date: 2023/05/10 02:16:31
func (r *AdminRuleAction) Edit(ctx context.Context, id uint64) (*model.AdminRuleAction, error) {
	row := &model.AdminRuleAction{}

	tx := db.Session(ctx).Model(&model.AdminRuleAction{}).Where("id = @id", sql.Named("id", id)).Limit(1).Find(row)

	if err := tx.Error; err != nil {
		log.WithError(err).Error("获取编辑的规则操作失败")
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, ErrAdminRuleActionNotFound
	}

	return row, nil
}

// Delete
// @date: 2023/05/10 02:16:31
func (r *AdminRuleAction) Delete(ctx context.Context, id uint64) error {
	err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Delete(&model.AdminRuleAction{}).Error
	if err != nil {
		log.WithError(err).Error("删除规则操作失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Update
// @date: 2023/05/10 02:16:31
func (r *AdminRuleAction) Update(ctx context.Context, id uint64, m *model.AdminRuleAction) error {
	if err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Updates(m).Error; err != nil {
		log.WithError(err).WithField("id", id).Error("更新规则操作失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Create
// @date: 2023/05/10 02:16:31
func (r *AdminRuleAction) Create(ctx context.Context, m *model.AdminRuleAction) error {
	if err := db.Session(ctx).Create(m).Error; err != nil {
		log.WithError(err).Error("创建规则操作失败")
		return errors.ErrInternalServer
	}
	return nil
}
