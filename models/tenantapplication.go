package models

import "github.com/astaxie/beego/orm"

type TenantApplication struct {
	Id        int    `orm:"column(Id);auto"`
	TenantId  int64  `orm:"column(TenantId);"`
	SysCode   string `orm:"column(SysCode);size(20);null"`
	MenuText  string `orm:"column(MenuText);size(255);null"`
	TransType string `orm:"column(TransType);"`
}

func (t *TenantApplication) TableName() string {
	return "tenantapplication"
}
func init() {
	orm.RegisterModel(new(TenantApplication))
}

//更新套餐关系表的菜单字段
func UpdateTenatMenuText(sysCode string, perMenu string, tenId int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE tenantapplication SET MenuText =? WHERE	TenantId =? AND SysCode =? ", perMenu, tenId, sysCode).Exec()
	return err
}

//GetTransType 获取运输类型
func GetTransType(sysCode string, tenId int64) (tenApp TenantApplication, err error) {
	o := orm.NewOrm()
	err = o.Raw("select * from tenantapplication where TenantId =? AND SysCode =?", tenId, sysCode).QueryRow(&tenApp)
	return tenApp, err
}
