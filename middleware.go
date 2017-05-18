package main

import (
	"fmt"

	"github.com/labstack/echo"
)

// MiddlewareTime is auth key
func MiddlewareTime(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := new(AuthKey)
		err := c.Bind(authKey)
		if err != nil {
			fmt.Println("authkey error")
			fmt.Println(err)
		}
		if authKey.Key != Config.Key {
			return c.JSON(400, MapY(400, "error", "Akey are error"))
		}

		return next(c)
	}
}

// MiddlewareKey is auth key
func MiddlewareKey(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authKey := new(AuthKey)
		key := c.QueryParam("key")
		authKey.Key = key

		if authKey.Key != Config.Key {
			return c.JSON(400, MapY(400, "error", "Akey are empty"))
		}

		return next(c)
	}
}

// MiddlewareLog is post log
func MiddlewareLog(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//recordLog.Type == "" || recordLog.Content == "" || recordLog.AppId == ""
		name := c.FormValue("name")
		content := c.FormValue("content")
		appId := c.FormValue("app_id")
		version := c.FormValue("version")

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
