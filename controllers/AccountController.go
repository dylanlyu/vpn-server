package controllers

import (
	"github.com/astaxie/beego"
	"vpn-server/structures"
	"vpn-server/models"
	"encoding/json"
	"time"
)

type AccountController struct {
	beego.Controller
}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /read [get]
func (c *AccountController) ReadAccount() {
	var account structures.WebAccount
	var client structures.ClientAccount
	var correspond structures.WebCorrespond

	//jerr := json.Unmarshal(c.Ctx.Input.RequestBody,&account)
	t := time.Date(time.Now().Year(),time.Now().Month(),time.Now().Day(),0,0,0,0,time.Local)
	beego.Debug(t.Unix())
	account = models.ReadWebAccount(t.Unix())
	_ , _ , _ = models.UpdateWebAccountIsaction(account.Id)
	client.Id = account.Id
	client.Account = account.Account
	client.Pass = account.Passwd

	_ , cityId , _ := models.ReadIdentity(account.IdentityId)
	_ , cityCode , _ := models.ReadCity(cityId)
	_ , correspond ,_ = models.ReadCorrespond(cityCode)

	client.Primary = correspond.First
	client.Secondary = correspond.Second
	client.Tertiary = correspond.Third
	client.IsAuto = account.IsAuto
	client.IsUserSet = account.IsUserSet
	client.Gas = account.Gas
	client.Action = account.Action
	client.IsError = account.IsError
	client.UpdateTime = account.UpdateTime

	json.Marshal(&client)

	c.Data["json"] = client
	c.ServeJSON()
}

// @Title 顯示IP
// @Description 顯示IP
// @Param	body		body 	Json	flase
// @Failure 403 body is Json
// @router /update [post]
func (a *AccountController) UpdateAccount() {
	var client structures.ClientAccount

	jerr := json.Unmarshal(a.Ctx.Input.RequestBody,&client)
	if jerr!=nil {

	}

	err,num,meg := models.UpdateWebAccount(client.IsError,client.UpdateTime,client.Id)

	beego.Debug(err)
	beego.Debug(num)
	beego.Debug(meg)
	a.Data["json"] = map[string]interface{}{"error": err, "message": meg}
	a.ServeJSON()

}
