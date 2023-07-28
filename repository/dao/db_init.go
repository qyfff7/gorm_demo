package dao

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm_demo/config"
	"gorm_demo/repository/models"
	"strings"
)

func InitDB() *gorm.DB {
	// 从viper获取配置
	mysqlConfig := config.Conf.MySQL
	username := mysqlConfig.UserName
	password := mysqlConfig.Password
	host := mysqlConfig.Host
	port := mysqlConfig.Port
	dbname := mysqlConfig.DbName
	charset := mysqlConfig.Charset
	parseTime := mysqlConfig.ParseTime
	loc := mysqlConfig.Loc

	// 根据配置拼接dsn
	dsn := strings.Join([]string{username, ":", password, "@tcp(", host, ":", port, ")/", dbname, "?charset=", charset, "&parseTime=", parseTime, "&loc=", loc}, "")
	db, err := Database(dsn) // 传入dsn获取db模型
	if err != nil {
		panic(err)
	}
	return db
}

// Database 根据dsn返回gorm.DB实例
func Database(conn string) (*gorm.DB, error) {
	// 开启连接并设置配置文件
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               conn,
		DefaultStringSize: 256,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: false, // 为true的话，不会给表加yes，false的话，不用写，这里只是列出来做提醒
		},
	})
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(models.User{}) // 自动迁移
	if err != nil {
		return nil, err
	}
	return db, err
}
