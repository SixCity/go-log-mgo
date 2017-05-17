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
}{}

//
//func Handle(idx int, mpool *mongopool.Pool, done func()) {
//	defer done()
//	err := handle(idx, mpool)
//	if err != nil {
//		log.Printf("ERROR (worker %d): %v", idx, err)
//	}
//}
//
//func handle(idx int, mpool *mongopool.Pool) error {
//	db, err := mpool.Acquire()
//	if err != nil {
//		return err
//	}
//	defer mpool.Release(db)
//
//	log.Printf("Worker %d proceeding.", idx)
//
//	err = db.C("foo").Insert(bson.M{"i": idx})
//	if err != nil {
//		return err
//	}
//
//	// Slow things down for more informative output.
//	time.Sleep(500 * time.Millisecond)
//
//	return nil
//}

var (
	dbName = "mgofun_test"
	dial   = "localhost"
)

type Person struct {
	NAME  string
	PHONE string
}
type Men struct {
	Persons []Person
}

var MDBS *mgo.Session

func init() {

	configor.Load(&Config, "config.yml")

	fmt.Printf("config : %#v", Config)

	uDB := Config.DB

	uri := uDB.User + ":" + uDB.Password + "@tcp(" + uDB.Host + ":" + uDB.Port + ")/" + uDB.Name + "?" + uDB.Option

	fmt.Println(uri)

	MDBS, _ = mgo.Dial("localhost:27017")

	//defer MDBS.Close()
	MDBS.SetMode(mgo.Monotonic, true)

	//db := MDBS.DB("sixCity")
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
	////*******插入元素*******
	//temp := &Person{
	//	PHONE: "18811577546",
	//	NAME:  "zhangzheHero",
	//}
	////一次可以插入多个对象 插入两个Person对象
	//err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"}, temp)
	//if err != nil {
	//	panic(err)
	//}

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
