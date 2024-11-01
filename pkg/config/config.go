package config

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var config *viper.Viper

var Appconfig *viper.Viper

func Init() {
	var err error
	config = viper.New()
	config.SetConfigType("yaml")
	config.SetConfigName("config")
	path, err := os.Getwd()
	if err != nil {
		log.Fatal("error while fetching path")
	}
	exepath := strings.Split(path, "\\")
	path = strings.Join(exepath[:len(exepath)-2], "\\")
	fmt.Println(exepath, path)
	config.AddConfigPath(path + "/pkg/config/")
	err = config.ReadInConfig()
	if err != nil {
		log.Fatal("error on parsing configuration file", err.Error())
	}
}

func GetConfig() *viper.Viper {
	return config
}

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
