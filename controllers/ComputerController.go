package controllers

import (
	"encoding/json"
	"strings"
	"vpn-server/models"
	"vpn-server/structures"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
)

type ComputerController struct {
	beego.Controller
}

// @Title 新增電腦
// @Description 新增電腦
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /add [post]
func (c *ComputerController) AddComputer() {
	var computers structures.WebComputer

	jerr := json.Unmarshal(c.Ctx.Input.RequestBody, &computers)
	if jerr != nil {
		c.Data["json"] = map[string]interface{}{"error": true, "msg": "jeadoraddson error"}
		c.ServeJSON()
		return
	}

	beego.Debug(strings.Split(computers.Computer, "-"))
	name := strings.Split(computers.Computer, "-")
	computers.Code = name[1]
	computers.City = name[0]

	beego.Debug(computers)
	err, id, msg := models.CreateComputer(computers)

	c.Data["json"] = map[string]interface{}{"error": err, "id": id, "msg": msg}
	c.ServeJSON()
}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	false
// @Failure 403 body is Json
// @router /update [post]
func (c *ComputerController) UpdateComputer() {
	var computers structures.WebComputer

	jerr := json.Unmarshal(c.Ctx.Input.RequestBody, &computers)
	if jerr != nil {
		c.Data["json"] = map[string]interface{}{"error": true, "msg": "jeadoraddson error"}
		c.ServeJSON()
		return
	}

	beego.Debug(strings.Split(computers.Computer, "-"))
	name := strings.Split(computers.Computer, "-")
	computers.Code = name[1]
	computers.City = name[0]

	beego.Debug(computers)
	err, id, msg := models.ReadComputer(computers)
	computers.Id = id
	err, id, msg = models.UpdateComputer(computers)

	c.Data["json"] = map[string]interface{}{"error": err, "id": id, "msg": msg}
	c.ServeJSON()

}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /pickup [post]
func (c *ComputerController) PickupComputer() {
	var computers structures.WebComputer
	jerr := json.Unmarshal(c.Ctx.Input.RequestBody, &computers)
	if jerr != nil {
		c.Data["json"] = map[string]interface{}{"error": true, "msg": "jeadoraddson error"}
		c.ServeJSON()
		return
	}

	computers = models.ReadRand(computers.City, 1)

	req := httplib.Post("http://localhost:8081/v1/server/update?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	computers.Status = 2
	req.JSONBody(&computers)

	str, err := req.String()

	if err != nil {
		beego.Error(err)
	}

	beego.Debug(str)

	c.Data["json"] = computers
	c.ServeJSON()

}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /used [post]
func (c *ComputerController) UsedIpAdd() {
	var computer structures.WebComputer
	jerr := json.Unmarshal(c.Ctx.Input.RequestBody, &computer)
	if jerr != nil {
		c.Data["json"] = map[string]interface{}{"error": true, "msg": "jeadoraddson error"}
		c.ServeJSON()
		return
	}

	req := httplib.Post("http://localhost:8081/v1/server/update?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	req.JSONBody(&computer)

	str, err := req.String()

	if err != nil {
		beego.Error(err)
	}

	beego.Debug(str)

	req = httplib.Get("http://" + computer.Ip + ":8080/v1/client/change?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	//req.JSONBody(&change)

	str, err = req.String()

	if err != nil {
		beego.Debug(err)
	}
	beego.Debug(str)

	c.Data["json"] = map[string]interface{}{"error": false, "message": "success"}
	c.ServeJSON()

}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /list [get]
func (c *ComputerController) List() {
	//var computer []structures.WebComputer

	num, call := models.ReadAll()
	//json.Unmarshal(call,&computer)
	userComputer := make([]structures.UserComputer, num)
	//json.Unmarshal(userComputer,call)
	for i := 0; i < int(num); i++ {
		userComputer[i].Code = call[i].Code
		userComputer[i].Name = models.ReadUserComputer(call[i].Code)
	}
	json.Marshal(&userComputer)
	beego.Debug(call)
	c.Data["json"] = userComputer

	c.ServeJSON()
}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /userpickup [post]
func (c *ComputerController) UserPickupComputer() {
	var computers structures.WebComputer
	jerr := json.Unmarshal(c.Ctx.Input.RequestBody, &computers)
	if jerr != nil {
		c.Data["json"] = map[string]interface{}{"error": true, "msg": "jeadoraddson error"}
		c.ServeJSON()
		return
	}

	computers = models.UserReadComputer(computers.Code)

	req := httplib.Post("http://localhost:8081/v1/server/update?token=")
	req.Header("Content-Type", "application/json")
	req.Header("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_0) AppleWebKit/537.36"+
		"(KHTML, like Gecko) Chrome/31.0.1650.57 Safari/537.36")
	computers.Status = 2
	req.JSONBody(&computers)

	str, err := req.String()

	if err != nil {
		beego.Error(err)
	}

	beego.Debug(str)

	c.Data["json"] = computers
	c.ServeJSON()

}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /readcheck [post]
func (c *ComputerController) ReadCheckComputer() {
	var computers structures.WebComputer
	jerr := json.Unmarshal(c.Ctx.Input.RequestBody, &computers)
	if jerr != nil {
		c.Data["json"] = map[string]interface{}{"error": true, "msg": "jeadoraddson error"}
		c.ServeJSON()
		return
	}

	err, ip, msg := models.ReadComputerIp(computers)

	c.Data["json"] = map[string]interface{}{"error": err, "ip": ip, "msg": msg}
	c.ServeJSON()
}
