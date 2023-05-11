package handler

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/httpx"
)

type AdminUser struct {
	AdminUserService *service.AdminUser
}

func NewAdminUser(adminUserService *service.AdminUser) *AdminUser {
	return &AdminUser{AdminUserService: adminUserService}
}

// Login
// @param c
// @date 2023-05-06 17:29:06
func (h *AdminUser) Login(c echo.Context) error {
	req := &schema.AdminUserLoginRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	result, err := h.AdminUserService.Login(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "登录成功", result)
}

// Info
// @param c
// @date 2023-05-07 21:28:13
func (h *AdminUser) Info(c echo.Context) error {

	return httpx.OkJSON(c, "获取成功", &schema.AdminUserInfo{})
}

func (h *AdminUser) Query(c echo.Context) error {
	req := &schema.AdminUserQueryRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	result, err := h.AdminUserService.Query(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "查询成功", result)
}

// Create
// @param c
// @date 2023-05-07 21:28:29
func (h *AdminUser) Create(c echo.Context) error {
	req := &schema.AdminUserCreateRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminUserService.Create(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "添加成功")
}

// Edit
// @param c
// @date 2023-05-07 22:26:40
func (h *AdminUser) Edit(c echo.Context) error {
	req := &schema.AdminUserEditRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	result, err := h.AdminUserService.Edit(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", result)
}

// Delete
// @param c
// @date 2023-05-11 00:41:17
func (h *AdminUser) Delete(c echo.Context) error {
	req := &schema.AdminUserDeleteRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminUserService.Delete(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "删除成功")
}

// Update
// @param c
// @date 2023-05-11 00:41:19
func (h *AdminUser) Update(c echo.Context) error {
	req := &schema.AdminUserUpdateRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminUserService.Update(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "添加成功")
}
