package validate

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

var (
	IsMobile = validation.NewStringRuleWithError(MobileRegex.MatchString, validation.NewError("1000", "手机号码有误"))
	IsIdCard = validation.NewStringRuleWithError(IdCardRegex.MatchString, validation.NewError("1000", "身份证号码有误"))
)
