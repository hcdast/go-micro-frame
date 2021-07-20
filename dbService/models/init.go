package models

import (
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/beego/beego/v2/core/logs"
	//go操作数据库的模块
	_ "github.com/go-sql-driver/mysql"
	"github.com/hcdast/go-micro-frame/mainService/conf"
)

// User 用户 table_name = user
type User struct {
	Id           int           `json:"user_id"`                       //用户编号
	Name         string        `orm:"size(32)"  json:"name"`          //用户昵称
	PasswordHash string        `orm:"size(128)" json:"password"`      //用户密码加密的
	Mobile       string        `orm:"size(11);unique"  json:"mobile"` //手机号
	RealName     string        `orm:"size(32)" json:"real_name"`      //真实姓名  实名认证
	IdCard       string        `orm:"size(20)" json:"id_card"`        //身份证号  实名认证
	AvatarUrl    string        `orm:"size(256)" json:"avatar_url"`    //用户头像路径       通过fastdfs进行图片存储
}

func init() {
	// 注册mysql的驱动
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	fmt.Println("-------------", conf.MysqlAddr)
	// 设置默认数据库
	_ = orm.RegisterDataBase("default", "mysql", conf.MysqlUser+":"+conf.MysqlPass+"@tcp("+conf.MysqlAddr+":"+conf.MysqlPort+")/"+conf.MysqlDbname+"?charset=utf8")
	// 注册model
	orm.RegisterModel(new(User))

	// create table
	// 第一个是别名
	// 第二个是是否强制替换模块 如果表变更就将false 换成true 之后再换回来表就便更好来了
	// 第三个参数是如果没有则同步或创建
	_ = orm.RunSyncdb("default", false, true)
}
