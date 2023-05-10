package model

type AdminUserGroup struct {
	UserID  uint64 `gorm:"index;default:0;"`
	GroupID uint64 `gorm:"index;default:0;"`
}

func (*AdminUserGroup) TableName() string {
	return "admin_user_groups"
}
