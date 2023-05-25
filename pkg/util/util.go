package util

import (
	"gorm.io/datatypes"
	"vab-admin/go/app/admin/schema"
	"vab-admin/go/pkg/model"
)

func AdminRouterToTree(rows []*schema.AdminRouter, pid uint64) []*schema.AdminRouter {
	var arr []*schema.AdminRouter
	for _, v := range rows {
		if pid == v.ParentId {
			children := AdminRouterToTree(rows, v.Id)
			v.Children = children
			arr = append(arr, v)
		}
	}
	return arr
}

// AdminRuleToTree
// @param rows
// @param pid
// @date 2023-05-25 21:11:15
func AdminRuleToTree(rows []*model.AdminRule, pid uint64) []*model.AdminRule {
	var arr []*model.AdminRule
	for _, v := range rows {
		if pid == v.GetParentId() {
			children := AdminRuleToTree(rows, v.GetId())
			v.Children = children
			arr = append(arr, v)
		}
	}
	return arr
}

// FindSubRules
// @param categories
// @param parentId
// @date 2023-05-25 21:11:14
func FindSubRules(categories []*model.AdminRule, parentId uint64) []*model.AdminRule {
	var result []*model.AdminRule

	// 遍历所有分类
	for _, category := range categories {
		// 如果当前分类的父分类ID等于指定的父分类ID，说明它是当前分类的子分类
		if category.GetParentId() == parentId {
			// 将当前分类添加到结果列表中
			result = append(result, category)

			// 递归处理当前分类的子分类，并将它们添加到结果列表中
			children := FindSubRules(categories, category.GetId())
			result = append(result, children...)
		}
	}

	return result
}

// DataTypeJSON DataTypeJSON[T any]
// @param data
// @date 2023-05-25 21:11:13
func DataTypeJSON[T any](data T) *datatypes.JSONType[T] {
	v := datatypes.NewJSONType(data)
	return &v
}
