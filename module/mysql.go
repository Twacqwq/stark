package module

import (
	"context"

	"github.com/FarmerChillax/stark"
	"github.com/FarmerChillax/stark/pkg/utils"
	"gorm.io/gorm"
)

type mysqlConn struct {
	client *gorm.DB
}

func (mc *mysqlConn) Get(ctx context.Context) *gorm.DB {
	return mc.client
}

func RegisterMysql(app *stark.Application) error {
	mysqlConf := app.Config.Mysql
	db, err := utils.NewMysql(mysqlConf)
	if err != nil {
		return err
	}
	stark.Mysql = &mysqlConn{client: db}
	return nil
}
