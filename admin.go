package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func AdminLog(c echo.Context) error {

	recordLog := new([]RecordLogs)

	return c.JSON(http.StatusOK, MapY(0, http.StatusText(200), recordLog))

}
