package service

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/pkg/model"
)

type AdminUserRole struct {
	AdminRoleRepo     *repository.AdminRole
	AdminUserRoleRepo *repository.AdminUserRole
}

// Create
// @param ctx
// @param userId
// @param roleId
// @date 2023-05-11 00:03:34
func (s *AdminUserRole) Create(ctx context.Context, userId uint64, roleId ...uint64) (err error) {
	roleId = pie.Unique(roleId)

	roleId, err = s.AdminRoleRepo.IdsByIds(ctx, roleId...)
	if err != nil {
		return err
	}

	if len(roleId) <= 0 {
		return nil
	}

	roles := pie.Map(roleId, func(id uint64) *model.AdminUserRole {
		return &model.AdminUserRole{UserID: userId, RoleID: id}
	})

	return s.AdminUserRoleRepo.Create(ctx, roles...)
}

// Update
// @param ctx
// @param userId
// @param roleId
// @date 2023-05-11 00:08:24
func (s *AdminUserRole) Update(ctx context.Context, userId uint64, roleId ...uint64) error {
	roleId = pie.Unique(roleId)
	// 获取到此用户已有的角色
	existRoleId, err := s.AdminUserRoleRepo.RoleIdByUserId(ctx, userId)
	if err != nil {
		return err
	}

	insertRoleId, deleteRoleId := pie.Diff(existRoleId, roleId)

	if len(deleteRoleId) > 0 {
		if err = s.AdminUserRoleRepo.DeleteByUserIdWithRoleId(ctx, userId, deleteRoleId...); err != nil {
			return err
		}
	}

	if len(insertRoleId) > 0 {
		if err = s.Create(ctx, userId, insertRoleId...); err != nil {
			return err
		}
	}

	return nil
}

// Delete
// @param ctx
// @param userId
// @date 2023-05-11 21:12:52
func (s *AdminUserRole) Delete(ctx context.Context, userId uint64) error {
	return s.AdminUserRoleRepo.DeleteByUserId(ctx, userId)
}
