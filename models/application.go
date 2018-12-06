package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Application struct {
	Id                     int       `orm:"column(Id);auto;pk"`
	SysCode                string    `orm:"column(SysCode);size(20)"`
	SysName                string    `orm:"column(SysName);size(45)"`
	SysUrl                 string    `orm:"column(SysUrl);size(255)"`
	CreationTime           time.Time `orm:"column(CreationTime);type(datetime);time"`
	CreatorUserId          int64     `orm:"column(CreatorUserId);null"`
	LastModificationTime   time.Time `orm:"column(LastModificationTime);type(datetime);time"`
	LastModificationUserId int64     `orm:"column(LastModificationUserId);null"`
	IsDeleted              int       `orm:"column(IsDeleted);0"`
	DeletionTime           time.Time `orm:"column(DeletionTime);type(datetime);time"`
	DeletionUserId         int64     `orm:"column(DeletionUserId);null"`
	IsValid                int       `orm:"column(IsValid);0"`
}

func (t *Application) TableName() string {
	return "application"
}

func init() {
	orm.RegisterModel(new(Application))
}

// AddApplication insert a new Application into database and returns
// last inserted Id on success.
func AddApplication(m *Application) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetApplicationById retrieves Application by Id. Returns error if
// Id doesn't exist
func GetApplicationById(id int) (v *Application, err error) {
	o := orm.NewOrm()
	v = &Application{Id: id}
	v.IsDeleted = 0
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//获取筛选条件下的数据总量
func GetTotalApplication(query map[string]string) (total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Application))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			k = k + "__icontains"
			qs = qs.Filter(k, v)
		}
	}
	total, err = qs.Count()
	return total, err
}

// GetAllApplication retrieves all Application matches certain condition. Returns empty list if
// no records exist
func GetAllApplication(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Application))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			k = k + "__icontains"
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

	var l []Application
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

// UpdateApplication updates Application by Id and returns error if
// the record to be updated doesn't exist
func UpdateApplicationById(m *Application) (err error) {
	o := orm.NewOrm()
	v := Application{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		m.CreationTime = v.CreationTime
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// 逻辑删除应用
// the record to be deleted doesn't exist
func DeleteApplication(id int) (err error) {
	o := orm.NewOrm()
	v := Application{Id: id}
	if err = o.Read(&v); err == nil {
		_, err = o.Raw("update application set IsDeleted=1 , DeletionTime = ?  where Id= ?  ", time.Now(), id).Exec()
	}
	return
}

// 系统名验证重复
func CheckRepeat(sysName string) (total int64, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Application)).Filter("SysName", sysName)
	return qs.Count()
}

func GenerateSysCode() (SysCode string) {
	var maps []orm.Params
	o := orm.NewOrm()
	o.Raw("select IFNULL(MAX(SysCode),'100000')+1  sysCode from application").Values(&maps)
	return maps[0]["sysCode"].(string)
}
