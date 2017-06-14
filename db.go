package main

import (
	"fmt"

	"github.com/jinzhu/configor"
	"gopkg.in/mgo.v2"
)

var Config = struct {
	APPName string `default:"app name"`
	Port    string `defaule:"1222"`
	DB      struct {
		Name     string
		User     string `default:"root"`
		Password string `required:"true" env:"DBPassword"`
		Port     string `default:"3306"`
		Host     string `default:"127.0.0.1"`
		Option   string
	}
	MDB struct {
		Name string
		URL  string `default:""`
		Port string
	}
	Key string
}{}

var MOBS *mgo.Session

func init() {

	configor.Load(&Config, "config.yml")

	fmt.Printf("config : %#v", Config)

	mongoURL := Config.MDB.URL + Config.MDB.Port

	MOBS, _ = mgo.Dial(mongoURL)

	MOBS.SetMode(mgo.Monotonic, true)
}
