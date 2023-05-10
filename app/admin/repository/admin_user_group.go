package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

type AdminUserGroup struct{}

// Create
// @param ctx
// @param userGroup
// @date 2023-05-11 00:03:28
func (*AdminUserGroup) Create(ctx context.Context, userGroup ...*model.AdminUserGroup) error {
	tx := db.Instance(ctx).Create(userGroup)
	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// GroupIdByUserId
// @param ctx
// @param userId
// @date 2023-05-11 00:06:08
func (r *AdminUserGroup) GroupIdByUserId(ctx context.Context, userId uint64) ([]uint64, error) {
	var groupId []uint64
	tx := db.Instance(ctx).Model(&model.AdminUserGroup{}).Where("user_id = @userId", sql.Named("userId", userId)).Pluck("group_id", &groupId)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据规则groupId获取有的RuleId失败")
		return nil, errors.ErrInternalServer
	}
	return groupId, nil
}

// DeleteByUserIdWithGroupId
// @param ctx
// @param userId
// @param groupId
// @date 2023-05-11 00:12:35
func (r *AdminUserGroup) DeleteByUserIdWithGroupId(ctx context.Context, userId uint64, groupId ...uint64) error {
	tx := db.Instance(ctx).Where("user_id = @userId AND group_id IN @groupId",
		sql.Named("userId", userId),
		sql.Named("groupId", groupId),
	).Delete(&model.AdminUserGroup{})

	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}
