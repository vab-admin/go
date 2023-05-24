package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

type AdminRole struct {
}

var (
	ErrAdminRoleNotFound = errors.New("不存在的分组")
)

// Query
// @date: 2023/05/10 20:35:43
func (r *AdminRole) Query(ctx context.Context, req *schema.AdminRoleQueryRequest) ([]*model.AdminRole, error) {
	var rows []*model.AdminRole

	tx := db.Instance(ctx).Model(&model.AdminRole{}).Select([]string{"id", "name", "created_at", "updated_at"}).Order("created_at DESC")

	if v := req.Name; v != "" {
		tx.Where("name = @name", sql.Named("name", strings.TrimSpace(v)))
	}

	tx.Find(&rows)

	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	return rows, nil
}

// Edit
// @date: 2023/05/10 20:35:43
func (r *AdminRole) Edit(ctx context.Context, id uint64) (*model.AdminRole, error) {
	row := &model.AdminRole{}

	tx := db.Instance(ctx).Model(&model.AdminRole{}).
		Preload("Rules", func(tx *gorm.DB) *gorm.DB {
			return tx.Select([]string{"id", "name"})
		}).
		Select([]string{"name", "id"}).Where("id = @id", sql.Named("id", id)).
		Limit(1).Find(row)

	if err := tx.Error; err != nil {
		log.WithError(err).Error("获取编辑的分组失败")
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, ErrAdminRoleNotFound
	}

	return row, nil
}

// Delete
// @date: 2023/05/10 20:35:43
func (r *AdminRole) Delete(ctx context.Context, id uint64) error {
	err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Delete(&model.AdminRole{}).Error
	if err != nil {
		log.WithError(err).Error("删除分组失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Update
// @date: 2023/05/10 20:35:43
func (r *AdminRole) Update(ctx context.Context, id uint64, m *model.AdminRole) error {
	if err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Updates(m).Error; err != nil {
		log.WithError(err).WithField("id", id).Error("更新分组失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Create
// @date: 2023/05/10 20:35:43
func (r *AdminRole) Create(ctx context.Context, m *model.AdminRole) error {
	if err := db.Session(ctx).Create(m).Error; err != nil {
		log.WithError(err).Error("创建分组失败")
		return errors.ErrInternalServer
	}
	return nil
}

// IdsByIds
// @param ctx
// @param ids
// @date 2023-05-11 00:40:25
func (r *AdminRole) IdsByIds(ctx context.Context, ids ...uint64) ([]uint64, error) {
	var newIds []uint64
	tx := db.Instance(ctx).Model(&model.AdminRole{}).Where("id IN @ids", sql.Named("ids", ids)).Pluck("id", &newIds)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据规则id获取存在的角色id失败")
		return nil, errors.ErrInternalServer
	}

	return newIds, nil
}
