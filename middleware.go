package main

import (
	"encoding/base64"
	"fmt"

	"github.com/labstack/echo"
)

// MiddlewareTime is auth key
func MiddlewareTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.QueryParam("iskey")
		decodeBytesKey, err := base64.StdEncoding.DecodeString(authKey)

		fmt.Printf(authKey)

		fmt.Print("\n")

		if err != nil || string(decodeBytesKey) != Config.Key {
			return c.JSON(400, MapY(400, "error", "A iskey are error"))
		}

		return next(c)
	}
}

// MiddlewareKey is auth key
func MiddlewareKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := c.QueryParam("iskey")
		decodeBytesKey, err := base64.StdEncoding.DecodeString(authKey)

		fmt.Printf(authKey)

		fmt.Print("\n")

		if err != nil || string(decodeBytesKey) != Config.Key {
			return c.JSON(400, MapY(400, "error", "A admin iskey are error"))
		}

		return next(c)
	}
}

// MiddlewareLog is post log
func MiddlewareLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := new(RecordLogs)
		err := c.Bind(authKey)
		if err != nil {
			fmt.Println("RecordLogs error")
			fmt.Println(err)
		}

		fmt.Print("\n")
		//recordLog.Type == "" || recordLog.Content == "" || recordLog.AppId == ""
		name := authKey.Name
		content := authKey.Content
		appId := authKey.AppId
		version := authKey.Version

		if name == "" || content == "" || appId == "" || version == "" {
			return c.JSON(400, MapY(400, "error", "Aname are empty"))
		}

		return next(c)
	}
}

// ServerHeader is Response Header
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "city/1.1")
		c.Response().Header().Set("Api-Version", "1.0.0")
		return next(c)
	}
}
func NotFind(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		return next(c)

	}
}
