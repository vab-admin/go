package schema

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"vab-admin/go/pkg/model"
	"vab-admin/go/pkg/pagination"
	"vab-admin/go/pkg/validate"
)

type (
	// AdminUserLoginRequest 管理员用户登录请求
	AdminUserLoginRequest struct {
		//  用户账号
		Username string `json:"username"`
		// 用户密码
		Password string `json:"password"`
	}

	// AdminUserLoginResponse 管理员登录响应
	AdminUserLoginResponse struct {
		// 登录token
		Token string `json:"token"`
	}

	// AdminUserInfo 管理员信息
	AdminUserInfo struct {
		// 管理员头像
		Avatar string `json:"avatar"`
		// 管理员用户名
		Username string `json:"username"`
		// 管理员角色
		Roles []string `json:"roles"`
		// 管理员按钮权限
		Permissions []string `json:"permissions"`
	}

	// AdminUserQueryRequest 查询管理员列表
	AdminUserQueryRequest struct {
		pagination.Param
		Timestamp
		// 用户名
		Username string `query:"username"`
		// 昵称
		Nickname string `query:"nickname"`
		// 手机号码
		Mobile string `query:"mobile"`
	}

	// AdminUserCreateRequest 创建管理员请求
	AdminUserCreateRequest struct {
		Username string   `json:"username"`
		Password string   `json:"password"`
		Nickname string   `json:"nickname"`
		Mobile   string   `json:"mobile"`
		Roles    []uint64 `json:"roles"`
	}

	// AdminUserEditRequest 编辑管理员请求
	AdminUserEditRequest struct {
		ID uint64 `param:"id"`
	}

	// AdminUserUpdateRequest 更新管理员请求
	AdminUserUpdateRequest struct {
		AdminUserEditRequest
		AdminUserCreateRequest
	}

	AdminUserDeleteRequest struct {
		ID uint64 `param:"id"`
	}

	AdminRouterMeta struct {
		Title       string `json:"title,omitempty"`
		Icon        string `json:"icon,omitempty"`
		Hidden      bool   `json:"hidden,omitempty"`
		LevelHidden bool   `json:"levelHidden,omitempty"`
	}
	AdminRouter struct {
		Id        uint64           `json:"-"`
		ParentId  uint64           `json:"-"`
		Path      string           `json:"path"`
		Name      string           `json:"name"`
		Component string           `json:"component"`
		Redirect  string           `json:"redirect,omitempty"`
		Meta      *AdminRouterMeta `json:"meta,omitempty"`
		Children  []*AdminRouter   `json:"children,omitempty"`
	}
)

var (
	// 用户名验证
	ruleAdminUserUsernameRequired = validation.Required.Error("用户名不得为空")
	ruleAdminUserUsernameLength   = validation.RuneLength(2, 20).Error("用户名长度在2-20位")
	ruleAdminUserUsernameUnique   = validate.Unique(&model.AdminUser{}, "account").Error("用户名已存在")

	// 昵称验证
	ruleAdminUserNicknameLength = validation.RuneLength(0, 20).Error("昵称长度不得大于20位字符")

	// 登录密码验证
	ruleAdminUserPasswordRequired = validation.Required.Error("登录密码不得为空")
	ruleAdminUserPasswordLength   = validation.RuneLength(6, 20).Error("登录密码在6-20位")

	// 手机号码验证
	ruleAdminUserMobileRequired = validation.Required.Error("手机号码不得为空")
	ruleAdminUserMobilePattern  = validate.IsMobile.Error("手机号码格式有误")
	ruleAdminUserMobileUnique   = validate.Unique(&model.AdminUser{}, "mobile").Error("手机号码已存在")

	// 登录账号验证
	ruleAdminUserLoginAccountRequired = ruleAdminUserUsernameRequired.Error("登录账号不得为空")
	ruleAdminUserLoginAccountLength   = ruleAdminUserUsernameLength.Error("登录账号长度在2-20位")

	// 用户id验证
	ruleAdminUserIdRequired = validation.Required.Error("用户id不得为空")
)

// Validate
// @date 2023-05-09 23:46:03
func (v *AdminUserLoginRequest) Validate() error {
	return validate.Check(
		validate.Field(&v.Username, ruleAdminUserLoginAccountRequired, ruleAdminUserLoginAccountLength),
		validate.Field(&v.Password, ruleAdminUserPasswordRequired, ruleAdminUserPasswordLength),
	)
}

// Validate
// @date 2023-05-09 23:46:02
func (v *AdminUserUpdateRequest) Validate() error {
	not := validate.UniqueNot("id", v.ID)

	return validate.Check(
		validate.Field(&v.ID, ruleAdminUserIdRequired),
		validate.Field(&v.Username, ruleAdminUserUsernameRequired, ruleAdminUserUsernameLength, ruleAdminUserUsernameUnique.Where(not)),
		validate.Field(&v.Nickname, ruleAdminUserNicknameLength),
		validate.Field(&v.Password, ruleAdminUserPasswordLength),
		validate.Field(&v.Mobile, ruleAdminUserMobileRequired, ruleAdminUserMobilePattern, ruleAdminUserMobileUnique.Where(not)),
	)
}

// Validate
// @date 2023-05-09 23:46:01
func (v *AdminUserCreateRequest) Validate() error {
	return validate.Check(
		validate.Field(&v.Username, ruleAdminUserUsernameRequired, ruleAdminUserUsernameLength, ruleAdminUserUsernameUnique),
		validate.Field(&v.Nickname, ruleAdminUserNicknameLength),
		validate.Field(&v.Password, ruleAdminUserPasswordRequired, ruleAdminUserPasswordLength),
		validate.Field(&v.Mobile, ruleAdminUserMobileRequired, ruleAdminUserMobilePattern, ruleAdminUserMobileUnique),
	)
}

// Validate
// @date 2023-05-09 23:46:00
func (v *AdminUserEditRequest) Validate() error {
	return ruleAdminUserIdRequired.Validate(&v.ID)
}
