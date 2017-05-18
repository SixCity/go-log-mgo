package main

import (
	"net/http"

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

	c.Bind(recordLog)

	op := NewMgoFun(MOBS, Config.MDB.Name, recordLog)

	op.Remove()

	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), recordLog))
}
