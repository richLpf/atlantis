package conf

import (
	"gopkg.in/ini.v1"
)

//App 参数
type App struct {
	Mode string
	Port string
}

//Mysql 参数
type Mysql struct {
	User     string
	Host     string
	Port     string
	Password string
	Database string
}

type IniParser struct {
	conf_reader *ini.File
}

type IniParseError struct {
	error_info string
}

func (e *IniParseError) Error() string { return e.error_info }

func (this *IniParser) Load(config_file_name string) error {
	conf, err := ini.Load(config_file_name)
	if err != nil {
		this.conf_reader = nil
		return err
	}
	this.conf_reader = conf
	return nil
}

func (this *IniParser) GetString(section string, key string) string {
	if this.conf_reader == nil {
		return ""
	}
	s := this.conf_reader.Section(section)
	if s == nil {
		return ""
	}
	return s.Key(key).String()
}
