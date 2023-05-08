package controllers

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
)

type ContactUsController struct {
}

func (p *ContactUsController) Get() mvc.Result {
	return mvc.View{
		Name: "ContactUs/ContactUs.html",
		Data: iris.Map{"curnav": "contactus"},
	}
}
