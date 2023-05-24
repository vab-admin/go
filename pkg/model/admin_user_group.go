package model

type AdminUserRole struct {
	UserID uint64 `gorm:"index;default:0;"`
	RoleID uint64 `gorm:"index;default:0;"`
}

func (*AdminUserRole) TableName() string {
	return "admin_user_roles"
}
