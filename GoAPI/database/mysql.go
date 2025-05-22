package database

import (
	"fmt"
	c "mind-set/config"
	log "mind-set/internal/utils/logger"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var MysqlDB *gorm.DB

type Writer interface {
	Printf(string, ...interface{})
}

type WriterLog struct{}

func (w WriterLog) Printf(format string, args ...interface{}) {
	if c.Config.Mysql.PrintSql {
		log.Logger.Sugar().Infof(format, args...)
	}
}

func initMysql() {
	logConfig := logger.New(
		WriterLog{},
		logger.Config{
			SlowThreshold:             0, //  SQL
			LogLevel:                  logger.LogLevel(c.Config.Mysql.LogLevel),
			IgnoreRecordNotFoundError: false, // ErrRecordNotFound
			Colorful:                  false,
		},
	)

	configs := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: c.Config.Mysql.TablePrefix,
		},
		Logger: logConfig,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		c.Config.Mysql.Username,
		c.Config.Mysql.Password,
		c.Config.Mysql.Host,
		c.Config.Mysql.Port,
		c.Config.Mysql.Database,
		c.Config.Mysql.Charset,
	)
	var err error
	MysqlDB, err = gorm.Open(mysql.Open(dsn), configs)

	if err != nil {
		panic("Mysql connection failed：" + err.Error())
	}

	sqlDB, _ := MysqlDB.DB()
	sqlDB.SetMaxIdleConns(c.Config.Mysql.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.Config.Mysql.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(c.Config.Mysql.MaxLifetime)
}
