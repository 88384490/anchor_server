package models

import (
	"anchor_server/types"
	orm "github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	var (
		user, _    = beego.AppConfig.String("MYSQL_USER")
		pwd, _     = beego.AppConfig.String("MYSQL_PASSWORD")
		sqlcoon, _ = beego.AppConfig.String("MYSQL_HOST")
		port, _    = beego.AppConfig.String("MYSQL_PORT")
		sqlUrls    = user + `:` + pwd + `@tcp(` + sqlcoon + `:` + port + `)/anchor?charset=utf8`
	)
	orm.RegisterModel(new(types.Movie))
	orm.RegisterDataBase("default", "mysql", sqlUrls)
	orm.RunSyncdb("default", false, true)
}

func QueryMovie() []types.Movie {
	var movieData []types.Movie
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("movie")
	sql := qb.String()
	o := orm.NewOrm()
	o.Raw(sql, 0).QueryRows(&movieData)
	println(movieData)
	return movieData
}
