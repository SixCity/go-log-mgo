package main

import (
	"fmt"

	"github.com/jinzhu/configor"
	"gopkg.in/mgo.v2"
)

var Config = struct {
	APPName string `default:"app name"`

	DB struct {
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
}{}

var MOBS *mgo.Session

func init() {

	configor.Load(&Config, "config.yml")

	fmt.Printf("config : %#v", Config)

	MOBS, _ = mgo.Dial("localhost:27017")

	MOBS.SetMode(mgo.Monotonic, true)

	//db := MOBS.DB("sixCity")
	//
	//collection := db.C("record_logs")
	//
	////*****集合中元素数目********
	//countNum, err := collection.Count()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Things objects count: ", countNum)
	//

	//*****查询单条数据*******
	//result := Person{}
	//err = collection.Find(bson.M{"phone": "456"}).One(&result)
	//fmt.Println("Phone:", result.NAME, result.PHONE)

	//*****查询多条数据*******
	//var personAll Men //存放结果
	//iter := collection.Find(nil).Iter()
	//for iter.Next(&result) {
	//	fmt.Printf("Result: %v\n", result.NAME)
	//	personAll.Persons = append(personAll.Persons, result)
	//}

	//*******更新数据**********
	//err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
	//err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
	//err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

	//******删除数据************
	//_, err = collection.RemoveAll(bson.M{"name": "Ale”}),

}
