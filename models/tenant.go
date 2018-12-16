package models

import (
	out "demo/outmodels"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Tenant struct {
	Id                   int       `orm:"column(Id);auto"`
	TenantName           string    `orm:"column(TenantName);size(64)"`
	TenantAddress        string    `orm:"column(TenantAddress);size(200)"`
	OrganizationCode     string    `orm:"column(OrganizationCode);size(45)"`
	BusinessLisenceUrl   string    `orm:"column(BusinessLisenceUrl);size(200)"`
	TaxFileNumber        string    `orm:"column(TaxFileNumber);size(40)"`
	LinkMan              string    `orm:"column(LinkMan);size(45)"`
	LinkPhone            string    `orm:"column(LinkPhone);size(45)"`
	Email                string    `orm:"column(Email);size(45)"`
	IsDeleted            int8      `orm:"column(IsDeleted);0"`
	CreationTime         time.Time `orm:"column(CreationTime);type(datetime)"`
	CreatorUserId        int64     `orm:"column(CreatorUserId);null"`
	LastModificationTime time.Time `orm:"column(LastModificationTime);type(datetime);"`
	LastModifierUserId   int64     `orm:"column(LastModifierUserId);null"`
	DeleterUserId        int64     `orm:"column(DeleterUserId);null"`
	DeletionTime         time.Time `orm:"column(DeletionTime);type(datetime);"`
}

type TenantInput struct {
	tenant  Tenant
	perId   string
	perName string
	sysCode string
}

func (t *Tenant) TableName() string {
	return "tenant"
}

func init() {
	orm.RegisterModel(new(Tenant))
}

// AddTenant insert a new Tenant into database and returns
// last inserted Id on success.
func AddTenant(m *Tenant) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTenantById retrieves Tenant by Id. Returns error if
// Id doesn't exist
func GetTenantById(id int) (v *Tenant, err error) {
	o := orm.NewOrm()
	v = &Tenant{Id: id}
	//v.IsDeleted = 0
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// UpdateTenant updates Tenant by Id and returns error if
// the record to be updated doesn't exist
func UpdateTenantById(m *Tenant) (err error) {
	o := orm.NewOrm()
	v := Tenant{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		m.CreationTime = v.CreationTime
		_, err = o.Update(m)
	}
	// 删除原先已有的权限信息

	//新增修改的权限信息
	return
}

// DeleteTenant deletes Tenant by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTenant(id int) (err error) {
	o := orm.NewOrm()
	v := Tenant{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Raw("update tenant set IsDeleted=1 ,DeletionTime = ?  where Id= ? ", time.Now(), id).Exec()
	}
	return
}

// 获取列表的信息
func GetTenantList(tenantName string, sysName string, offset int64, limit int64) (result []out.UserManageInfo, err error) {
	o := orm.NewOrm()
	var sql = `SELECT t1.Id id,t1.TenantName tenant_name,t3.SysName sys_name,t2.MenuText menu_text,t4.Name operator,t2.SysCode sys_code FROM tenant t1 LEFT JOIN tenantapplication t2 ON t1.Id = t2.TenantId
		left join application t3 on t2.SysCode = t3.SysCode LEFT JOIN user t4 on t1.CreatorUserId = t4.id `
	conditions := []string{}
	if tenantName != "" {
		conditions = append(conditions, " t1.TenantName like '%"+tenantName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t3.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " where " + strings.Join(conditions, " and ")
	}
	sql = sql + " limit " + strconv.FormatInt(limit, 10) + "  offset " + strconv.FormatInt(offset, 10)
	_, err = o.Raw(sql).QueryRows(&result)
	return result, err
}

// 统计查询条件的数量
func CountTenantInfo(tenantName string, sysName string) (total int64) {
	o := orm.NewOrm()
	conditions := []string{}
	var sql = `SELECT count(0) total FROM tenant t1 LEFT JOIN tenantapplication t2 ON t1.Id = t2.TenantId
	left join application t3 on t2.SysCode = t3.SysCode LEFT JOIN user t4 on t1.CreatorUserId = t4.id `
	if tenantName != "" {
		conditions = append(conditions, " t1.TenantName like '%"+tenantName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t3.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " where " + strings.Join(conditions, " and ")
	}
	var maps []orm.Params
	o.Raw(sql).Values(&maps)
	total, _ = strconv.ParseInt(maps[0]["total"].(string), 10, 64)
	return total
}

// 获取企业所拥有的所有权限
func GetPerInfoForTenant(sysCode string, tenantId int) (result []out.PermissionCheckInfo, err error) {
	o := orm.NewOrm()
	_, err = o.Raw(`SELECT t5.DisplayName display_name ,t5.NAME name,GROUP_CONCAT(t5.perName) code_name,	GROUP_CONCAT(t5.perId) code,GROUP_CONCAT(t5.flag) flag
	FROM (SELECT t3.DisplayName, t3.MenuCode as NAME,t1.DisplayName perName,t1. NAME perId,CASE 	WHEN t2. NAME IS NULL THEN 	0 ELSE 1 END flag
	FROM (SELECT 	MenuCode,	DisplayName,NAME FROM	permission WHERE SysCode = ? AND IsMenu = 1 ) t1
	LEFT JOIN ( SELECT DisplayName,	NAME FROM	permission WHERE 	TenantId = ? AND IsMenu = 1 ) t2 ON t1. NAME = t2. NAME
	LEFT JOIN ( SELECT	MenuCode,	DisplayName,NAME	FROM permission WHERE SysCode = ? 	AND IsMenu = 0 ) t3 ON t3.MenuCode = t1.MenuCode
	GROUP BY 	t1. NAME ) t5 GROUP BY NAME`, sysCode, tenantId, sysCode).QueryRows(&result)
	return result, err
}

func UpdateTenantPermission(sysCode string, perIdStr string, perMenu string, tenId int) (err error) {
	err = DeleteTenatPermssion(sysCode, tenId)
	if err != nil {
		return err
	}
	err = InsertTenantPermission(perIdStr, tenId)
	if err != nil {
		return err
	}
	err = UpdateTenatMenuText(sysCode, perMenu, tenId)
	if err != nil {
		return err
	}
	return
}

//新增套餐已经勾选的信息
func InsertTenantPermission(perIdStr string, tenId int) (err error) {
	arr := strings.Split(perIdStr, ",")
	var param string
	for _, x := range arr {
		param += "'" + x + "',"
	}
	length := len(param) - 1
	params := param[0:length]
	o := orm.NewOrm()
	var sql = `INSERT INTO permission (NAME,tenantId,DisplayName,SysCode,MenuCode,CreationTime,IsMenu
		) SELECT NAME,?,DisplayName,SysCode,MenuCode,?,IsMenu FROM permission t1
		WHERE t1.TenantId = 0 AND IsMenu = 1 AND t1.Name IN (` + params + `)`
	_, err = o.Raw(sql, tenId, time.Now()).Exec()
	return err
}

//删除原先配置的权限信息
func DeleteTenatPermssion(sysCode string, tenId int) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("DELETE FROM	permission WHERE TenantId = ? and SysCode=? and RoleId=0 ", tenId, sysCode).Exec()
	return err
}
