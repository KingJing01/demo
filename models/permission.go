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
	TenantId               int       `orm:"column(TenantId);0"`
	RoleId                 int       `orm:"column(RoleId);0"`
	UserId                 int64     `orm:"column(UserId);0"`
	DisplayName            string    `orm:"column(DisplayName);size(50)"`
	SysCode                int       `orm:"column(SysCode);"`
	DeletionTime           time.Time `orm:"column(DeletionTime);type(datetime)"`
	DeleterUserId          int64     `orm:"column(DeleterUserId);null"`
	LastModificationTime   time.Time `orm:"column(LastModificationTime);type(datetime)"`
	LastModificationUserId int64     `orm:"column(LastModificationUserId);null"`
	IsDeleted              int       `orm:"column(IsDeleted);0"`
	MenuCode               int       `orm:"column(MenuCode);null"`
	MenuText               string    `orm:"column(MenuText);null"`
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
func AddPermission(m map[string]interface{}, userID int64) (id int64, err error) {
	sysCodeStr := m["SysCode"].(string)
	sysCode, _ := strconv.Atoi(sysCodeStr)
	menuCode := GenerMenuCode()
	o := orm.NewOrm()
	//批量插入数据组装菜单数据和菜单的权限
	var permissionList []Permission
	var permiss Permission
	permiss.CreationTime = time.Now()
	permiss.SysCode = sysCode
	permiss.MenuCode = menuCode
	permiss.CreatorUserId = userID
	arr := m["PerData"].([]interface{})
	var menuText string
	for _, per := range arr {
		value := per.(map[string]interface{})
		menu := value["DisplayName"].(string)
		menuText += "," + menu
		permiss.DisplayName = menu
		permiss.Name = value["Name"].(string)
		permiss.IsMenu = 1
		permissionList = append(permissionList, permiss)
	}
	// 菜单数据的拼装
	permiss.DisplayName = m["DisplayName"].(string)
	permiss.Name = m["Name"].(string)
	permiss.IsMenu = 0
	rs := []rune(menuText)
	lth := len(rs)
	if lth > 0 {
		menuText = string(rs[1:lth])
		permiss.MenuText = menuText
	}
	permissionList = append(permissionList, permiss)
	id, err = o.InsertMulti(len(arr)+1, permissionList)
	return
}

