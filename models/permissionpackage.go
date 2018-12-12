package models

import (
	out "demo/outmodels"

	"github.com/astaxie/beego/orm"
)

type PermissionPackage struct {
	Id             int    `orm:"column(Id);auto"`
	PermissionCode string `orm:"column(PermissionCode);size(45);null"`
	SetMealCode    string `orm:"column(SetMealCode);size(45);null"`
}

func (t *PermissionPackage) TableName() string {
	return "permissionpackage"
}
func init() {
	orm.RegisterModel(new(PermissionPackage))
}

//	AddPermissionPackage 新增权限信息
func AddPermissionPackage(p *PermissionPackage) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(p)
	return
}

func GetPermBySetMealCode(setMeatCode string, sysCode string) (result []out.PermissionCheckInfo, err error) {
	o := orm.NewOrm()
	_, err = o.Raw(`SELECT t3.NAME name,t3.DisplayName display_name ,GROUP_CONCAT(t1. NAME) code,GROUP_CONCAT(t1.DisplayName) code_name,
		GROUP_CONCAT(CASE WHEN t2.PermissionCode IS NULL THEN 	0 	ELSE 1 END ) flag FROM (SELECT  Id, NAME,DisplayName,MenuCode FROM
		permission WHERE 	SysCode = ? AND isMenu = 1 ) t1 LEFT JOIN (
		SELECT
			PermissionCode
		FROM
			permissionpackage
		WHERE
			SetMealCode = ?
	) t2 ON t1. NAME = t2.PermissionCode
	LEFT JOIN (SELECT
				NAME,
				DisplayName,
				MenuCode
			FROM
				permission
			WHERE
				SysCode = ?
			AND isMenu = 0) t3 on t1.MenuCode = t3.MenuCode
	GROUP BY
		t1.MenuCode order by t1.MenuCode,t1.Id asc`, sysCode, setMeatCode, sysCode).QueryRows(&result)
	return result, err
}
