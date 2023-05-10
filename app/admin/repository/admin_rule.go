package repository

import (
	"context"
	"database/sql"
	"erp/pkg/db"
	"erp/pkg/errors"
	"erp/pkg/model"
	log "github.com/sirupsen/logrus"
)

type AdminRule struct {
}

var (
	ErrAdminRuleNotFound = errors.New("不存在的菜单")
)

// ById
// @date: 2023/05/07 23:05:42
func (r *AdminRule) ById(ctx context.Context, id uint64) (*model.AdminRule, error) {
	row := &model.AdminRule{}

	tx := db.Session(ctx).Model(&model.AdminRule{}).Where("id = @id", sql.Named("id", id)).Limit(1).Find(row)

	if err := tx.Error; err != nil {
		log.WithError(err).Error("获取编辑的菜单失败")
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, ErrAdminRuleNotFound
	}

	return row, nil
}

// Delete
// @date: 2023/05/07 23:05:42
func (r *AdminRule) Delete(ctx context.Context, id uint64) error {
	err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Delete(&model.AdminRule{}).Error
	if err != nil {
		log.WithError(err).Error("删除菜单失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Update
// @date: 2023/05/07 23:05:42
func (r *AdminRule) Update(ctx context.Context, id uint64, m any) error {
	if err := db.Session(ctx).Model(&model.AdminRule{}).Where("id = @id", sql.Named("id", id)).Updates(m).Error; err != nil {
		log.WithError(err).WithField("id", id).Error("更新菜单失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Create
// @date: 2023/05/07 23:05:42
func (r *AdminRule) Create(ctx context.Context, m *model.AdminRule) error {
	if err := db.Session(ctx).Create(m).Error; err != nil {
		log.WithError(err).Error("创建菜单失败")
		return errors.ErrInternalServer
	}
	return nil
}

// All
// @param ctx
// @date 2023-05-08 00:28:54
func (r *AdminRule) All(ctx context.Context) ([]*model.AdminRule, error) {
	var rows []*model.AdminRule
	tx := db.Instance(ctx).Order("id DESC").Find(&rows)
	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}
	return rows, nil
}

// IdsByIds
// @param ctx
// @param ids
// @date 2023-05-11 00:03:00
func (r *AdminRule) IdsByIds(ctx context.Context, ids ...uint64) ([]uint64, error) {
	var newIds []uint64
	tx := db.Instance(ctx).Model(&model.AdminRule{}).Where("id IN @ids", sql.Named("ids", ids)).Pluck("id", &newIds)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据规则id获取存在的规则id失败")
		return nil, errors.ErrInternalServer
	}

	return newIds, nil
}
