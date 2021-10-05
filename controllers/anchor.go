package controllers

import (
	"anchor_server/models"
	beego "github.com/beego/beego/v2/server/web"
)

type AnchorController struct {
	beego.Controller
}

// Post
// @Title CreateAnchor
// @Success 200 {[]} models.Anchor
// @Failure 403 body is empty
// @router / [post]
func (a *AnchorController) Post() {
	data := models.QueryAll()
	println(data)
	a.Data["json"] = data
	a.ServeJSON()
}
