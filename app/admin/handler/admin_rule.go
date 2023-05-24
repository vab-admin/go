package handler

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/httpx"
)

type AdminRule struct {
	AdminRuleService *service.AdminRule
}

// Create
// @date: 2023/05/07 23:03:41
func (h *AdminRule) Create(c echo.Context) error {
	req := &schema.AdminRuleCreateRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}
	ctx := c.Request().Context()
	if err := h.AdminRuleService.Create(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "创建成功")
}

// Edit
// @date: 2023/05/07 23:03:41
func (h *AdminRule) Edit(c echo.Context) error {
	req := &schema.AdminRuleEditRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.AdminRuleService.Edit(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// UpdateField
// @param c
// @date 2023-05-10 01:23:55
func (h *AdminRule) UpdateField(c echo.Context) error {
	req := &schema.AdminRuleUpdateFieldRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminRuleService.UpdateField(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "更新成功")
}

// Update
// @date: 2023/05/07 23:03:41
func (h *AdminRule) Update(c echo.Context) error {
	req := &schema.AdminRuleUpdateRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminRuleService.Update(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "更新成功")
}

// Delete
// @date: 2023/05/07 23:03:41
func (h *AdminRule) Delete(c echo.Context) error {
	req := &schema.AdminRuleDeleteRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.AdminRuleService.Delete(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "删除成功")
}

// Query
// @param c
// @date 2023-05-11 00:41:09
func (h *AdminRule) Query(c echo.Context) error {
	ctx := c.Request().Context()

	result, err := h.AdminRuleService.Tree(ctx)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", result)
}
