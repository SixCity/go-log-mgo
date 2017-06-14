## /api/log

?key=123

```go
name := c.FormValue("name")
content := c.FormValue("content")
appId := c.FormValue("app_id")
version := c.FormValue("version")
```


```go
type RecordLogs struct {
	ExModel `bson:",inline"`
	// 记录类型
	Name string `form:"name" query:"name" json:"name"`
	// 客户端栈
	Model string `form:"model" query:"model" json:"model"`
	// 客户端记录内容
	Content string `form:"content" query:"content"  json:"content"`
	// 客户端名称
	AppId string `bson:"app_id,omitempty" form:"app_id" query:"app_id" json:"app_id"`
	// 客户端版本号
	Version string `form:"version" query:"version" json:"version"`
	// 客户端 IP
	Ip string `form:"ip" query:"ip" json:"ip"`
	// 客户端用户信息
	User string `form:"user" query:"user" json:"user"`
	// 服务返回信息
	Serves string `form:"serves" query:"serves" json:"serves"`
	// 客户提交信息
	Client string `form:"client" query:"client" json:"client"`
}
```