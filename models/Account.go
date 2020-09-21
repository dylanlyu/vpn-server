package models

import (
	"vpn-server/structures"

	"github.com/astaxie/beego"
)



func ReadWebAccount(times int64) (c structures.WebAccount)  {
	err := o.Raw("SELECT * FROM web_account WHERE forum_id = 34 AND iserror = 0 AND istask = 0 AND israise = 1 AND isaction = 0 AND update_time < ? ORDER BY RAND() LIMIT 1",times).QueryRow(&c)
	if err !=nil {
		beego.Debug(err.Error())
	}
	return c
}

func UpdateWebAccountIsaction(id int) (error bool,num int64,mesg string){
	res, err := o.Raw("UPDATE web_account SET isaction = 1 WHERE id = ?", id).Exec()
	if err !=nil {
		beego.Debug(err.Error())
		return true,0,err.Error()
	}
	num, _ = res.RowsAffected()
	//fmt.Println("mysql row affected nums: ", num)
	return false,num,"update success"
}

func UpdateWebAccount(iserror int,updatetime int,id int) (error bool,num int64,mesg string){
	res, err := o.Raw("UPDATE web_account SET isaction = 0,iserror = ? ,update_time = ? WHERE id = ?", iserror,updatetime,id).Exec()
	if err !=nil {
		beego.Debug(err.Error())
		return true,0,err.Error()
	}
	num, _ = res.RowsAffected()
	//fmt.Println("mysql row affected nums: ", num)
	return false,num,"update success"
}