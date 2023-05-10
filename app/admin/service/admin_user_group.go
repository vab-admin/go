package service

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/pkg/model"
)

type AdminUserGroup struct {
	AdminUserGroupRepo *repository.AdminUserGroup
	AdminGroupRepo     *repository.AdminGroup
}

// Create
// @param ctx
// @param userId
// @param groupId
// @date 2023-05-11 00:03:34
func (s *AdminUserGroup) Create(ctx context.Context, userId uint64, groupId ...uint64) (err error) {
	groupId = pie.Unique(groupId)

	groupId, err = s.AdminGroupRepo.IdsByIds(ctx, groupId...)
	if err != nil {
		return err
	}

	if len(groupId) <= 0 {
		return nil
	}

	groups := pie.Map(groupId, func(gid uint64) *model.AdminUserGroup {
		return &model.AdminUserGroup{UserID: userId, GroupID: gid}
	})

	return s.AdminUserGroupRepo.Create(ctx, groups...)
}

// Update
// @param ctx
// @param userId
// @param groupId
// @date 2023-05-11 00:08:24
func (s *AdminUserGroup) Update(ctx context.Context, userId uint64, groupId ...uint64) error {
	groupId = pie.Unique(groupId)
	// 获取到此用户已有的角色
	existGroupId, err := s.AdminUserGroupRepo.GroupIdByUserId(ctx, userId)
	if err != nil {
		return err
	}

	insertGroupId, deleteGroupId := pie.Diff(existGroupId, groupId)

	if len(deleteGroupId) > 0 {
		if err = s.AdminUserGroupRepo.DeleteByUserIdWithGroupId(ctx, userId, deleteGroupId...); err != nil {
			return err
		}
	}

	if len(insertGroupId) > 0 {
		if err = s.Create(ctx, userId, insertGroupId...); err != nil {
			return err
		}
	}

	return nil
}
