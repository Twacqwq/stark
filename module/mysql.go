package module

import (
	"context"

	"github.com/FarmerChillax/stark"
	"gorm.io/gorm"
)

// var mysqlOnce sync.Once

type mysqlConn struct {
	client *gorm.DB
}

func (mc *mysqlConn) Get(ctx context.Context) *gorm.DB {
	return mc.client
}

func RegisterMySQL(db *gorm.DB) {
	stark.Mysql = &mysqlConn{client: db}
}
