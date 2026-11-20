package orm

import (
	"fmt"
	"net/url"
	"time"

	"go-biu/library/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	Db  *gorm.DB
	err error
)

func InitGorm() *gorm.DB {
	dbConf := config.YamlConf.Db
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=%s", dbConf.Username, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.Database, url.QueryEscape("Asia/Shanghai"))

	log := logger.Silent
	if "debug" == config.YamlConf.Base.Model {
		log = logger.Info
	}

	Db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(log),
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDb, _ := Db.DB()
	sqlDb.SetMaxIdleConns(5)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	return Db
}

func CloseGorm() {
	sql, _ := Db.DB()
	sql.Close()
}

//func Builder () {
//	dsn := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", config.YamlConf.Db.Username, config.YamlConf.Db.Password, config.YamlConf.Db.Host, config.YamlConf.Db.Port, config.YamlConf.Db.Database)
//	Db, err = gorm.Open("mysql", dsn)
//	if err != nil {
//		panic(err)
//	}
//	// 连接池
//	Db.DB().SetMaxIdleConns(10)
//	Db.DB().SetMaxOpenConns(100)
//	// 数据库日志
//	if config.YamlConf.Base.Model == "debug" {
//		Db.LogMode(true)
//	} else {
//		Db.LogMode(false)
//	}
//
//	// 全局禁用表名复数
//	Db.SingularTable(true)
//	//defer Db.Close()
//}
