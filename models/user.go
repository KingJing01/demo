package models

import (
	"crypto/sha256"
	input "demo/inputmodels"
	out "demo/outmodels"
	"encoding/hex"
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
	Password               string    `orm:"column(Password);size(64);null"`
	TenantId               int64     `orm:"column(TenantId);null"`
	SysCode                string    `orm:"column(SysCode)"`
	SsoID                  int64     `orm:"column(SsoId)"`
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
func AddUser(m *User, roleIds []string, sysCodes []string, tenantID int64, userID int64) (tmsUser out.TMSUser, err error) {
	o := orm.NewOrm()
	o.Begin()
	//新增一个ssouer
	ssoUser := SsoUser{}
	ssoUser.Phone = m.PhoneNumber
	ssoUser.Email = m.EmailAddress
	ssoID, err := o.Insert(&ssoUser)
	if err != nil {
		//回滚
		o.Rollback()
		return tmsUser, err
	}
	userrole := UserRole{}
	userrole.CreationTime = time.Now()
	userrole.TenantId = tenantID
	m.SsoID = ssoID
	m.TenantId = tenantID
	m.Password = GetDefaultPassword("")
	m.CreatorUserId = userID
	sysCodeStr := ""
	for j, t := range roleIds {

		for k, z := range sysCodes {
			if j == k {
				sysCodeStr += z + ","
				m.SysCode = z
				roleID, _ := strconv.Atoi(t)
				userrole.SysCode = z
				userrole.RoleId = roleID
				_, err = o.Insert(m)
				if err != nil {
					o.Rollback()
					return tmsUser, err
				}
				//角色权限插入
				var sql = `INSERT INTO permission (NAME,tenantId,DisplayName,SysCode,MenuCode,CreationTime,IsMenu,RoleId,UserId
						) SELECT NAME,tenantId,DisplayName,SysCode,MenuCode,?,IsMenu,RoleId,? FROM permission t1
						WHERE t1.TenantId = ? AND IsMenu = 1 AND t1.RoleId =?`
				_, err = o.Raw(sql, time.Now(), m.Id, tenantID, t).Exec()
				if err != nil {
					o.Rollback()
					return tmsUser, err
				}
				userrole.UserId = m.Id
				userrole.CreatorUserId = userID
				_, err = o.Insert(&userrole)
				if err != nil {
					o.Rollback()
					return tmsUser, err
				}
				m.Id = 0
				userrole.Id = 0
			}
		}
	}
	v := Tenant{Id: tenantID}
	o.Read(&v)
	tmsUser.SsoUID = strconv.FormatInt(ssoID, 10)
	tmsUser.UserCode = m.UserName
	tmsUser.Email = m.EmailAddress
	tmsUser.Mobile = m.PhoneNumber
	tmsUser.CompanyID = strconv.FormatInt(tenantID, 10)
	tmsUser.CompanyName = v.TenantName
	tmsUser.SysID = sysCodeStr
	tmsUser.Contact = m.Name
	tmsUser.ShortCompanyName = v.ShortName
	tmsUser.IsAdmin = "0"
	respCode, err := out.SendUserInfoToTms(tmsUser)
	fmt.Println("################接口返回的标记值################ ", respCode)
	if respCode != 200 {
		o.Rollback()
		return tmsUser, err
	}
	//事务提交
	o.Commit()
	return tmsUser, err
}

// GetUserByID retrieves User by Id. Returns error if
// Id doesn't exist
func GetUserByID(id int64) (v *out.UserInfo, err error) {
	o := orm.NewOrm()
	err = o.Raw(`select t1.Id id ,t1.UserName user_name,t1.Name name,t1.PhoneNumber phone_number,t1.EmailAddress email_address,t1.SysCode sys_code,t3.RoleName role_name,t2.RoleId role_code from user t1 left join userrole
	 t2 on t1.Id = t2.UserId left join role t3 on t2.RoleId = t3.id where t1.id= ?`, id).QueryRow(&v)
	if err != nil {
		return nil, err
	}
	return v, err
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

// UpdateUserByID updates User by Id and returns error if
// the record to be updated doesn't exist
func UpdateUserByID(m *User, roleCode string, userID int64) (err error) {
	o := orm.NewOrm()
	v := &User{Id: m.Id}
	if err = o.Read(v); err == nil {
		m.CreationTime = v.CreationTime
		m.CreatorUserId = v.CreatorUserId
		m.TenantId = v.TenantId
		m.SsoID = v.SsoID
		m.LastModificationTime = time.Now()
		m.LastModifierUserId = userID
		_, err = o.Update(m)
		_, err = o.Raw("update userrole set RoleId=?  where  SysCode=? and UserId=?", roleCode, m.SysCode, m.Id).Exec()
		if err != nil {
			o.Rollback()
		}
	}
	return
}

// 根据用户名查询当前人是否存在
func LoginValidUser(username string) (total int64, err error) {
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw("SELECT count(0) total FROM USER t1  WHERE t1.IsDeleted = 0 and (t1.PhoneNumber=? or t1.UserName=? or t1.EmailAddress=?)", username, username, username).Values(&maps)
	total, _ = strconv.ParseInt(maps[0]["total"].(string), 10, 64)
	return total, err
}

//根据用户名、密码查询系统
func LoginValidSys(username string) (sysCode string, err error) {
	o := orm.NewOrm()
	var maps []orm.Params
	_, err = o.Raw("select GROUP_CONCAT(SysCode) sysCode from user where  IsDeleted=0 and (UserName = ? or PhoneNumber =? or EmailAddress=?)", username, username, username).Values(&maps)
	if maps[0]["sysCode"] == nil {
		return "", err
	}
	sysCode = maps[0]["sysCode"].(string)
	return sysCode, err
}

//根据用户名、密码，系统查询
func LoginCheck(username string, password string, SysCode string) (result bool, user User, err error) {
	valid := validation.Validation{}
	resultEmail := valid.Email(username, "username")
	password = GetDefaultPassword(password)
	o := orm.NewOrm()
	u := &User{}
	result = true
	//登录名格式分析  邮箱 登陆  用户名和手机号
	if resultEmail.Ok {
		err = o.Raw("SELECT t1.* FROM USER t1  WHERE t1.IsDeleted = 0 and t1.SysCode=? and t1.EmailAddress=? and t1.Password=?", SysCode, username, password).QueryRow(&u)
	} else {
		err = o.Raw("SELECT t1.* FROM USER t1  WHERE t1.IsDeleted = 0 and t1.SysCode=? and (t1.PhoneNumber=? or t1.UserName=?) and t1.Password=?", SysCode, username, username, password).QueryRow(&u)
	}
	user = *u
	// 判断是否有错误的返回
	if err != nil || int(user.Id) == 0 {
		result = false
		return result, user, err
	}
	return true, user, nil
}

//根据用户名、密码查询
func AuthLoginCheck(username string, password string) (result bool, user User, err error) {
	o := orm.NewOrm()
	password = GetDefaultPassword(password)
	err = o.Raw(`SELECT t1.* FROM USER t1 LEFT JOIN userrole t2 on t1.Id =t2.UserId WHERE t1.IsDeleted = 0
	AND t2.RoleId in (0,1) AND (t1.PhoneNumber = ? OR t1.UserName = ? OR t1.EmailAddress = ?)
	AND t1. PASSWORD = ?`, username, username, username, password).QueryRow(&user)
	// 判断是否有错误的返回
	if err != nil {
		return false, user, err
	}
	return true, user, nil
}

func RegistUser(loginInfo *input.LoginInfo, SysCode string) (ssoId int64, err error) {
	o := orm.NewOrm()
	o.Begin()
	ssoUser := new(SsoUser)
	ssoUser.Phone = loginInfo.UserName
	_, err = o.Insert(ssoUser)
	if err != nil {
		o.Rollback()
	}
	ssoId = ssoUser.Id
	user := new(User)
	user.Password = GetDefaultPassword(loginInfo.Password)
	user.Name = loginInfo.UserName
	user.PhoneNumber = loginInfo.UserName
	user.SsoID = ssoId
	user.SysCode = SysCode
	user.TenantId = 1
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
func GetUserList(roleName string, sysName string, userName string, offset int64, limit int64, tenantID int64) (result []out.UserInfo, err error) {
	o := orm.NewOrm()
	var sql = `SELECT t2.EmailAddress email_address ,t2.PhoneNumber phone_number,t2.UserName user_name,t2.Id id,t3.SysName sys_name,t4.RoleName role_name,
	t4.AuthText auth_text,t4.IsValid is_valid FROM	USER t2 LEFT JOIN  userrole t1  ON t1.UserId = t2.Id
	LEFT JOIN application t3 ON t2.SysCode = t3.SysCode LEFT JOIN role t4 ON t1.RoleId = t4.Id
	where t2.TenantId=?  and t2.IsDeleted=0`
	conditions := []string{}
	if roleName != "" {
		conditions = append(conditions, " t4.roleName like '%"+roleName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t3.SysName  like '%"+sysName+"%'")
	}
	if userName != "" {
		conditions = append(conditions, " t2.PhoneNumber  like '%"+userName+"%' or t2.EmailAddress like '%"+userName+"%' or t2.UserName like '%"+userName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " and " + strings.Join(conditions, " and ")
	}
	sql = sql + " limit " + strconv.FormatInt(limit, 10) + "  offset " + strconv.FormatInt(offset, 10)
	_, err = o.Raw(sql, tenantID).QueryRows(&result)
	return result, err
}

// 统计查询条件的数量
func CountUserInfo(roleName string, sysName string, userName string, tenantID int64) (total int64) {
	o := orm.NewOrm()
	conditions := []string{}
	var sql = "SELECT count(0) total FROM	USER t2 LEFT JOIN  userrole t1  ON t1.UserId = t2.Id LEFT JOIN application t3 ON t2.SysCode = t3.SysCode LEFT JOIN role t4 ON t1.RoleId = t4.Id where t2.TenantId=? and t2.IsDeleted=0"
	if roleName != "" {
		conditions = append(conditions, " t4.RoleName like '%"+roleName+"%'")
	}
	if sysName != "" {
		conditions = append(conditions, " t3.SysName  like '%"+sysName+"%'")
	}
	if userName != "" {
		conditions = append(conditions, " t2.PhoneNumber  like '%"+userName+"%' or t2.EmailAddress like '%"+userName+"%' or t2.UserName like '%"+userName+"%'")
	}
	if len(conditions) > 0 {
		sql = sql + " and " + strings.Join(conditions, " and ")
	}
	var maps []orm.Params
	o.Raw(sql, tenantID).Values(&maps)
	total, _ = strconv.ParseInt(maps[0]["total"].(string), 10, 64)
	return total
}

//UpdateUserValidStatus 更新isValid数据状态
func UpdateUserValidStatus(id int, isValid int64, userID int64) (err error) {
	o := orm.NewOrm()
	_, err = o.Raw("UPDATE user SET IsValid = ?, LastModificationTime =? , LastModifierUserId =? WHERE Id =?", isValid, time.Now(), userID, id).Exec()
	return err
}

//DeleteUser 删除角色信息
func DeleteUser(ids string, userID int64) (err error) {
	arr := strings.Split(ids, ",")
	var param string
	for _, x := range arr {
		param += x + ","
	}
	length := len(param) - 1
	params := param[0:length]
	var sql = "update user set IsDeleted=1 , DeletionTime = ? ,DeleterUserId = ? where Id in ( " + params + ")"
	o := orm.NewOrm()
	_, err = o.Raw(sql, time.Now(), userID).Exec()
	return
}

//GetDefaultPassword 参数为"" 生成默认的加密密码 sha256 不为空对密码进行加密
func GetDefaultPassword(val string) (passwd string) {
	if val == "" {
		val = "123456"
	}
	hash := sha256.New()
	hash.Write([]byte(val))
	md := hash.Sum(nil)
	passwd = hex.EncodeToString(md)
	return passwd
}
