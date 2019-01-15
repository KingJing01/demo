package controllers

import (
	input "demo/inputmodels"
	"demo/models"
	out "demo/outmodels"
	tools "demo/tools"
	"encoding/json"
	"strconv"

	"github.com/astaxie/beego"
)

// RoleController operations for Role
type RoleController struct {
	beego.Controller
}

// URLMapping ...
func (c *RoleController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("updateValidStatus", c.UpdateValidStatus)
}

// Post ...
// @Title Post
// @Description create Role
// @Param	body		body 	models.Role	true		"body for Role content"
// @Success 201 {int} models.Role
// @Failure 403 body is empty
// @router / [post]
func (c *RoleController) Post() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, userID, _ := tools.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	var v input.RoleInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddRole(&v, userID, tenantID); err == nil {
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
// @Description get Role by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Role
// @Failure 403 :id is empty
// @router /:id [get]
func (c *RoleController) GetOne() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRoleById(id)
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
// @Description 获取角色信息
// @Param	roleName 	query	string	false	"角色名称名称"
// @Param	sysName	query	string	false	"系统名称"
// @Param	pageSize	query	string	false	 "一页显示数据量 后台默认为10 "
// @Param	offset	query	string	false	"数据下标"
// @Success 200 [object] models.SetMeal
// @Failure 403
// @router / [get]
func (c *RoleController) GetAll() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, _, _ := tools.GetInfoFromToken(originToken)
	var roleName string
	var sysName string
	var limit int64 = 10
	var offset int64
	// pageSize: 10 (default is 10)
	if v, err := c.GetInt64("pageSize"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// roleName
	if v := c.GetString("roleName"); v != "" {
		roleName = v
	}
	// sysName
	if v := c.GetString("sysName"); v != "" {
		sysName = v
	}
	data, err := models.GetRoleList(roleName, sysName, offset, limit, tenantID)
	total := models.CountRoleInfo(roleName, sysName, tenantID)
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
// @Description update the Role
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Role	true		"body for Role content"
// @Success 200 {object} models.Role
// @Failure 403 :id is not int
// @router / [put]
func (c *RoleController) Put() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tools.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	var v input.RoleInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateRoleById(&v, userID); err == nil {
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
// @Description delete the Role
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *RoleController) Delete() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tools.GetInfoFromToken(originToken)
	ids := c.Ctx.Input.Param(":id")
	if err := models.DeleteRole(ids, userID); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// UpdateValidStatus 列表修改角色有效状态
// @Title Delete
// @Description 列表修改角色有效状态
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /updateValidStatus/:id [put]
func (c *RoleController) UpdateValidStatus() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tools.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	isValid, _ := c.GetInt64("IsValid")
	if err := models.UpdateValidStatus(id, isValid, userID); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetRoleBySysCode 根据系统获取系统下的角色
// @Title UpdateValidStatus
// @Description 根据系统获取系统下的角色
// @Param	sysCodes		query 	string	true		"需要查询的系统编号"
// @Success 200 {string} get success!
// @Failure 403 sysCodes is empty
// @router /getRoleBySysCode [get]
func (c *RoleController) GetRoleBySysCode() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, _, _ := tools.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	sysCode := c.GetString("sysCode")
	if data, err := models.GetRoleBySysCode(sysCode, tenantID); err == nil {
		radioData := out.ParseCheckRadioData(data)
		result.Result = 1
		result.Data = radioData
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}
