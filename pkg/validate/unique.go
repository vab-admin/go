package validate

import (
	"context"
	log "github.com/sirupsen/logrus"

	"erp/pkg/db"
	"erp/pkg/errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"gorm.io/gorm/clause"
)

func Unique(model interface{}, field string) UniqueRule {
	return UniqueRule{
		field: field,
		model: model,
		err:   validation.NewError("100", field+" already exists"),
	}
}

type UniqueRule struct {
	field string
	err   validation.Error
	model interface{}
	where []clause.Expression
}

type UniqueWhere struct {
	Field string

	Value interface{}
}

func (r UniqueRule) Validate(value interface{}) error {
	ctx := context.Background()

	var count int64
	tx := db.Session(ctx).Model(r.model).Clauses(clause.Eq{Column: r.field, Value: value})
	if r.where != nil {
		tx.Clauses(r.where...)
	}

	tx.Count(&count)

	if err := tx.Error; err != nil {
		log.WithError(err).WithField("model", r.model).Error("查询数据是否唯一失败")
		return errors.ErrInternalServer
	}

	if count > 0 {
		return r.err
	}

	return nil
}

func (r UniqueRule) Error(message string) UniqueRule {
	r.err = r.err.SetMessage(message)
	return r
}

func (r UniqueRule) Where(cond ...*UniqueWhere) UniqueRule {
	for _, not := range cond {
		r.where = append(r.where, clause.Neq{Column: not.Field, Value: not.Value})
	}

	return r
}

func (r UniqueRule) And(cond ...*UniqueWhere) UniqueRule {
	for _, not := range cond {
		r.where = append(r.where, clause.Eq{Column: not.Field, Value: not.Value})
	}

	return r
}

func UniqueNot(field string, value interface{}) *UniqueWhere {
	return &UniqueWhere{
		Field: field,
		Value: value,
	}
}

func NewUniqueWhere(f string, v any) *UniqueWhere {
	return &UniqueWhere{
		Field: f,
		Value: v,
	}
}
