package models

import (
	"vpn-server/structures"
	"github.com/astaxie/beego/orm"
)


func ReadCity(id int)(error bool, code string, msg string) {
	city := structures.WebCity{Id: id}

	err := o.Read(&city)

	if err == orm.ErrNoRows {
		return false,"0", "Could not find"
	} else if err == orm.ErrMissPK {
		return false,"0", "No PK"
	} else {
		return false, city.Code, "Read success"
	}

}

