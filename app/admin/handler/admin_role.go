package handler

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/httpx"
)

type AdminRole struct {
	AdminRoleService *service.AdminRole
}

// Query
// @date: 2023/05/10 20:34:51
func (h *AdminRole) Query(c echo.Context) error {
	req := &schema.AdminRoleQueryRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminRoleService.Query(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Create
// @date: 2023/05/10 20:34:51
func (h *AdminRole) Create(c echo.Context) error {
	req := &schema.AdminRoleCreateRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminRoleService.Create(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "创建成功")
}

// Edit
// @date: 2023/05/10 20:34:51
func (h *AdminRole) Edit(c echo.Context) error {
	req := &schema.AdminRoleEditRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminRoleService.Edit(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Update
// @date: 2023/05/10 20:34:51
func (h *AdminRole) Update(c echo.Context) error {
	req := &schema.AdminRoleUpdateRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminRoleService.Update(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "更新成功")
}

// Delete
// @date: 2023/05/10 20:34:51
func (h *AdminRole) Delete(c echo.Context) error {
	req := &schema.AdminRoleDeleteRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	if err := h.AdminRoleService.Delete(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "删除成功")
}
