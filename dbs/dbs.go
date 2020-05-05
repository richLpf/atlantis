package dbs

import (
	"fmt"
	"go-project-init/conf"
	"go-project-init/model"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	//DB 数据库变量
	DB  *gorm.DB
	err error
)

//RunMysql 启动mysql
func RunMysql(mysql conf.Mysql) {
	connectSQL := mysql.User + ":" + mysql.Password + "@tcp(" + mysql.Host + ":" + mysql.Port + ")/" + mysql.Database + "?charset=utf8mb4,utf8&parseTime=true&loc=Local"
	fmt.Println("connet", connectSQL)
	DB, err = gorm.Open("mysql", connectSQL)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB", DB)
	DB.Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8").AutoMigrate(
		&model.User{},
	)
	DB.DB().Ping()
}
