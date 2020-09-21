package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["vpn-server/controllers:AccountController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:AccountController"],
		beego.ControllerComments{
			"ReadAccount",
			`/read`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:AccountController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:AccountController"],
		beego.ControllerComments{
			"UpdateAccount",
			`/update`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"AddComputer",
			`/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"UpdateComputer",
			`/update`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"PickupComputer",
			`/pickup`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"UsedIpAdd",
			`/used`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"List",
			`/list`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"UserPickupComputer",
			`/userpickup`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"] = append(beego.GlobalControllerRouter["vpn-server/controllers:ComputerController"],
		beego.ControllerComments{
			"ReadCheckComputer",
			`/readcheck`,
			[]string{"post"},
			nil})

}
