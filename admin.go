package main

import (
	"net/http"

	"fmt"

	"github.com/labstack/echo"
	"gopkg.in/mgo.v2/bson"
)

func AdminLogs(c echo.Context) error {

	recordLog := new(RecordLogs)

	op := NewMgoFun(MOBS, Config.MDB.Name, recordLog)

	var recordLogs []*RecordLogs
	op.FindAll(&recordLogs)

	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), recordLogs))

}

func AdminLog(c echo.Context) error {
	recordLog := new(RecordLogs)

	op := NewMgoFun(MOBS, Config.MDB.Name, recordLog)

	id := c.Param("id")

	recordLog.Id = bson.ObjectIdHex(id)

	op.Query = bson.M{"_id": recordLog.Id}
	op.GetByQ()

	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), recordLog))
}

func AdminLogDel(c echo.Context) error {
	recordLog := new(RecordLogs)
	id := c.FormValue("id")

	fmt.Printf(id)

	if id == "" {
		return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), "id is emp"))
	}

	recordLog.Id = bson.ObjectIdHex(id)

	fmt.Printf(string(recordLog.Id))

	op := NewMgoFun(MOBS, Config.MDB.Name, recordLog)

	op.RemoveDel()

	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), recordLog))
}
