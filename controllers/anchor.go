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
	a.Data["json"] = data
	a.ServeJSON()
}

// Get
// @Title QueryAnchors
// @Success 200 {[]} models.Anchor
// @Failure 403 body is empty
// @router /:id [get]
func (a *AnchorController) Get() {
	param := a.Ctx.Input.Param(":id")
	println(a.Ctx.Input)
	if param != "" {
		println(param)
	}
}