// GetPermissionById retrieves Permission by Id. Returns error if
// Id doesn't exist
func GetPermissionById(id int) (v *out.PermissonInfo, err error) {
	o := orm.NewOrm()
	err = o.Raw("select t1.Id id ,DisplayName display_name,Name name ,MenuText menu_text,MenuCode menu_code,t2.SysName sys_name,t1.SysCode sys_code from permission t1 LEFT JOIN application t2 on t1.SysCode = t2.SysCode where t1.Id = ?", id).QueryRow(&v)
	var lists []out.PerInfo
	_, err = o.Raw("select DisplayName display_name,Name name from permission where MenuCode = ? and IsMenu=1 and TenantId = 0 ", v.MenuCode).QueryRows(&lists)
	v.PerData = lists
	return v, err
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

func GetPermissionByUser(userid int64, sysCode string) (perInfos []out.PerInfo) {
	o := orm.NewOrm()
	if sysCode != "" {
		o.Raw("select t2.childData display_name ,t1.Name name from permission t1 right join (select GROUP_CONCAT(Name) childData ,MenuCode from permission where UserId=?   group by MenuCode ) t2 on t1.MenuCode = t2.MenuCode where t1.IsMenu=0 and t1.SysCode=?", userid, sysCode).QueryRows(&perInfos)
	}
	return perInfos
}

// 获取列表的信息
func GetPermissionList(menuName string, sysName string, offset int64, limit int64) (result []out.MenuInfo, err error) {
	o := orm.NewOrm()
	var sql = "SELECT t1.DisplayName menu_name,t2.SysName sys_name,t2.SysCode sys_code,t1.MenuText menu_text,t1.id FROM permission t1 LEFT JOIN application t2 ON t1.SysCode = t2.SysCode WHERE t1.isMenu=0 "
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

//GetPerInfoBySysCode 根据系统编号和租户Id(token获取)得到权限
func GetPerInfoBySysCode(SysCode string, TenantID int64) (result []out.PermissionCheckInfo, err error) {
	o := orm.NewOrm()
	_, err = o.Raw(`SELECT t1. NAME name ,t1.DisplayName display_name,GROUP_CONCAT(t2.NAME) code ,GROUP_CONCAT(t2.DisplayName) code_name FROM permission t1 
		LEFT JOIN permission t2 ON t1.MenuCode = t2.MenuCode WHERE t1.SysCode = t2.SysCode AND t1.IsMenu = 0 AND  t1.SysCode = ? AND t2.IsMenu = 1 and  t2.RoleId=0 and t2.TenantId=?
		group by t1.Name order by t1.MenuCode,t1.Id asc`, SysCode, TenantID).QueryRows(&result)
	return result, err
}

//GenerMenuCode 字段生成菜单号
func GenerMenuCode() (menuCode int) {
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw("select IFNULL(MAX(MenuCode),'2000')+1  menuCode from permission").Values(&maps)
	menuCode, _ = strconv.Atoi(maps[0]["menuCode"].(string))
	return menuCode
}

// 更新基本权限信息
func UpdatePermission(m map[string]interface{}, id int, userID int64) (err error) {
	o := orm.NewOrm()
	v := Permission{Id: id}
	o.Read(&v)
	var permissionList []Permission
	var permiss Permission
	permiss.CreationTime = time.Now()
	permiss.CreatorUserId = userID
	permiss.SysCode = v.SysCode
	arr := m["PerData"].([]interface{})
	var menuText string
	for _, per := range arr {
		value := per.(map[string]interface{})
		menu := value["DisplayName"].(string)
		menuText += "," + menu
		permiss.DisplayName = menu
		permiss.Name = value["Name"].(string)
		permiss.IsMenu = 1
		permiss.MenuCode = v.MenuCode
		permissionList = append(permissionList, permiss)
	}
	displayName := m["DisplayName"].(string)
	name := m["Name"].(string)
	rs := []rune(menuText)
	lth := len(rs)
	menuText = string(rs[1:lth])
	_, err = o.Raw("UPDATE permission set DisplayName=?,Name=?,LastModificationTime=?,MenuText=?,LastModificationUserId=? where Id=?", displayName, name, time.Now(), menuText, userID, id).Exec()
	_, err = o.Raw("DELETE FROM  permission where TenantId = 0 and IsMenu = 1  and  MenuCode=?  ", v.MenuCode).Exec()
	_, err = o.InsertMulti(len(arr), permissionList)
	return err
}

//GetPerInfoByRoleCode ...
func GetPerInfoByRoleCode(roleID string, sysCode string, TenantID int64, userID int64) (result []out.PermissionCheckInfo, err error) {
	o := orm.NewOrm()
	_, err = o.Raw(`SELECT t3.Name name ,t3.DisplayName display_name,GROUP_CONCAT(t1.NAME) code,GROUP_CONCAT(t1.DisplayName) code_name,
	GROUP_CONCAT(CASE	WHEN t2.NAME IS NULL THEN	0 ELSE 1 END ) flag
FROM ( SELECT MenuCode,NAME,DisplayName FROM permission WHERE TenantId = ? AND RoleId = 0 AND SysCode = ?) t1
LEFT JOIN (SELECT	NAME,	DisplayName FROM permission WHERE TenantId = ? AND UserId = ? AND RoleId = ? AND SysCode = ?
) t2 ON t2.NAME = t1.NAME  LEFT JOIN (
SELECT MenuCode,NAME,DisplayName FROM permission WHERE IsMenu=0) t3 on t3.MenuCode = t1.MenuCode group by t1.MenuCode`, TenantID, sysCode, TenantID, userID, roleID, sysCode).QueryRows(&result)
	return result, err
}
