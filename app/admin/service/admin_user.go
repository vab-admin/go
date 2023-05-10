package service

import (
	"context"
	"github.com/alibabacloud-go/tea/tea"
	"time"
	"vab-admin/go/app/admin/repository"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/auth/jwtauth"
	"vab-admin/go/pkg/db"
	"vab-admin/go/pkg/errors"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/pagination"
	"vab-admin/go/pkg/util/hash"
)

type AdminUser struct {
	AdminUserRepo         *repository.AdminUser
	AdminUserGroupService *AdminUserGroup
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
	if token, err = jwtauth.CreateToken("", jwtauth.UserInfo{UserID: user.GetId()}); err != nil {
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

		return s.AdminUserGroupService.Create(ctx, user.GetId(), req.Groups...)
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

func (s *AdminUser) Update(ctx context.Context, req *schema.AdminUserUpdateRequest) error {
	if err := req.Validate(); err != nil {
		return err
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

		return s.AdminUserGroupService.Update(ctx, req.ID, req.Groups...)
	})
}
