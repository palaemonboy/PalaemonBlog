package model

import (
	"PalaemonBlog/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB
var err error

func InitDB() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		utils.DBUser,
		utils.DBPassword,
		utils.DBHost,
		utils.DBPort,
		utils.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		// gorm log mode : silent | gorm日志模式：silent
		Logger: logger.Default.LogMode(logger.Silent),
		// Foreign Key Constraints | 外键约束
		DisableForeignKeyConstraintWhenMigrating: true,
		// Disable default transactions (improves runtime) | 禁用默认事务（提高运行速度）
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			// Use the singular table name to enable this option, where the table name of `User` should be `user` | 使用单数表名，启用该选项，此时，`User` 的表名应该是 `user`
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("Connect DB failed,pls check!", err)
	}
	// Migrate data tables, it is recommended that comments are not implemented when there are no data table structure changes | 迁移数据表，在没有数据表结构变更时候，建议注释不执行
	//_ = db.AutoMigrate(&User{}, &Article{}, &Category{})

	sqlDB, _ := db.DB()

	// SetMaxIdleCons 设置连接池中的最大闲置连接数
	sqlDB.SetMaxOpenConns(10)

	// SetMaxOpenCons 设置数据库的最大连接数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetiment 设置连接的最大可复用时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	fmt.Println("Connect DB SUCCESS!")

	Db = db
	return err

}
