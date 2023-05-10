package handler

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/httpx"
)

type AdminRuleAction struct {
	AdminRuleActionService *service.AdminRuleAction
}

// Query
// @date: 2023/05/10 02:03:14
func (h *AdminRuleAction) Query(c echo.Context) error {
	req := &schema.AdminRuleActionQueryRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminRuleActionService.Query(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Create
// @date: 2023/05/10 02:03:14
func (h *AdminRuleAction) Create(c echo.Context) error {
	req := &schema.AdminRuleActionCreateRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminRuleActionService.Create(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "创建成功")
}

// Edit
// @date: 2023/05/10 02:03:14
func (h *AdminRuleAction) Edit(c echo.Context) error {
	req := &schema.AdminRuleActionEditRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminRuleActionService.Edit(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Update
// @date: 2023/05/10 02:03:14
func (h *AdminRuleAction) Update(c echo.Context) error {
	req := &schema.AdminRuleActionUpdateRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminRuleActionService.Update(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "更新成功")
}

// Delete
// @date: 2023/05/10 02:03:14
func (h *AdminRuleAction) Delete(c echo.Context) error {
	req := &schema.AdminRuleActionDeleteRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminRuleActionService.Delete(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "删除成功")
}
