package repository

import (
	"context"
	"database/sql"
	log "github.com/sirupsen/logrus"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
)

type AdminUserRole struct{}

// Create
// @param ctx
// @param userRole
// @date 2023-05-11 00:03:28
func (*AdminUserRole) Create(ctx context.Context, userRole ...*model.AdminUserRole) error {
	tx := db.Instance(ctx).Create(userRole)
	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// RoleIdByUserId
// @param ctx
// @param userId
// @date 2023-05-11 00:06:08
func (r *AdminUserRole) RoleIdByUserId(ctx context.Context, userId uint64) ([]uint64, error) {
	var roleId []uint64
	tx := db.Instance(ctx).Model(&model.AdminUserRole{}).Where("user_id = @userId", sql.Named("userId", userId)).Pluck("role_id", &roleId)
	if err := tx.Error; err != nil {
		log.WithError(err).Error("根据规则roleId获取有的RuleId失败")
		return nil, errors.ErrInternalServer
	}
	return roleId, nil
}

// DeleteByUserIdWithRoleId
// @param ctx
// @param userId
// @param roleId
// @date 2023-05-11 00:12:35
func (r *AdminUserRole) DeleteByUserIdWithRoleId(ctx context.Context, userId uint64, roleId ...uint64) error {
	tx := db.Instance(ctx).Where("user_id = @userId AND role_id IN @roleId",
		sql.Named("userId", userId),
		sql.Named("roleId", roleId),
	).Delete(&model.AdminUserRole{})

	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}

	return nil
}

// DeleteByUserId
// @param ctx
// @param userId
// @date 2023-05-11 21:12:23
func (r *AdminUserRole) DeleteByUserId(ctx context.Context, userId uint64) error {
	return db.Session(ctx).Where("user_id IN @userId", sql.Named("userId", userId)).Delete(&model.AdminUserRole{}).Error
}
