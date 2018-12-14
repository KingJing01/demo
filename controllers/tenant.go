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

// TenantController operations for Tenant
type TenantController struct {
	beego.Controller
}

// URLMapping ...
func (c *TenantController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetTenantPermission", c.GetTenantPermission)
}

// Post ...
// @Title Post
// @Description create Tenant
// @Param	body		body 	models.Tenant	true		"body for Tenant content"
// @Success 201 {int} models.Tenant
// @Failure 403 body is empty
// @router / [post]
func (c *TenantController) Post() {
	result := &out.OperResult{}
	var v models.Tenant
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.CreationTime = time.Now()
		if _, err := models.AddTenant(&v); err == nil {
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
// @Description get Tenant by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TenantController) GetOne() {
	result := &out.OperResult{}
	//sysCode := c.GetString("sysCode")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTenantById(id)
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
// @Description get Tenant
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	pageSize	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Tenant
// @Failure 403
// @router / [get]
func (c *TenantController) GetAll() {
	var tenantName string
	var sysName string
	var limit int64 = 10
	var offset int64
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("pageSize"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sysCode
	if v := c.GetString("tenantName"); v != "" {
		tenantName = v
	}
	// sysName
	if v := c.GetString("sysName"); v != "" {
		sysName = v
	}
	result := &out.OperResult{}
	data, err := models.GetTenantList(tenantName, sysName, offset, limit)
	total := models.CountTenantInfo(tenantName, sysName)
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
// @Description update the Tenant
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Tenant	true		"body for Tenant content"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TenantController) Put() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Tenant{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.LastModificationTime = time.Now()
		if err := models.UpdateTenantById(&v); err == nil {
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
// @Description 逻辑删除
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TenantController) Delete() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTenant(id); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetTenantPermission ...
// @Title GetTenantPermission
// @Description 获取企业所有的权限信息
// @Param	sysCode path 	string	true		"The id you want to update"
// @Param	tenId  path 	string	true		"The id you want to update"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /getTenantPermission [get]
func (c *TenantController) GetTenantPermission() {
	result := &out.OperResult{}
	sysCode := c.GetString("sysCode")
	idStr := c.GetString("tenId")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetPerInfoForTenant(sysCode, id)
	if err != nil {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	} else {
		permissionList := tool.ParsePermissionDataForCheckboxUpdate(v)
		result.Result = 1
		result.Data = permissionList
		c.Data["json"] = result
	}
	c.ServeJSON()
}
