package handler

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/httpx"
)

type AdminGroup struct {
	AdminGroupService *service.AdminGroup
}

// Query
// @date: 2023/05/10 20:34:51
func (h *AdminGroup) Query(c echo.Context) error {
	req := &schema.AdminGroupQueryRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminGroupService.Query(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Create
// @date: 2023/05/10 20:34:51
func (h *AdminGroup) Create(c echo.Context) error {
	req := &schema.AdminGroupCreateRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminGroupService.Create(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "创建成功")
}

// Edit
// @date: 2023/05/10 20:34:51
func (h *AdminGroup) Edit(c echo.Context) error {
	req := &schema.AdminGroupEditRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminGroupService.Edit(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Update
// @date: 2023/05/10 20:34:51
func (h *AdminGroup) Update(c echo.Context) error {
	req := &schema.AdminGroupUpdateRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminGroupService.Update(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "更新成功")
}

// Delete
// @date: 2023/05/10 20:34:51
func (h *AdminGroup) Delete(c echo.Context) error {
	req := &schema.AdminGroupDeleteRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminGroupService.Delete(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "删除成功")
}
