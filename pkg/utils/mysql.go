package utils

import (
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/FarmerChillax/stark/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func validateMysqlConfig(mysqlSetting *config.MysqlConfig) error {
	if mysqlSetting == nil {
		return fmt.Errorf("mysqlSetting is nil")
	}
	if mysqlSetting.Username == "" {
		return fmt.Errorf("lack of mysqlSetting.UserName")
	}
	if mysqlSetting.Password == "" {
		return fmt.Errorf("lack of mysqlSetting.Password")
	}
	if mysqlSetting.Host == "" {
		return fmt.Errorf("lack of mysqlSetting.Host")
	}
	if mysqlSetting.DBName == "" {
		return fmt.Errorf("lack of mysqlSetting.DBName")
	}
	if mysqlSetting.Charset == "" {
		return fmt.Errorf("lack of mysqlSetting.Charset")
	}
	if mysqlSetting.Loc == "" {
		mysqlSetting.Loc = "Local"
	} else {
		mysqlSetting.Loc = url.QueryEscape(mysqlSetting.Loc)
	}
	return nil
}

func NewMysql(mysqlConf *config.MysqlConfig) (*gorm.DB, error) {
	dsnTmplate := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"

	dsn := fmt.Sprintf(dsnTmplate, mysqlConf.Username, mysqlConf.Password,
		mysqlConf.Host, mysqlConf.Port, mysqlConf.DBName)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}

	sqldb, err := db.DB()
	if err != nil {
		// logrus.Errorf("NewMysql.db.DB err: %v", err)
		return nil, err
	}

	maxIdle := 10
	maxOpen := 30
	sqldb.SetMaxIdleConns(maxIdle)
	sqldb.SetMaxOpenConns(maxOpen)

	return db, nil
}

func NewMySQLConn(mysqlSetting *config.MysqlConfig) (*gorm.DB, error) {
	if mysqlSetting.Dsn == "" {
		if err := validateMysqlConfig(mysqlSetting); err != nil {
			return nil, err
		}
		mysqlSetting.Dsn = fmt.Sprintf("%s:%s@(%s)/%s?charset=%s&parseTime=%t&loc=%s&&timeout=%ds",
			mysqlSetting.Username,
			mysqlSetting.Password,
			mysqlSetting.Host,
			mysqlSetting.DBName,
			mysqlSetting.Charset,
			mysqlSetting.ParseTime,
			mysqlSetting.Loc,
			mysqlSetting.Timeout,
		)
	}

	dsn := mysqlSetting.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if os.Getenv("ENV") == "dev" {
		db.Logger.LogMode(logger.Info)
	} else {
		db.Logger.LogMode(logger.Error)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	maxIdle := 10
	maxOpen := 30
	if mysqlSetting.MaxOpen > 0 && mysqlSetting.MaxIdle > 0 {
		maxIdle = mysqlSetting.MaxIdle
		maxOpen = mysqlSetting.MaxOpen
	}
	if mysqlSetting.ConnMaxLifeSecond > 0 {
		sqlDB.SetConnMaxLifetime(time.Duration(mysqlSetting.ConnMaxLifeSecond) * time.Second)
	}
	sqlDB.SetMaxIdleConns(maxIdle)
	sqlDB.SetMaxOpenConns(maxOpen)

	return db, nil
}
