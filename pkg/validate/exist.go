package validate

import (
	"context"
	log "github.com/sirupsen/logrus"

	"erp/pkg/db"
	"erp/pkg/errors"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"gorm.io/gorm/clause"
)

type ExistRule struct {
	field string
	err   validation.Error
	model interface{}
	where []clause.Expression
}

func Exist(model interface{}, field string) ExistRule {
	return ExistRule{
		field: field,
		model: model,
		err:   validation.NewError("100", "100"),
	}
}

type ExistWhere struct {
	Field string
	Value interface{}
}

func (r ExistRule) Validate(value interface{}) error {
	ctx := context.Background()

	var count int64
	tx := db.Session(ctx).Model(r.model).Clauses(clause.Eq{Column: r.field, Value: value})
	if r.where != nil {
		tx.Clauses(r.where...)
	}

	tx.Count(&count)

	if err := tx.Error; err != nil {
		log.WithError(err).WithFields(log.Fields{"model": r.model}).Error("查询数据是否存在失败")
		return errors.ErrInternalServer
	}

	if count == 0 {
		return r.err
	}

	return nil
}

func (r ExistRule) Error(message string) ExistRule {
	r.err = r.err.SetMessage(message)
	return r
}

func (r ExistRule) Where(cond ...*UniqueWhere) ExistRule {
	for _, not := range cond {
		r.where = append(r.where, clause.Neq{Column: not.Field, Value: not.Value})
	}

	return r
}

func ExistNot(field string, value interface{}) *ExistWhere {
	return &ExistWhere{
		Field: field,
		Value: value,
	}
}
