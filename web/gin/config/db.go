package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var LiveRecordConn *gorm.DB

func initDbConfig() {
	var dbDialector gorm.Dialector
	if val, err := GetLiveRecordDialector(); err != nil {
		fmt.Println("初始化MySQL失败")
	} else {
		dbDialector = val
	}

	gormDb, err := gorm.Open(dbDialector, &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		//Logger:                 redefineLog(dbType), //拦截、接管 gorm v2 自带日志
	})

	if err != nil {
		//gorm 数据库驱动初始化失败
		fmt.Println("数据库驱动初始化失败")
		return
	}

	if rawDb, err := gormDb.DB(); err == nil {
		rawDb.SetConnMaxIdleTime(time.Second * 30)
		rawDb.SetConnMaxLifetime(GlobalConfig.GetDuration("Gormv2.LiveRecord.Write.SetConnMaxLifetime") * time.Second)
		rawDb.SetMaxIdleConns(GlobalConfig.GetInt("Gormv2.LiveRecord.Write.SetMaxIdleConns"))
		rawDb.SetMaxOpenConns(GlobalConfig.GetInt("Gormv2.LiveRecord.Write.SetMaxOpenConns"))
		LiveRecordConn = gormDb
	}

}

func GetLiveRecordDialector() (gorm.Dialector, error) {
	Host := GlobalConfig.GetString("Gormv2.LiveRecord.Write.Host")
	DataBase := GlobalConfig.GetString("Gormv2.LiveRecord.Write.DataBase")
	Port := GlobalConfig.GetInt("Gormv2.LiveRecord.Write.Port")
	User := GlobalConfig.GetString("Gormv2.LiveRecord.Write.User")
	Pass := GlobalConfig.GetString("Gormv2.LiveRecord.Write.Pass")
	Charset := GlobalConfig.GetString("Gormv2.LiveRecord.Write.Charset")
	addressUrl := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", User, Pass, Host, Port, DataBase, Charset)
	return mysql.Open(addressUrl), nil
}
