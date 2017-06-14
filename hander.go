package main

import (
	"fmt"
	"net/http"

	"time"

	"github.com/labstack/echo"
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

	name := recordLog.Name
	content := recordLog.Content
	appId := recordLog.AppId
	version := recordLog.Version

	if name == "" || content == "" || appId == "" || version == "" {
		return c.JSON(400, MapY(400, "error", "Aname are empty"))
	}

	recordLog.Ip = c.RealIP()

	s := MOBS

	cp := NewMgoFun(s, Config.MDB.Name, recordLog)
	cp.Create()

	return c.JSON(http.StatusOK, MapY(200, "ok", recordLog))

}
