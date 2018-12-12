package controllers

import (
	"demo/models"
	out "demo/outmodels"
	tool "demo/tools"
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// PermissionController operations for Permission
type PermissionController struct {
	beego.Controller
}

// URLMapping ...
func (c *PermissionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("getPerInfoBySysCode", c.GetPerInfoBySysCode)
}

// Post ...
// @Title Post
// @Description create Permission
// @Param	body		body 	models.Permission	true		"body for Permission content"
// @Success 201 {int} models.Permission
// @Failure 403 body is empty
// @router / [post]
func (c *PermissionController) Post() {
	result := &out.OperResult{}
	var v models.Permission
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.CreationTime = time.Now()

		if _, err := models.AddPermission(&v); err == nil {
			result.Result = 1
			c.Data["json"] = result
		} else {
			result.Result = 0
			result.Message = err.Error()
			c.Data["json"] = result
		}
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Permission by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Permission
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PermissionController) GetOne() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPermissionById(id)
	if err != nil {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	} else {
		result.Result = 1
		result.Data = v
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Permission
// @Param   menuName query   string false
// @Param   sysName query   string	false
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Permission
// @Failure 403
// @router / [get]
func (c *PermissionController) GetAll() {
	var menuName string
	var sysName string
	var limit int64 = 10
	var offset int64
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// menuName
	if v := c.GetString("menuName"); v != "" {
		menuName = v
	}
	// sysName
	if v := c.GetString("sysName"); v != "" {
		sysName = v
	}
	result := &out.OperResult{}
	data, err := models.GetPermissionList(menuName, sysName, offset, limit)
	total := models.CountPermissionInfo(menuName, sysName)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		result.Result = 1
		var page = make(map[string]interface{})
		page["list"] = data
		var ListQuery = make(map[string]int64)
		ListQuery["limit"] = limit
		ListQuery["page"] = offset
		page["listQuery"] = ListQuery
		page["total"] = total
		result.Data = page
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Permission
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Permission	true		"body for Permission content"
// @Success 200 {object} models.Permission
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PermissionController) Put() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Permission{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.LastModificationTime = time.Now()
		if err := models.UpdatePermissionById(&v); err == nil {
			result.Result = 1
			c.Data["json"] = result
		} else {
			result.Result = 0
			result.Message = err.Error()
			c.Data["json"] = result
		}
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Permission
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PermissionController) Delete() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeletePermission(id); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// 根据系统编号获取对应的权限
// @router /getPerInfoBySysCode/:sysCode [get]
func (c *PermissionController) GetPerInfoBySysCode() {
	result := &out.OperResult{}
	sysCode := c.Ctx.Input.Param(":sysCode")
	if data, err := models.GetPerInfoBySysCode(sysCode); len(data) > 0 {
		permissionList := tool.ParsePermissionDataForCheckbox(data)
		result.Result = 1
		result.Data = permissionList
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// 根据套餐编号和系统编号获取对应的权限
// @router /getPerInfoBySysCodeUpdate [get]
func (c *PermissionController) GetPerInfoBySysCodeUpdate() {
	result := &out.OperResult{}
	sysCode := c.GetString("sysCode")
	setMealCode := c.GetString("setMealCode")
	if data, err := models.GetPermBySetMealCode(setMealCode, sysCode); err == nil {
		permissionList := tool.ParsePermissionDataForCheckboxUpdate(data)
		result.Result = 1
		result.Data = permissionList
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}
