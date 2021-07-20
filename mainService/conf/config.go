package conf

import (
	"github.com/beego/beego/v2/core/config"
	"github.com/beego/beego/v2/core/logs"

	//使用了beego框架的配置文件读取模块
)

var (
	ServerName  string //项目名称
	ServerAddr  string //服务器ip地址
	serverPort  string //服务器端口
	RedisAddr   string //redis ip地址
	RedisPort   string //redis port端口
	RedisDbnum  string //redis db 编号
	MysqlAddr   string //mysql ip 地址
	MysqlPort   string //mysql 端口
	MysqlDbname string //mysql db name
	MysqlUser	string //mysql db user
	MysqlPass	string //mysql db pass
)

func InitConfig() {
	//从配置文件读取配置信息
	appconf, err := config.NewConfig("ini", "./conf/app.conf")
	if err != nil {
		logs.Debug(err)
		return
	}
	ServerName, _ = appconf.String("appname")
	ServerAddr, _ = appconf.String("httpaddr")
	serverPort, _ = appconf.String("httpport")
	RedisAddr, _ = appconf.String("redisaddr")
	RedisPort, _ = appconf.String("redisport")
	RedisDbnum, _ = appconf.String("redisdbnum")
	MysqlAddr, _ = appconf.String("mysqladdr")
	MysqlPort, _ = appconf.String("mysqlport")
	MysqlDbname, _ = appconf.String("mysqldbname")
	MysqlUser, _ = appconf.String("mysqluser")
	MysqlPass, _ = appconf.String("mysqlpass")
	return
}

func init() {
	InitConfig()
}