package models

import (
	out "demo/outmodels"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Permission struct {
	Id                     int       `orm:"column(Id);auto"`
	CreationTime           time.Time `orm:"column(CreationTime);type(datetime)"`
	CreatorUserId          int64     `orm:"column(CreatorUserId);null"`
	Discriminator          string    `orm:"column(Discriminator);size(300)"`
	Name                   string    `orm:"column(Name);size(128)"`
	TenantId               int       `orm:"column(TenantId);null"`
	RoleId                 int       `orm:"column(RoleId);null"`
	UserId                 int64     `orm:"column(UserId);null"`
	DisplayName            string    `orm:"column(DisplayName);size(50)"`
	SysCode                int       `orm:"column(SysCode);0"`
	DeletionTime           time.Time `orm:"column(DeletionTime);type(datetime)"`
	DeleterUserId          int64     `orm:"column(DeleterUserId);null"`
	LastModificationTime   time.Time `orm:"column(LastModificationTime);type(datetime)"`
	LastModificationUserId int64     `orm:"column(LastModificationUserId);null"`
	IsDeleted              int       `orm:"column(IsDeleted);0"`
	MenuCode               int       `orm:"column(MenuCode);null"`
	IsMenu                 int       `orm:"column(IsMenu);0"`
}

func (t *Permission) TableName() string {
	return "permission"
}

func init() {
	orm.RegisterModel(new(Permission))
}

// AddPermission insert a new Permission into database and returns
// last inserted Id on success.
func AddPermission(m *Permission) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPermissionById retrieves Permission by Id. Returns error if
// Id doesn't exist
func GetPermissionById(id int) (v *Permission, err error) {
	o := orm.NewOrm()
	v = &Permission{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPermission retrieves all Permission matches certain condition. Returns empty list if
// no records exist
func GetAllPermission(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Permission
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdatePermission updates Permission by Id and returns error if
// the record to be updated doesn't exist
func UpdatePermissionById(m *Permission) (err error) {
	o := orm.NewOrm()
	v := Permission{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

//获取筛选条件下的数据总量
func GetTotalPermission(query map[string]string) (total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Permission))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	total, err = qs.Count()
	return total, err
}

// DeletePermission deletes Permission by Id and returns error if
// the record to be deleted doesn't exist
func DeletePermission(id int) (err error) {
	o := orm.NewOrm()
	v := Permission{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		_, err = o.Raw("update permission set IsDeleted=1 ,DeletionTime = ?  where Id= ? ", time.Now(), id).Exec()
	}
	return
}

func GetPermissionByUserAndPermission(userid int64, permissionName string) (p *Permission, err error) {
	o := orm.NewOrm()
	p = &Permission{}
	err = o.QueryTable("permission").Filter("userid", userid).Filter("name", permissionName).One(p)
	if err == orm.ErrMultiRows {
		// 多条的时候报错
		return nil, err
	}
	if err == orm.ErrNoRows {
		// 没有找到记录
		return nil, err
	}
	return p, nil
}

func GetPermissionByUser(userid int64, sysCode string) (permissions []Permission, num int64) {
	o := orm.NewOrm()
	num, _ = o.QueryTable("permission").Filter("UserId", userid).Filter("SysCode", sysCode).All(&permissions)
	return permissions, num
}

// 获取列表的信息
func GetPermissionList(menuName string, sysName string, offset int64, limit int64) (result []out.MenuInfo, err error) {
	o := orm.NewOrm()
	var sql = "SELECT t1.DisplayName menu_name,t2.SysName sys_name,t1.MenuText menu_text,t1.id FROM permission t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode WHERE t1.isMenu=0 "
	conditions := []string{}
	if menuName != "" {
		conditions = append(conditions, " t1.DisplayName like '%"+menuName+"%'")
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
func CountPermissionInfo(menuName string, sysName string) (total int64) {
	o := orm.NewOrm()
	conditions := []string{}
	var sql = "SELECT count(0) total FROM permission t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode WHERE t1.isMenu=0"
	if menuName != "" {
		conditions = append(conditions, " t1.DisplayName like '%"+menuName+"%'")
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

func GetPerInfoBySysCode(SysCode string) (result []out.PermissionCheckInfo, err error) {
	o := orm.NewOrm()
	_, err = o.Raw("SELECT t1. NAME name ,t1.DisplayName display_name,GROUP_CONCAT(t2.NAME) code ,GROUP_CONCAT(t2.DisplayName) code_name FROM permission t1 LEFT JOIN permission t2 ON t1.MenuCode = t2.MenuCode WHERE t1.SysCode = t2.SysCode AND t1.IsMenu = 0 AND t1.SysCode = ? AND t2.IsMenu = 1 group by t1.Name order by t1.MenuCode,t1.Id asc", SysCode).QueryRows(&result)
	return result, err
}

//根据租户id和系统编号 修改租户下的权限信息
func UpdateTenantPerission(sysCode string, tenId string) {

}
