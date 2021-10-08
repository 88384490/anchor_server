package models

import (
	orm "github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

var userName = "root"
var password = "Vkx9r5xy"
var (
	sqlUrl = userName + `:` + password + `@tcp(rm-bp187qy95g848549ulo.mysql.rds.aliyuncs.com:3306)/anchor?charset=utf8`
)

type Anchor struct {
	Id            int
	AnchorName    string
	LastStartTime string
	FanNumber     int
	ChatNumber    int
	Status        bool
}

func init() {
	orm.RegisterModel(new(Anchor))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", sqlUrl)
	orm.RunSyncdb("default", false, true)
}

func QueryAll() []Anchor {
	var anchorData []Anchor
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("anchor")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, 0).QueryRows(&anchorData)
	println(anchorData)
	return anchorData
}
