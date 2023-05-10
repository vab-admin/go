package model

type Supplier struct {
	ID *uint64 `gorm:"size:11;primaryKey;comment:主键;column:id" json:"id,omitempty"`
}
