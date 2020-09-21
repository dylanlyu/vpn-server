package models

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"vpn-server/structures"
)

var (
	o orm.Ormer
)

func Init() {
	//beego.Info("Open DB connect!")

	fmt.Println("Open DB connect!")

	db_host := beego.AppConfig.String("db_host")
	db_port := beego.AppConfig.String("db_port")
	db_user := beego.AppConfig.String("db_user")
	db_pass := beego.AppConfig.String("db_pass")
	db_name := beego.AppConfig.String("db_name")
	//db_sslmode := beego.AppConfig.String("db_sslmode")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(structures.WebComputer),new(structures.WebAccount),
		new(structures.WebIdentity),new(structures.WebCity),new(structures.WebCorrespond),new(structures.Webusercomputer))

	dns := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8",
		db_user,db_pass,db_host,db_port,db_name)

	//dns := fmt.Sprintf("%s:%s@unix(/tmp/mysql.sock)/%s?charset=utf8",
	//	db_user,db_pass,db_name)

	fmt.Println(dns)


	orm.RegisterDataBase("default", "mysql", dns)
	if beego.BConfig.RunMode == "dev" {
		orm.Debug = true
	}
	o = orm.NewOrm()
	o.Using(db_name)
}
