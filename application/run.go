package application

import (
	"flag"
	"fmt"
	"atlantis/conf"
	"atlantis/dbs"
	apiRouter "atlantis/router"
	"os"
)

//Run app run
func Run(app conf.App) {
	router := apiRouter.Router(app)
	router.Run(app.Port)
}

func init() {
	mode := flag.String("mode", "dev", "eventment")
	flag.Parse()
	mysql, err := getMysql(*mode)
	if err != nil {
		panic(err)
	}
	dbs.RunMysql(mysql)
	app, err := getApp(*mode)
	fmt.Println("app", app)
	if err != nil {
		panic(err)
	}
	Run(app)
}

func readIni(mode string) (iniParser conf.IniParser, err error) {
	confFileName := "dev.ini"
	if mode == "prod" {
		confFileName = "prod.ini"
	}
	//iniParser = conf.IniParser{}
	dir, _ := os.Getwd()
	pathName := dir + `/conf/` + confFileName
	err = iniParser.Load(pathName)
	return iniParser, err
}

func getApp(mode string) (app conf.App, err error) {
	iniParser, err := readIni(mode)
	if err != nil {
		return app, err
	}
	app = conf.App{
		Mode: mode,
		Port: iniParser.GetString("app", "port"),
	}
	return app, nil
}

func getMysql(mode string) (mysql conf.Mysql, err error) {
	iniParser, err := readIni(mode)
	if err != nil {
		return mysql, err
	}
	mysql = conf.Mysql{
		User:     iniParser.GetString("mysql", "user"),
		Password: iniParser.GetString("mysql", "password"),
		Host:     iniParser.GetString("mysql", "host"),
		Port:     iniParser.GetString("mysql", "port"),
		Database: iniParser.GetString("mysql", "database"),
	}
	return mysql, err
}
