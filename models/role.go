package models

import (
	input "demo/inputmodels"
	out "demo/outmodels"
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Role struct {
	ID                   int       `orm:"column(Id);auto"`
	RoleCode             string    `orm:"column(RoleCode);size(20)"`
	RoleName             string    `orm:"column(RoleName);size(45)"`
	TenantID             int64     `orm:"column(TenantId);null"`
	SysCode              string    `orm:"column(SysCode);size(20)"`
	IsValid              int       `orm:"column(IsValid);0"`
	IsDeleted            int8      `orm:"column(IsDeleted);0"`
	AuthText             string    `orm:"column(AuthText);size(1024)"`
	CreationTime         time.Time `orm:"column(CreationTime);type(datetime)"`
	CreatorUserID        int64     `orm:"column(CreatorUserId);null"`
	DeleterUserID        int64     `orm:"column(DeleterUserId);null"`
	DeletionTime         time.Time `orm:"column(DeletionTime);type(datetime);null"`
	LastModificationTime time.Time `orm:"column(LastModificationTime);type(datetime);null"`
	LastModifierUserID   int64     `orm:"column(LastModifierUserId);null"`
}

func (t *Role) TableName() string {
	return "role"
}

func init() {
	orm.RegisterModel(new(Role))
}

// AddRole insert a new Role into database and returns
// last inserted Id on success.
func AddRole(m *input.RoleInput, userID int64, tenantID int64) (id int64, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	o.Raw("select IFNULL(MAX(RoleCode),'300000')+1  sysCode from role").Values(&maps)
	sysCode := m.SysCode
	nowTime := time.Now()
	o.Begin()
	role := new(Role)
	role.TenantID = tenantID
	role.CreationTime = nowTime
	role.AuthText = m.PerName
	role.RoleName = m.RoleName
	role.SysCode = sysCode
	role.CreatorUserID = userID
	role.RoleCode = maps[0]["sysCode"].(string)
	id, err = o.Insert(role)
	if err != nil {
		o.Rollback()
	}
	arr := strings.Split(m.PerId, ",")
	var param string
	for _, x := range arr {
		param += "'" + x + "',"
	}
	length := len(param) - 1
	params := param[0:length]
	var sql = `INSERT INTO permission (NAME,tenantId,DisplayName,SysCode,MenuCode,CreationTime,IsMenu,UserId,RoleId
		) SELECT NAME,?,DisplayName,?,MenuCode,?,IsMenu,?,? FROM permission t1
		WHERE t1.TenantId = 0 AND IsMenu = 1 AND t1.Name IN (` + params + `)`
	_, err = o.Raw(sql, tenantID, sysCode, nowTime, userID, role.ID).Exec()
	if err != nil {
		o.Rollback()
	}
	o.Commit()
	return
}

// GetRoleById retrieves Role by Id. Returns error if
// Id doesn't exist
func GetRoleById(id int) (v *Role, err error) {
	o := orm.NewOrm()
	v = &Role{ID: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//获取筛选条件下的数据总量
func GetTotalRole(query map[string]string) (total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
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

// GetAllRole retrieves all Role matches certain condition. Returns empty list if
// no records exist
func GetAllRole(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Role))
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

	var l []Role
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

// UpdateRole updates Role by Id and returns error if
// the record to be updated doesn't exist
func UpdateRoleById(m *Role) (err error) {
	o := orm.NewOrm()
	v := Role{ID: m.ID}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.SysCode = v.SysCode
		m.TenantID = v.TenantID
		m.CreationTime = v.CreationTime
		m.CreatorUserID = v.CreatorUserID
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 获取角色列表的信息
func GetRoleList(roleName string, sysName string, offset int64, limit int64, tenantID int64) (result []out.RoleInfo, err error) {
	o := orm.NewOrm()
	var sql = `SELECT  t1.RoleName role_name,  t1.RoleCode role_code, t2.SysName sys_name, t1.isValid is_valid, t1.AuthText auth_text,  t1.ID id
	FROM  role t1 LEFT JOIN   application t2 ON t1.SysCode = t2.SysCode WHERE  t1.TenantId = ? AND t1.isDeleted = 0 and t2.isDeleted=0 and t2.isValid = 0 `
	conditions := []string{}
	if roleName != "" {
		conditions = append(conditions, " t1.roleName like '%"+roleName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t2.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " and " + strings.Join(conditions, " and ")
	}
	sql = sql + " limit " + strconv.FormatInt(limit, 10) + "  offset " + strconv.FormatInt(offset, 10)
	_, err = o.Raw(sql, tenantID).QueryRows(&result)
	return result, err
}

// 统计查询条件的数量
func CountRoleInfo(roleName string, sysName string, tenantID int64) (total int64) {
	o := orm.NewOrm()
	conditions := []string{}
	var sql = "SELECT count(0) total FROM  role t1 LEFT JOIN   application t2 ON t1.SysCode = t2.SysCode WHERE  t1.TenantId = ? AND t1.isDeleted = 0 and t2.isDeleted=0 and t2.isValid = 0  "
	if roleName != "" {
		conditions = append(conditions, " t1.RoleName like '%"+roleName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t2.SysName  like '%"+sysName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " and " + strings.Join(conditions, " and ")
	}
	var maps []orm.Params
	o.Raw(sql, tenantID).Values(&maps)
	total, _ = strconv.ParseInt(maps[0]["total"].(string), 10, 64)
	return total
}

//UpdateValidStatus 更新isValid数据状态
func UpdateValidStatus(id int, isValid int64, userID int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE role SET IsValid = ?, LastModificationTime =? , LastModifierUserId =? WHERE Id =?", isValid, time.Now(), userID, id).Exec()
	return err
}

//DeleteRole 删除角色信息
func DeleteRole(ids string, userID int64) (err error) {
	arr := strings.Split(ids, ",")
	var param string
	for _, x := range arr {
		param += x + ","
	}
	length := len(param) - 1
	params := param[0:length]
	var sql = "update role set IsDeleted=1 , DeletionTime = ? ,DeleterUserId = ? where Id in ( " + params + ")"
	o := orm.NewOrm()
	_, err = o.Raw(sql, time.Now(), userID).Exec()
	return
}
