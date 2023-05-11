package handler

import (
	"github.com/labstack/echo/v5"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/app/admin/service"
	"vab-admin/go/pkg/httpx"
)

type SystemApi struct {
	SystemApiService *service.SystemApi
}

// Query
// @date: 2023/05/10 02:03:14
func (h *SystemApi) Query(c echo.Context) error {
	req := &schema.SystemApiQueryRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	resp, err := h.SystemApiService.Query(ctx, req)

	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Create
// @date: 2023/05/10 02:03:14
func (h *SystemApi) Create(c echo.Context) error {
	req := &schema.SystemApiCreateRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.SystemApiService.Create(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "创建成功")
}

// Edit
// @date: 2023/05/10 02:03:14
func (h *SystemApi) Edit(c echo.Context) error {
	req := &schema.SystemApiEditRequest{}
	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()
	resp, err := h.SystemApiService.Edit(ctx, req)
	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}

// Update
// @date: 2023/05/10 02:03:14
func (h *SystemApi) Update(c echo.Context) error {
	req := &schema.SystemApiUpdateRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.SystemApiService.Update(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "更新成功")
}

// Delete
// @date: 2023/05/10 02:03:14
func (h *SystemApi) Delete(c echo.Context) error {
	req := &schema.SystemApiDeleteRequest{}

	if err := httpx.Bind(c, req); err != nil {
		return err
	}

	ctx := c.Request().Context()

	if err := h.SystemApiService.Delete(ctx, req); err != nil {
		return err
	}

	return httpx.OK(c, "删除成功")
}

func (h *SystemApi) All(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := h.SystemApiService.All(ctx)

	if err != nil {
		return err
	}

	return httpx.OkJSON(c, "获取成功", resp)
}
