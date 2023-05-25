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
	"vab-admin/go/pkg/pagination"
)

type AdminUser struct{}

// Login
// @param ctx
// @param account
// @date 2023-05-06 17:45:07
func (*AdminUser) Login(ctx context.Context, account string) (*model.AdminUser, error) {
	row := &model.AdminUser{}

	tx := db.Instance(ctx).
		Model(&model.AdminUser{}).
		Where("account = @account", sql.Named("account", account)).
		Order("id DESC").
		Limit(1).
		Select([]string{"id", "account", "password"}).Find(row)

	if err := tx.Error; err != nil {
		log.WithError(err).WithField("account", account).Error("管理员登录失败")
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected < 1 {
		return nil, errors.ErrAdminUserAccountNotFound
	}

	return row, nil
}

// Create
// @param ctx
// @param user
// @date 2023-05-07 21:31:05
func (u *AdminUser) Create(ctx context.Context, user *model.AdminUser) error {
	tx := db.Session(ctx).Create(user)

	if err := tx.Error; err != nil {
		log.WithError(err).WithField("user", user).Error("创建管理员用户失败")
		return errors.ErrInternalServer
	}

	return nil
}

// Query
// @param ctx
// @param req
// @date 2023-05-07 21:51:31
func (u *AdminUser) Query(ctx context.Context, req *schema.AdminUserQueryRequest) (*pagination.Paginator[[]*model.AdminUser], error) {
	tx := db.Instance(ctx).Model(&model.AdminUser{}).Order("created_at DESC").Select([]string{"id", "account", "nickname", "mobile", "created_at"})

	if v := req.Mobile; v != "" {
		tx.Where("mobile = @mobile", sql.Named("mobile", strings.TrimSpace(v)))
	}

	if v := req.Username; v != "" {
		tx.Where("username = @username", sql.Named("username", strings.TrimSpace(v)))
	}

	if v := req.Nickname; v != "" {
		tx.Where("nickname = @nickname", sql.Named("nickname", strings.TrimSpace(v)))
	}

	if !req.IsZero() {
		tx.Where("created_at BETWEEN @start AND @end",
			sql.Named("start", req.Start),
			sql.Named("end", req.End),
		)
	}

	return pagination.Paging[[]*model.AdminUser](tx, &req.Param)
}

// Edit
// @param ctx
// @param id
// @date 2023-05-07 22:26:54
func (*AdminUser) Edit(ctx context.Context, id uint64) (*model.AdminUser, error) {
	row := &model.AdminUser{}
	tx := db.Instance(ctx).
		Model(&model.AdminUser{}).
		Select([]string{"account", "nickname", "mobile", "id"}).
		Preload("Roles", func(tx *gorm.DB) *gorm.DB {
			return tx.Select([]string{"id", "name"})
		}).
		Where("id = @id", sql.Named("id", id)).
		Order("id DESC").Limit(1).Find(row)

	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, errors.ErrAdminUserAccountNotFound
	}

	return row, nil
}

// Update
// @param ctx
// @param id
// @param user
// @date 2023-05-07 22:39:08
func (u *AdminUser) Update(ctx context.Context, id uint64, user *model.AdminUser) error {
	tx := db.Session(ctx).Where("id = @id", sql.Named("id", id)).Limit(1).Updates(user)
	if err := tx.Error; err != nil {
		return errors.ErrInternalServer
	}
	return nil
}

// ByDeleteId
// @param ctx
// @param userId
// @date 2023-05-12 00:07:43
func (*AdminUser) ByDeleteId(ctx context.Context, userId uint64) (*model.AdminUser, error) {
	row := &model.AdminUser{}
	tx := db.Instance(ctx).
		Model(&model.AdminUser{}).
		Select([]string{"account", "nickname", "mobile", "id"}).
		Where("id = @id", sql.Named("id", userId)).
		Order("id DESC").Limit(1).Find(row)

	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, errors.ErrAdminUserAccountNotFound
	}

	return row, nil
}

// DeleteById
// @param ctx
// @param id
// @date 2023-05-12 00:07:45
func (u *AdminUser) DeleteById(ctx context.Context, id uint64) error {
	return db.Session(ctx).Where("id = @id", sql.Named("id", id)).Limit(1).Delete(&model.AdminUser{}).Error
}

// Router
// @param ctx
// @param id
// @date 2023-05-17 17:12:32
func (u *AdminUser) Router(ctx context.Context, userId uint64) ([]*model.AdminRule, error) {

	var rows []*model.AdminRule

	tx := db.Session(ctx).Model(&model.AdminRule{}).Where("type = ?", model.AdminRuleTypeMenu).Order("sort ASC")

	if userId > 1 {
		roleSubQuery := db.Session(ctx).
			Model(&model.AdminUserRole{}).
			Where("user_id = @userID", sql.Named("userID", userId)).Select("role_id")

		menuSubQuery := db.Session(ctx).
			Model(&model.AdminRoleRule{}).
			Where("role_id IN (@roleID)", sql.Named("roleID", roleSubQuery)).Select("rule_id")

		tx.Where("id IN (@menuID)", sql.Named("menuID", menuSubQuery))
	}

	tx.Find(&rows)

	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	return rows, nil
}

func (u *AdminUser) Info(ctx context.Context, userId uint64) (*model.AdminUser, error) {
	user := &model.AdminUser{}
	tx := db.Instance(ctx).Model(&model.AdminUser{}).
		Select([]string{"id", "account"}).
		Preload("Roles", func(tx *gorm.DB) *gorm.DB {
			return tx.Select([]string{"id", "code"})
		}).
		Where("id = @userID", sql.Named("userID", userId)).
		Order("id DESC").Limit(1).Find(user)

	if err := tx.Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	if tx.RowsAffected == 0 {
		return nil, errors.New("获取信息失败")
	}

	return user, nil
}

func (u *AdminUser) InfoPermissions(ctx context.Context, userId uint64) ([]string, error) {
	tx := db.Session(ctx).Model(&model.AdminRule{}).Where("code IS NOT NULL AND code <> ''")

	roleSub := db.Session(ctx).Model(&model.AdminUserRole{}).Where("user_id = @userId", sql.Named("userId", userId)).Select("role_id")

	ruleIdSub := db.Session(ctx).Model(&model.AdminRoleRule{}).Where("role_id IN (@roleId)", sql.Named("roleId", roleSub)).Select("rule_id")

	tx.Where("id IN (@ruleIds)", sql.Named("ruleIds", ruleIdSub))

	var code []string

	tx.Pluck("code", &code)

	return code, nil
}
