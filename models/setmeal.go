package models

import (
	input "demo/inputmodels"
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
	PermissionText         string    `orm:"column(PermissionText);size(1000)"`
	CreationTime           time.Time `orm:"column(CreationTime);type(datetime);time"`
	CreatorUserId          int64     `orm:"column(CreatorUserId);null"`
	LastModificationTime   time.Time `orm:"column(LastModificationTime);type(datetime);time"`
	LastModificationUserId int64     `orm:"column(LastModificationUserId);null"`
	IsDeleted              int       `orm:"column(IsDeleted);0"`
	DeletionTime           time.Time `orm:"column(DeletionTime);type(datetime);time"`
	DeletionUserId         int64     `orm:"column(DeletionUserId);null"`
	SysCode                string    `orm:"column(SysCode);size(20)"`
}

//TableName 设置表名
func (t *SetMeal) TableName() string {
	return "setmeal"
}

func init() {
	orm.RegisterModel(new(SetMeal))
}

//GetSetMealList 获取套餐列表的信息
func GetSetMealList(setMealName string, sysName string, offset int64, limit int64) (result []out.SetMealInfo, err error) {
	o := orm.NewOrm()
	var sql = "SELECT t1.Id id,t1.SetMealCode set_meal_code, t1.SetMealName set_meal_name, t2.SysName sys_name,t1.PermissionText permission_text,t1.IsDeleted is_deleted,t1.SysCode sys_code FROM setmeal t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode "
	conditions := []string{}
	if setMealName != "" {
		conditions = append(conditions, " t1.SetMealName like '%"+setMealName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t2.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " where " + strings.Join(conditions, " and ")
	}
	sql = sql + " limit " + strconv.FormatInt(limit, 10) + "  offset " + strconv.FormatInt(offset, 10)
	_, err = o.Raw(sql).QueryRows(&result)
	return result, err
}

//CountSetMealInfo 统计查询条件的数量
func CountSetMealInfo(setMealName string, sysName string) (total int64) {
	o := orm.NewOrm()
	conditions := []string{}
	var sql = "SELECT count(0) total FROM setmeal t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode "
	if setMealName != "" {
		conditions = append(conditions, " t1.SetMealName like '%"+setMealName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t2.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " where " + strings.Join(conditions, " and ")
	}
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	total, _ = strconv.ParseInt(maps[0]["total"].(string), 10, 64)
	return total
}

//GenerateSetMeatCode 生成套餐编号
func GenerateSetMeatCode() (SetMealCode string) {
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw("select IFNULL(MAX(SetMealCode),'1000')+1  setMealCode from setmeal").Values(&maps)
	return maps[0]["setMealCode"].(string)
}

//AddSetMeal 新增套餐
func AddSetMeal(setMeatInfo *input.SetMeatInput, userID int64) (id int64, err error) {
	//生成套餐编号
	setMeatCode := GenerateSetMeatCode()
	// 套餐表
	setMeal := new(SetMeal)
	setMeal.SetMealName = setMeatInfo.SetMealName
	setMeal.SetMealCode = setMeatCode
	setMeal.PermissionText = setMeatInfo.PerName
	setMeal.CreationTime = time.Now()
	setMeal.CreatorUserId = userID
	setMeal.SysCode = setMeatInfo.SysCode
	o := orm.NewOrm()
	id, err = o.Insert(setMeal)
	length := len(setMeatInfo.PerId)
	//权限套餐关系数据录入
	var permission []PermissionPackage
	arr := strings.Split(setMeatInfo.PerId, ",")
	for _, per := range arr {
		var permObject PermissionPackage
		permObject.PermissionCode = per
		permObject.SetMealCode = setMeatCode
		permission = append(permission, permObject)
	}
	id, err = o.InsertMulti(length, permission)
	return id, err
}

//DeleteSetMeal 禁用套餐信息
func DeleteSetMeal(ids string, userID int64) (err error) {
	arr := strings.Split(ids, ",")
	var param string
	for _, x := range arr {
		param += x + ","
	}
	length := len(param) - 1
	params := param[0:length]
	var sql = "update setmeal set IsDeleted=1 , DeletionTime = ? ,DeletionUserId = ? where Id in ( " + params + ")"
	o := orm.NewOrm()
	_, err = o.Raw(sql, time.Now(), userID).Exec()
	return
}

//UpdateSetMeal 更新套餐信息
func UpdateSetMeal(setMeatInfo *input.SetMeatInput, userID int64) (id int64, err error) {
	o := orm.NewOrm()
	o.Raw("update setmeal set SetMealName=?,SysCode=?,LastModificationTime=?,PermissionText=?,LastModificationUserId=? where Id=? ", setMeatInfo.SetMealName, setMeatInfo.SysCode, time.Now(), setMeatInfo.PerName, userID, setMeatInfo.Id).Exec()
	o.Raw("delete  from permissionpackage where SetMealCode= ?", setMeatInfo.SetMealCode).Exec()
	//权限套餐关系数据录入
	var permission []PermissionPackage
	arr := strings.Split(setMeatInfo.PerId, ",")
	for _, per := range arr {
		var permObject PermissionPackage
		permObject.PermissionCode = per
		permObject.SetMealCode = setMeatInfo.SetMealCode
		permission = append(permission, permObject)
	}
	id, err = o.InsertMulti(len(arr), permission)
	return id, err
}

//GetSetMealRadio 获取套餐信息 radio
func GetSetMealRadio(sysCodes string) (data []out.PerInfo, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("select SetMealCode name,SetMealName display_name from setmeal where SysCode=?", sysCodes).QueryRows(&data)
	return data, err
}

//UpdateSetMealStatus  更新套餐的状态
func UpdateSetMealStatus(id string, status string, userID int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("update setmeal set IsDeleted=? , DeletionTime = ? ,DeletionUserId = ? where Id =?", status, time.Now(), userID, id).Exec()
	return err
}
