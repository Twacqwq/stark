package module

import (
	"context"

	"github.com/FarmerChillax/stark/config"
	"github.com/FarmerChillax/stark/pkg/helper"
	"gorm.io/gorm"
)

// var mysqlOnce sync.Once

type databseConn struct {
	client *gorm.DB
}

func (mc *databseConn) Get(ctx context.Context) *gorm.DB {
	return mc.client
}

func wrapGorm(db *gorm.DB, err error) (*databseConn, error) {
	return &databseConn{
		client: db,
	}, err
}

func NewDatabase(conf *config.DatabseConfig) (*databseConn, error) {
	return wrapGorm(helper.NewGormDB(conf))
}
