package base

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	"myblog/conf"
	"fmt"
)

var DBMyBlog *gorm.DB




func InitGormDbPool(config *conf.MysqlConfig) (*gorm.DB, error) {

	db, err := gorm.Open("mysql", config.Conn)
	if err != nil {
		fmt.Println("init db err : ", config, err)
		return nil, err
	}

	db.DB().SetMaxOpenConns(config.PoolSize)
	db.DB().SetMaxIdleConns(config.PoolSize / 2)
	if conf.Config.Debug {
		db.LogMode(conf.Config.Debug)
	}

	db.SingularTable(true)

	err = db.DB().Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("init db pool OK")

	return db, nil
}


func InitDB() {
	var err error
  	DBMyBlog, err = InitGormDbPool(&conf.Config.DBMyBlog)
  	if nil != err {
		Logger.Errorf("init feeddb err :%v,%v",  conf.Config.DBMyBlog.Conn, err)
		return
	}

}