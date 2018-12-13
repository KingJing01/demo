package models

import (
	out "demo/outmodels"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Tenant struct {
	Id                   int       `orm:"column(Id);auto"`
	TenancyName          string    `orm:"column(TenancyName);size(64)"`
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
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
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
