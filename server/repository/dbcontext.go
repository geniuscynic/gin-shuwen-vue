package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var dbClient *gorm.DB

func GetDbContext() *gorm.DB {
	if dbClient == nil {
		dbClient = initDbContext()
	}

	return dbClient
}

func InitDbAutoMigrate() {
	GetDbContext().AutoMigrate(&User{})
	GetDbContext().AutoMigrate(&Role{})
	GetDbContext().AutoMigrate(&UserRole{})
}

func initDbContext() *gorm.DB {

	return initDbMySqlContext()

}

func initDbMySqlContext() *gorm.DB {
	dsn := "root:123456@tcp(127.0.0.1:3306)/xjm_admin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil
	}

	sqlDb, _ := db.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(10)

	return db
}
