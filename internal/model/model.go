package model

import (
	"fmt"
	"my-service/pkg/setting"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID uint32 `gorm:"primary_key" json:"id"`
}

type SwaggerSuccess struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	CheckpointThyroid = 0
	CheckpointBreast  = 1
	CheckpointHeart   = 2
)

var (
	ImageTypeScan   = 0
	ImageTypeFreeze = 1
)

func NewDBEngine(dbSettings *setting.DatabaseSettings) (*gorm.DB, error) {
	// dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
	// 	dbSettings.UserName,
	// 	dbSettings.Password,
	// 	dbSettings.Host,
	// 	dbSettings.DBName,
	// 	dbSettings.Charset,
	// 	dbSettings.ParseTime,
	// )
	// fmt.Println("dsn", dsn)
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	db, err := gorm.Open(dbSettings.DBType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=%t&loc=Local",
		dbSettings.UserName,
		dbSettings.Password,
		dbSettings.Host,
		dbSettings.DBName,
		dbSettings.Charset,
		dbSettings.ParseTime,
	))
	if err != nil {
		return nil, err
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return dbSettings.TablePrefix + defaultTableName
	}
	db.LogMode(true)
	db.SingularTable(true)
	//db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	//db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	//db.Callback().Delete().Replace("gorm:AfterDelete", deleteCallback)
	db.DB().SetMaxIdleConns(dbSettings.MaxIdleConns)
	db.DB().SetMaxOpenConns(dbSettings.MaxOpenConns)

	return db, nil
}

// func deleteCallback(scope *gorm.Scope) {
// 	if !scope.HasError() {
// 		fmt.Println(scope.DB().RowsAffected)
// 		tbName := scope.QuotedTableName()
// 		if "t_download_file_info" != tbName {
// 			return
// 		}
// 		if filedTmp, ok := scope.FieldByName("downloadTimes"); ok {
// 			fmt.Println(filedTmp.Field)
// 		}
// 	}
// }
