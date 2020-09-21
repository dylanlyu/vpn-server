package models

import (
	"vpn-server/structures"
	"github.com/astaxie/beego/orm"
)


func ReadCorrespond(word string) (error bool, c structures.WebCorrespond, msg string) {
	correspond := structures.WebCorrespond{First:word}

	err := o.Read(&correspond, "First")

	if err == orm.ErrNoRows {
		return false,correspond, "Could not find"
	} else if err == orm.ErrMissPK {
		return false,correspond, "No PK"
	} else {
		return false, correspond, "Read success"
	}

}
