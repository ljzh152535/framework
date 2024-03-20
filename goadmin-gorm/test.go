package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strings"
	"time"
)

func dbConfig() *gorm.DB {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:               "root:123mysql@tcp(127.0.0.1:3306)/go_admin?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 171,                                                                                   // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction: false, // 是否跳过默认事务
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "t_",                              //表名前缀，`User`的表为`t_users`
			SingularTable: false,                             //使用单数表名，启用此选项后，`User`的表将为`user`
			NoLowerCase:   true,                              //跳过名称的snake_casing
			NameReplacer:  strings.NewReplacer("CID", "Cid"), //在将结构/字段名称转换为数据库名称之前，使用name replacer更改结构/字段名称
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 在 AutoMigrate 或 CreateTable 时，GORM 会自动创建外键约束，若要禁用该特性，可将其设置为 true
		// 逻辑外键 (在代码里自动体现外键关系)
	})
	fmt.Println(db, err)

	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
	return db

}

type User struct {
	Name string
}

func main() {
	db := dbConfig()

	// 创建表
	db.AutoMigrate(&User{})
}
