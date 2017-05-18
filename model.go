package main

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

type AuthKey struct {
	Id  int    `gorm:"primary_key;AUTO_INCREMENT" form:"id" json:"id"`
	Key string `form:"key" json:"key"`
}

type User struct {
	BaseModel `bson:",inline"`
	Name      string `bson:"name,omitempty"`
	Age       int    `bson:"age,omitempty"`
}
