package models

import (
	"vpn-server/structures"
	"github.com/astaxie/beego/orm"
)

func ReadIdentity(id uint)(error bool, cityid int, msg string) {
	identity := structures.WebIdentity{Id: int(id)}
	err := o.Read(&identity)

	if err == orm.ErrNoRows {
		return false,0, "Could not find"
	} else if err == orm.ErrMissPK {
		return false,0, "No PK"
	} else {
		return false, int(identity.CityId), "Read success"
	}

}
