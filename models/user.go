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
	"github.com/astaxie/beego/validation"
)

type User struct {
	Id                     int64     `orm:"column(Id);auto"`
	CreationTime           time.Time `orm:"column(CreationTime);type(datetime)"`
	CreatorUserId          int64     `orm:"column(CreatorUserId);null"`
	DeleterUserId          int64     `orm:"column(DeleterUserId);null"`
	DeletionTime           time.Time `orm:"column(DeletionTime);type(datetime);null"`
	EmailAddress           string    `orm:"column(EmailAddress);size(256)"`
	EmailConfirmationCode  string    `orm:"column(EmailConfirmationCode);size(328);null"`
	IsDeleted              int8      `orm:"column(IsDeleted)"`
	IsEmailConfirmed       int8      `orm:"column(IsEmailConfirmed)"`
	IsPhoneNumberConfirmed int8      `orm:"column(IsPhoneNumberConfirmed)"`
	LastModificationTime   time.Time `orm:"column(LastModificationTime);type(datetime);null"`
	LastModifierUserId     int64     `orm:"column(LastModifierUserId);null"`
	Name                   string    `orm:"column(Name);size(32)"`
	UserName               string    `orm:"column(UserName);size(32)"`
	PasswordResetCode      string    `orm:"column(PasswordResetCode);size(328);null"`
	PhoneNumber            string    `orm:"column(PhoneNumber);size(32);null"`
	TenantId               int       `orm:"column(TenantId);null"`
	SysCode                string    `orm:"column(SysCode)"`
	SsoID                  int       `orm:"column(SsoId)"`
	UserUrl                string    `orm:"column(UserUrl)"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetUserById retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserById(id int64) (v *User, err error) {
	o := orm.NewOrm()
	v = &User{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllUser retrieves all User matches certain condition. Returns empty list if
// no records exist
func GetAllUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))
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

	var l []User
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

// UpdateUser updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserById(m *User) (err error) {
	o := orm.NewOrm()
	v := User{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteUser deletes User by Id and returns error if
// the record to be deleted doesn't exist
func DeleteUser(id int64) (err error) {
	o := orm.NewOrm()
	v := User{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&User{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}

//根据用户名、密码查询
func LoginCheck(username string, password string, SysCode string) (result bool, user User, err error) {
	valid := validation.Validation{}
	resultMobile := valid.Mobile(username, "username")
	o := orm.NewOrm()
	u := &User{}
	result = true
	//登录名格式分析  手机号码直接 ssoUser验证 其他的使用user--->sso关联
	if resultMobile.Ok {
		err = o.Raw("select t2.*, t1.Phone  SsoPhone from ssouser t1 left join user t2 on t1.id = t2.SsoId where t2.SysCode=? and t1.Phone=? and t1.Passwd=? ", SysCode, username, password).QueryRow(&u)
	} else {
		err = o.Raw("select t2.*,t1.Phone SsoPhone from ssouser t1 left join user t2 on t1.id = t2.SsoId where t2.SysCode=? and t2.UserName=? and t1.Passwd=? ", SysCode, username, password).QueryRow(&u)
	}
	user = *u
	// 判断是否有错误的返回
	if err != nil || int(user.Id) == 0 {
		result = false
		return result, user, err
	}
	return true, user, nil
}

func RegistUser(loginInfo *input.LoginInfo, SysCode string) (ssoId int, err error) {
	o := orm.NewOrm()
	o.Begin()
	ssoUser := new(SsoUser)
	ssoUser.Phone = loginInfo.UserName
	ssoUser.Passwd = loginInfo.Password
	_, err = o.Insert(ssoUser)
	if err != nil {
		o.Rollback()
	}
	ssoId = ssoUser.Id
	user := new(User)
	user.Name = loginInfo.UserName
	user.PhoneNumber = loginInfo.UserName
	user.SsoID = ssoUser.Id
	user.SysCode = SysCode
	user.CreationTime = time.Now()
	_, err = o.Insert(user)
	if err != nil {
		o.Rollback()
	}
	o.Commit()
	return
}

func PasswdUpdate(info *input.LoginInfo, SysCode string) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("update ssouser set Passwd=? where Phone=? or Email=?", info.Password, info.UserName, info.UserName).Exec()
	return err
}

// 获取角色列表的信息
func GetUserList(roleName string, sysName string, offset int64, limit int64, tenantID int64) (result []out.RoleInfo, err error) {
	o := orm.NewOrm()
	var sql = `SELECT 	t2.EmailAddress,t2.PhoneNumber,t2.UserName,t2.Id,t3.SysName,t4.RoleName,
	t4.AuthText,t4.IsValid FROM	USER t2 LEFT JOIN  userrole t1  ON t1.UserId = t2.Id
	LEFT JOIN application t3 ON t2.SysCode = t3.SysCode LEFT JOIN role t4 ON t1.RoleId = t4.Id
	where t2.TenantId=? `
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
func CountUserInfo(roleName string, sysName string, tenantID int64) (total int64) {
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
