package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

/*
time
*/
func HandTime(c echo.Context) error {
	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), time.Now().UTC()))
}

/*
日志
*/
func HandLog(c echo.Context) error {

	//var recordLog RecordLogs
	recordLog := new(RecordLogs)

	err := c.Bind(recordLog)

	if err != nil {
		fmt.Println("error")
		fmt.Println(err)

	}

	fmt.Println(recordLog)

	if recordLog.Type == "" || recordLog.Content == "" || recordLog.AppId == "" {

		return c.JSON(422, MapY(422, "error", "Fields are empty"))
	}

	recordLog.Ip = c.RealIP()

	db := MDBS.DB("sixCity")

	collection := db.C("record_logs")

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	//一次可以插入多个对象 插入两个Person对象
	err = collection.Insert(recordLog)
	if err != nil {
		panic(err)
	}

	s := MDBS

	user := new(User)
	user.Name = "Tom"
	user.Age = 10
	user.Id = bson.NewObjectId()

	user2 := new(User)
	user2.Name = "Jack"
	user2.Age = 20
	user2.Id = bson.NewObjectId()

	cp := NewMgoFun(s, dbName, recordLog)
	cp.Create()

	//conduct op
	op := NewMgoFun(s, dbName, user)
	op.Create()
	op = NewMgoFun(s, dbName, user2)
	err = op.Create()
	if err != nil {
		fmt.Println("Err during save: ", err)
	}

	return c.JSON(http.StatusOK, MapY(200, "ok", recordLog))

}
