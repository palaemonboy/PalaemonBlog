package utils

import (
	"fmt"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string
	JwtKey   string

	DB         string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
)

func init() {
	file, err := ini.Load("config/config.ini")
	if err != nil {
		fmt.Println("Can't read config file.", err)
	}
	LoadServer(file)
	LodaData(file)
}
func LoadServer(file *ini.File) {
	AppMode = file.Section("server").Key("AppMode").MustString("debug")
	HttpPort = file.Section("server").Key("HttpPort").MustString(":3000")
	JwtKey = file.Section("server").Key("JwtKey").MustString("s22717gmy")
}

func LodaData(file *ini.File) {
	DB = file.Section("database").Key("DB").MustString("mysql")
	DBHost = file.Section("database").Key("DBHost").MustString("192.168.50.102")
	DBPort = file.Section("database").Key("DBPort").MustString("3306")
	DBUser = file.Section("database").Key("DBUser").MustString("root")
	DBPassword = file.Section("database").Key("DBPassword").MustString("www.nas.beer")
	DBName = file.Section("database").Key("DBName").MustString("test")
}
