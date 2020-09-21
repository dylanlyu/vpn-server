package main

import (
	_ "vpn-server/docs"
	"vpn-server/models"
	_ "vpn-server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	//beego.SetLogger("file", `{"filename":"logs/test.log"}`)

	var FilterToken = func(ctx *context.Context) {
		//beego.Debug(ctx.Request.Header.Get("Authorization"))
		if ctx.Request.URL.Query()["token"] == nil {
			beego.Info("URL=error")
			beego.Error("URL:" + ctx.Input.URL() + "  " + ctx.Input.IP())
			ctx.Redirect(302, "/")
			return
		}

		key := ""

		if ctx.Request.URL.Query()["token"][0] != key {
			beego.Info("Token=error")
			beego.Error("URL:" + ctx.Input.URL() + "  " + ctx.Input.IP())
			ctx.Redirect(302, "/")
			return
		}

	}
	models.Init()
	beego.InsertFilter("/v1/*", beego.BeforeRouter, FilterToken)
	beego.Run()
}
