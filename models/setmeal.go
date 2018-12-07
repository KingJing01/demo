package models

import (
	out "demo/outmodels"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type SetMeal struct {
	Id                     int       `orm:"column(Id);auto;pk"`
	SetMealName            string    `orm:"column(SetMealName);size(45)"`
	SetMealCode            string    `orm:"column(SetMealCode);size(45)"`
	PermissionText         string    `orm:"column(PermissionText);size(200)"`
	CreationTime           time.Time `orm:"column(CreationTime);type(datetime);time"`
	CreatorUserId          int64     `orm:"column(CreatorUserId);null"`
	LastModificationTime   time.Time `orm:"column(LastModificationTime);type(datetime);time"`
	LastModificationUserId int64     `orm:"column(LastModificationUserId);null"`
	IsDeleted              int       `orm:"column(IsDeleted);0"`
	DeletionTime           time.Time `orm:"column(DeletionTime);type(datetime);time"`
	DeletionUserId         int64     `orm:"column(DeletionUserId);null"`
	SysCode                string    `orm:"column(SysCode);size(20)"`
}

func (t *SetMeal) TableName() string {
	return "setmeal"
}

func init() {
	orm.RegisterModel(new(SetMeal))
}

// 获取列表的信息
func GetSetMealList(setMealName string, sysName string, offset int64, limit int64) (result []out.SetMealInfo, err error) {
	o := orm.NewOrm()
	var sql = "SELECT t1.Id id,t1.SetMealCode set_meal_code, t1.SetMealName set_meal_name, t2.SysName sys_name,t1.PermissionText permission_text FROM setmeal t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode WHERE  t1.IsDeleted = 0"
	conditions := []string{}
	if setMealName != "" {
		conditions = append(conditions, " t1.SetMealName like '%"+setMealName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t2.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " and " + strings.Join(conditions, " and ")
	}
	sql = sql + " limit " + strconv.FormatInt(limit, 10) + "  offset " + strconv.FormatInt(offset, 10)
	_, err = o.Raw(sql).QueryRows(&result)
	return result, err
}

// 统计查询条件的数量
func CountSetMealInfo(setMealName string, sysName string) (total int64) {
	o := orm.NewOrm()
	conditions := []string{}
	var sql = "SELECT count(0) total FROM setmeal t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode WHERE  t1.IsDeleted = 0"
	if setMealName != "" {
		conditions = append(conditions, " t1.SetMealName like '%"+setMealName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t2.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " and " + strings.Join(conditions, " and ")
	}
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	total, _ = strconv.ParseInt(maps[0]["total"].(string), 10, 64)
	return total
}
