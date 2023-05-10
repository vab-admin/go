package validate

import (
	"reflect"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type (
	FieldRules struct {
		fieldPtr interface{}
		rules    []validation.Rule
	}
)

func Field(fieldPtr interface{}, rules ...validation.Rule) *FieldRules {
	return &FieldRules{
		fieldPtr: fieldPtr,
		rules:    rules,
	}
}

func Check(fields ...*FieldRules) error {

	for i, fr := range fields {
		fv := reflect.ValueOf(fr.fieldPtr)
		if fv.Kind() != reflect.Ptr {
			return validation.NewInternalError(validation.ErrFieldPointer(i))
		}

		if err := validation.Validate(fv.Elem().Interface(), fr.rules...); err != nil {

			if ie, ok := err.(validation.InternalError); ok && ie.InternalError() != nil {
				return err
			}

			return err
		}
	}

	return nil
}
