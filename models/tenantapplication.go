package models

import "github.com/astaxie/beego/orm"

type TenantApplication struct {
	Id       int    `orm:"column(Id);auto"`
	TenantId int    `orm:"column(TenantId);"`
	SysCode  string `orm:"column(SysCode);size(20);null"`
	MenuText string `orm:"column(MenuText);size(255);null"`
}

func (t *TenantApplication) TableName() string {
	return "tenantapplication"
}
func init() {
	orm.RegisterModel(new(TenantApplication))
}
