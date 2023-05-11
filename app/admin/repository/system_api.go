package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/pagination"
)

var (
	ErrSystemApiNotFound = errors.New("不存在的接口")
)

type SystemApi struct{}

// Query
// @date: 2023/05/10 02:16:31
func (r *SystemApi) Query(ctx context.Context, req *schema.SystemApiQueryRequest) (*pagination.Paginator[[]*model.SystemApi], error) {

	tx := db.Session(ctx).Model(&model.SystemApi{}).Order("created_at DESC")

	return pagination.Paging[[]*model.SystemApi](tx, &req.Param)
}

// Edit
// @date: 2023/05/10 02:16:31
func (r *SystemApi) Edit(ctx context.Context, id uint64) (*model.SystemApi, error) {
	row := &model.SystemApi{}

	tx := db.Session(ctx).Model(&model.SystemApi{}).Where("id = @id", sql.Named("id", id)).Limit(1).Find(row)

	if err := tx.Error; err != nil {
		log.WithError(err).Error("获取编辑的规则操作失败")
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, ErrSystemApiNotFound
	}

	return row, nil
}

// Delete
// @date: 2023/05/10 02:16:31
func (r *SystemApi) Delete(ctx context.Context, id uint64) error {
	err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Delete(&model.SystemApi{}).Error
	if err != nil {
		log.WithError(err).Error("删除规则操作失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Update
// @date: 2023/05/10 02:16:31
func (r *SystemApi) Update(ctx context.Context, id uint64, m *model.SystemApi) error {
	if err := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Updates(m).Error; err != nil {
		log.WithError(err).WithField("id", id).Error("更新规则操作失败")
		return errors.ErrInternalServer
	}
	return nil
}

// Create
// @date: 2023/05/10 02:16:31
func (r *SystemApi) Create(ctx context.Context, m *model.SystemApi) error {
	if err := db.Session(ctx).Create(m).Error; err != nil {
		log.WithError(err).Error("创建规则操作失败")
		return errors.ErrInternalServer
	}
	return nil
}

// All
// @param ctx
// @date 2023-05-11 22:33:44
func (*SystemApi) All(ctx context.Context) ([]*model.SystemApi, error) {
	var rows []*model.SystemApi
	return rows, db.Instance(ctx).Find(&rows).Error
}

func (r *SystemApi) IdsByIds(ctx context.Context, ids ...uint64) ([]uint64, error) {
	var newIds []uint64
	tx := db.Instance(ctx).Model(&model.SystemApi{}).Where("id IN @ids", sql.Named("ids", ids)).Pluck("id", &newIds)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据接口id获取存在的接口id失败")
		return nil, errors.ErrInternalServer
	}

	return newIds, nil
}
