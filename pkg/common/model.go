package common

import "gorm.io/plugin/soft_delete"

type ColumnCreateModifyDeleteTime struct {
	ID          int64                 `gorm:"primary_key;AUTO_INCREMENT" json:"id" db:"id"`
	CreateTime  int64                 `json:"create_time" db:"create_time" gorm:"autoCreateTime"`
	ModifyTime  int64                 `json:"modify_time" db:"modify_time" gorm:"autoUpdateTime"`
	DeletedTime int64                 `json:"deleted_time" db:"deleted_time"`
	IsDel       soft_delete.DeletedAt `json:"is_del" db:"is_del" gorm:"softDelete:flag,DeletedAtField:DeletedTime"`
}
