package service

import (
	"context"
	"github.com/elliotchance/pie/v2"
	"github.com/thoas/go-funk"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/util"
)

type AdminRule struct {
	EnforcerService     *Enforcer
	AdminRuleApiService *AdminRuleApi
	AdminRuleRepo       *repository.AdminRule
}

// Create
// @param ctx
// @param req
// @date 2023-05-08 00:31:38
func (l *AdminRule) Create(ctx context.Context, req *schema.AdminRuleCreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	if err := l.checkParent(ctx, req.ParentId); err != nil {
		return err
	}

	m := req.ToAdminRuleModel()

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := l.AdminRuleRepo.Create(ctx, m); err != nil {
			return err
		}

		return l.AdminRuleApiService.Create(ctx, m.GetId(), req.Apis...)
	})
}

// Edit
// @param ctx
// @param req
// @date 2023-05-08 00:31:37
func (l *AdminRule) Edit(ctx context.Context, req *schema.AdminRuleEditRequest) (*model.AdminRule, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	return l.AdminRuleRepo.Edit(ctx, req.Id)
}

// UpdateField
// @param ctx
// @param req
// @date 2023-05-10 01:25:02
func (l *AdminRule) UpdateField(ctx context.Context, req *schema.AdminRuleUpdateFieldRequest) error {
	return l.AdminRuleRepo.Update(ctx, req.Id, map[string]any{
		req.Field: req.Value,
	})
}

// Update
// @param ctx
// @param req
// @date 2023-05-08 00:31:36
func (l *AdminRule) Update(ctx context.Context, req *schema.AdminRuleUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	menu, err := l.AdminRuleRepo.ById(ctx, req.Id)
	if err != nil {
		return err
	}

	if err = l.checkParent(ctx, req.ParentId); err != nil {
		return err
	}

	if req.ParentId == menu.GetId() {
		return errors.New("无法将自己设置为上级")
	}

	var subMenus []*model.AdminRule
	if subMenus, err = l.subMenus(ctx, req.Id); err != nil {
		return err
	}

	subMenuIds := pie.Map(subMenus, func(t *model.AdminRule) uint64 { return t.GetId() })

	if funk.InUInt64s(subMenuIds, req.ParentId) {
		return errors.New("无法将当前菜单的下级设置为自己的上级")
	}

	m := req.ToAdminRuleModel()

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err = l.AdminRuleRepo.Update(ctx, req.Id, m); err != nil {
			return err
		}

		return l.AdminRuleApiService.Update(ctx, req.Id, req.Apis...)
	})
}

// Delete
// @param ctx
// @param req
// @date 2023-05-08 00:31:35
func (l *AdminRule) Delete(ctx context.Context, req *schema.AdminRuleDeleteRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	subMenus, err := l.subMenus(ctx, req.Id)
	if err != nil {
		return err
	}

	if len(subMenus) > 0 {
		return errors.New("此菜单下有子菜单，无法删除")
	}

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err = l.AdminRuleRepo.Delete(ctx, req.Id); err != nil {
			return err
		}

		if err = l.AdminRuleApiService.Delete(ctx, req.Id); err != nil {
			return err
		}

		return l.EnforcerService.LoadPolicy()
	})
}

// Tree
// @param ctx
// @date 2023-05-08 00:34:17
func (l *AdminRule) Tree(ctx context.Context) ([]*model.AdminRule, error) {
	rows, err := l.AdminRuleRepo.All(ctx)
	if err != nil {
		return nil, err
	}

	return util.AdminRuleToTree(rows, 0), nil
}

// subMenus
// @param ctx
// @param id
// @date 2023-05-10 00:34:35
func (l *AdminRule) subMenus(ctx context.Context, id uint64) ([]*model.AdminRule, error) {
	menus, err := l.AdminRuleRepo.All(ctx)
	if err != nil {
		return nil, err
	}

	return util.FindSubRules(menus, id), nil
}

// checkParent
// @param ctx
// @param parentId
// @date 2023-05-10 00:38:06
func (l *AdminRule) checkParent(ctx context.Context, parentId uint64) error {
	if parentId > 0 {
		parentMenu, err := l.AdminRuleRepo.ById(ctx, parentId)
		if err != nil {
			return err
		}

		if parentMenu.GetType() != model.AdminRuleTypeMenu {
			return errors.New("无法给菜单操作添加下级规则")
		}
	}
	return nil
}
