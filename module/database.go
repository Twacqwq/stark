package module

import (
	"context"

	"gorm.io/gorm"
)

// var mysqlOnce sync.Once

type databseConn struct {
	client *gorm.DB
}

func (mc *databseConn) Get(ctx context.Context) *gorm.DB {
	return mc.client
}
