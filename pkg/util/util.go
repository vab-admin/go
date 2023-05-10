package util

import (
	"erp/pkg/model"
	"gorm.io/datatypes"
)

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

func DataTypeJSON[T any](data T) *datatypes.JSONType[T] {
	v := datatypes.NewJSONType(data)
	return &v
}
