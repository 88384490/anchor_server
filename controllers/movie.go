package controllers

import (
	"anchor_server/models"
	beego "github.com/beego/beego/v2/server/web"
)

type MovieController struct {
	beego.Controller
}

func (m *MovieController) Post() {
	data := models.QueryMovie()
	m.Data["json"] = data
	m.ServeJSON()
}
