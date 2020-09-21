// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"vpn-server/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func init() {
	beego.Any("/", func(ctx *context.Context) {
		ctx.Output.Body([]byte(``))
	})

	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/server",
			beego.NSInclude(
				&controllers.ComputerController{},
			),
		),
		beego.NSNamespace("/account",
			beego.NSInclude(
				&controllers.AccountController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
