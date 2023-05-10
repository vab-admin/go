package validate

import "regexp"

var (
	// MobileRegex 是否为有效的手机
	MobileRegex = regexp.MustCompile(`^1[3-9]\d{9}$`)

	// IdCardRegex 有效的身份证格式
	IdCardRegex = regexp.MustCompile(`(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[\dXx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`)

	// ChsRegex 只能是汉字
	ChsRegex = regexp.MustCompile(`^[\x{4e00}-\x{9fa5}\x{9fa6}-\x{9fef}\x{3400}-\x{4db5}\x{20000}-\x{2ebe0}]+$`)

	BankCardRegex = regexp.MustCompile(`^([1-9])(\d{15}|\d{18})$`)
)
