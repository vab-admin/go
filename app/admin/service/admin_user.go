package service

import (
	"context"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/elliotchance/pie/v2"
	"time"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/auth/jwtauth"
	"vab-admin/go/pkg/contextx"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/pagination"
	"vab-admin/go/pkg/util"
	"vab-admin/go/pkg/util/hash"
)

type AdminUser struct {
	EnforcerService      *Enforcer
	AdminUserRoleService *AdminUserRole
	AdminUserRepo        *repository.AdminUser
}

// Login
// @param ctx
// @param req
// @date 2023-05-06 17:51:12
func (s *AdminUser) Login(ctx context.Context, req *schema.AdminUserLoginRequest) (*schema.AdminUserLoginResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := s.AdminUserRepo.Login(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	if !hash.PasswordVerify(user.GetPassword(), req.Password) {
		return nil, errors.ErrAdminUserPasswordInvalid
	}

	var token string
	if token, err = jwtauth.CreateToken("api", jwtauth.UserInfo{UserID: user.GetId()}); err != nil {
		return nil, err
	}

	resp := &schema.AdminUserLoginResponse{Token: token}

	return resp, nil
}

// Create
// @param ctx
// @param req
// @date 2023-05-07 21:31:35
func (s *AdminUser) Create(ctx context.Context, req *schema.AdminUserCreateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	now := tea.Int64(time.Now().Unix())

	user := &model.AdminUser{
		Mobile:    tea.String(req.Mobile),
		Account:   tea.String(req.Username),
		Password:  tea.String(hash.PasswordHash(req.Password)),
		CreatedAt: now,
		UpdatedAt: now,
	}

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := s.AdminUserRepo.Create(ctx, user); err != nil {
			return err
		}

		if err := s.AdminUserRoleService.Create(ctx, user.GetId(), req.Roles...); err != nil {
			return err
		}

		return s.EnforcerService.LoadPolicy()
	})
}

// Query
// @param ctx
// @param req
// @date 2023-05-07 22:26:48
func (s *AdminUser) Query(ctx context.Context, req *schema.AdminUserQueryRequest) (*pagination.Paginator[[]*model.AdminUser], error) {
	return s.AdminUserRepo.Query(ctx, req)
}

// Edit
// @param ctx
// @param req
// @date 2023-05-07 22:26:46
func (s *AdminUser) Edit(ctx context.Context, req *schema.AdminUserEditRequest) (*model.AdminUser, error) {
	return s.AdminUserRepo.Edit(ctx, req.ID)
}

// Update
// @param ctx
// @param req
// @date 2023-05-11 16:27:18
func (s *AdminUser) Update(ctx context.Context, req *schema.AdminUserUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
	}

	userId := contextx.FromUserID(ctx)
	if req.ID == 1 && userId != 1 {
		return errors.New("无权限操作超级管理员")
	}

	now := tea.Int64(time.Now().Unix())

	user := &model.AdminUser{
		Mobile:    tea.String(req.Mobile),
		Account:   tea.String(req.Username),
		Nickname:  tea.String(req.Nickname),
		UpdatedAt: now,
	}

	if v := req.Password; v != "" {
		user.Password = tea.String(hash.PasswordHash(req.Password))
	}

	return db.Transaction(ctx, func(ctx context.Context) error {
		if err := s.AdminUserRepo.Update(ctx, req.ID, user); err != nil {
			return err
		}

		if err := s.AdminUserRoleService.Update(ctx, req.ID, req.Roles...); err != nil {
			return err
		}

		return s.EnforcerService.LoadPolicy()
	})
}

// Delete
// @param ctx
// @date 2023-05-11 20:28:18
func (s *AdminUser) Delete(ctx context.Context, req *schema.AdminUserDeleteRequest) error {
	user, err := s.AdminUserRepo.ByDeleteId(ctx, req.ID)
	if err != nil {
		return err
	}

	if user.GetId() == 1 {
		return errors.New("超级管理员无法删除")
	}

	return db.Transaction(ctx, func(ctx context.Context) error {

		if err = s.AdminUserRepo.DeleteById(ctx, req.ID); err != nil {
			return errors.ErrInternalServer
		}

		if err = s.AdminUserRoleService.Delete(ctx, req.ID); err != nil {
			return errors.ErrInternalServer
		}

		return s.EnforcerService.LoadPolicy()
	})

}

// Router
// @param ctx
// @date 2023-05-25 23:13:18
func (s *AdminUser) Router(ctx context.Context) ([]*schema.AdminRouter, error) {
	userId := contextx.FromUserID(ctx)

	routers, err := s.AdminUserRepo.Router(ctx, userId)
	if err != nil {
		return nil, err
	}

	var rows = pie.Map(routers, func(rule *model.AdminRule) *schema.AdminRouter {
		return &schema.AdminRouter{
			Id:        rule.GetId(),
			ParentId:  rule.GetParentId(),
			Path:      rule.GetPath(),
			Name:      rule.GetName(),
			Component: rule.GetComponent(),
			Redirect:  rule.GetRedirect(),
			Meta: &schema.AdminRouterMeta{
				Title:       rule.GetTitle(),
				Icon:        rule.GetIcon(),
				Hidden:      rule.GetHidden(),
				LevelHidden: rule.GetLevelHidden(),
			},
		}
	})

	rows = util.AdminRouterToTree(rows, 0)

	root := &schema.AdminRouter{
		Path:      "/",
		Name:      "Layout",
		Component: "Layout",
		Meta:      &schema.AdminRouterMeta{LevelHidden: true},
		Children:  rows,
	}

	return []*schema.AdminRouter{root}, nil
}

// Info
// @param ctx
// @date 2023-05-26 00:36:24
func (s *AdminUser) Info(ctx context.Context) (*schema.AdminUserInfo, error) {
	userId := contextx.FromUserID(ctx)

	user, err := s.AdminUserRepo.Info(ctx, userId)
	if err != nil {
		return nil, err
	}

	info := &schema.AdminUserInfo{
		Avatar:   "",
		Username: user.GetUsername(),
	}

	if userId == 1 {
		info.Roles = []string{"root"}
	} else {
		info.Roles = pie.Map(user.Roles, func(t *model.AdminRole) string { return t.GetCode() })
	}

	if userId > 1 {
		info.Permissions, _ = s.AdminUserRepo.InfoPermissions(ctx, userId)
	}

	return info, nil
}
