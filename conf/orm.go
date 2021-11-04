package conf

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	authModel "ssopa/model/auth"
	reportTemplateModel "ssopa/model/report_template"
	"time"
)

var (
	Orm *gorm.DB
)

func SetupOrm()  {

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GetConfig("mysql::user"),
		GetConfig("mysql::password"),
		GetConfig("mysql::host"),
		GetConfig("mysql::db"))
	Orm, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err!= nil {
		log.Panicln(err)
	}

	_ = Orm.AutoMigrate(&authModel.SsoPaUsers{})
	_ = Orm.AutoMigrate(&reportTemplateModel.ReportTemplate{})
	_ = Orm.AutoMigrate(&reportTemplateModel.ReportTemplateVar{})
}