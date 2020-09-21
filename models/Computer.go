package models

import (
	"vpn-server/structures"
	"github.com/astaxie/beego"
	"time"
	"github.com/astaxie/beego/orm"
)

func CreateComputer(c structures.WebComputer) (error bool, id int64, msg string) {
	beego.Debug("add=================")
	beego.Debug(c)
	c.Time = time.Now()
	id, err := o.Insert(&c)
	if err != nil {
		beego.Error(id)
		beego.Error("error:" + err.Error())
		return true, 0, "add error"
	}
	return false, id, "add success"
}

func ReadComputer(c structures.WebComputer) (error bool, id int64, msg string) {
	computer := structures.WebComputer{Computer: c.Computer}

	err := o.Read(&computer, "computer")

	if err == orm.ErrNoRows {
		return false,0, "Could not find"
	} else if err == orm.ErrMissPK {
		return false,0, "No PK"
	} else {
		return false, computer.Id, "Read success"
	}

}

func UpdateComputer(c structures.WebComputer) (error bool, id int64, msg string) {
	computer := structures.WebComputer{Id: c.Id}
	if o.Read(&computer) == nil {
		id := computer.Id
		computer.Id = id
		t := computer.Time
		computer = c
		computer.Time = t
		beego.Debug(computer)
		_, err := o.Update(&computer)
		if err != nil {
			return true, 0, err.Error()
		}
	}
	return false, c.Id, "update success"
}

func ReadRand(city string,status int64) (c structures.WebComputer)  {
	err := o.Raw("SELECT * FROM web_computer WHERE city = ? and status = ? ORDER BY RAND() LIMIT 1", city,status).QueryRow(&c)
	if err !=nil {
		beego.Debug(err.Error())
	}
	return c
}

func UserReadComputer(code string) (c structures.WebComputer)  {
	err := o.Raw("SELECT * FROM web_computer WHERE code = ? and status = 1 ORDER BY RAND() LIMIT 1",code).QueryRow(&c)
	if err !=nil {
		beego.Debug(err.Error())
	}
	return c
}

func ReadAll() (num int64,c []structures.WebComputer)  {
	num, err := o.QueryTable("web_computer").Filter("status", "1").GroupBy("code").OrderBy("code").All(&c)
	//err := o.Raw("SELECT * FROM web_computer WHERE status = 1").QueryRow(&c)
	if err !=nil {
		beego.Debug(err.Error())
	}
	//beego.Debug(num)
	return num,c
}

func ReadUserComputer(code string) (name string)  {
	var webusercomputer structures.Webusercomputer
	err := o.Raw("SELECT name FROM web_usercomputer WHERE code = ?", code).QueryRow(&webusercomputer)
	if err !=nil {
		beego.Debug(err.Error())
	}
	return webusercomputer.Name
}

func ReadComputerIp(c structures.WebComputer) (error bool, ip string, msg string) {
	computer := structures.WebComputer{Computer: c.Computer}

	err := o.Read(&computer, "computer")

	if err == orm.ErrNoRows {
		return false,"0", "Could not find"
	} else if err == orm.ErrMissPK {
		return false,"0", "No PK"
	} else {
		return false, computer.Ip, "Read success"
	}

}