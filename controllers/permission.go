package controllers

import (
	"demo/models"
	out "demo/outmodels"
	tool "demo/tools"
	"encoding/json"
	"strconv"

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
	userID := c.GetSession("userId")
	if userID == nil {
		result.Result = 0
		result.Message = "seesion失效"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	var mystruct map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct); err == nil {
		if _, err := models.AddPermission(mystruct, userID.(int64)); err == nil {
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
// @Description 获取菜单及对应操作信息
// @Param	id		path 	string	true		"权限ID"
// @Success 200  result:1(success)  0(false)
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
// @Param	pageSize	query	string	false	"Limit the size of result set. Must be an integer"
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
	if v, err := c.GetInt64("pageSize"); err == nil {
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
// @Description 更新权限信息
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Permission	true		"body for Permission content"
// @Success 200 {object} models.Permission
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PermissionController) Put() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tool.GetInfoFromToken(originToken)
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	var mystruct map[string]interface{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct); err == nil {
		if err := models.UpdatePermission(mystruct, id, userID); err == nil {
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

// Delete 逻辑删除权限
// @Title Delete
// @Description 逻辑删除权限
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

// GetPerInfoBySysCode 根据系统编号获取企业下对应的权限
// @Title GetPerInfoBySysCode
// @Description  根据系统编号获取对应的权限
// @Param	sysCode	 path 	string 	true		"系统编号"
// @Success 200  result:1(success)  0(false)
// @Failure 403 sysCode is empty
// @router /getPerInfoBySysCode/:sysCode [get]
func (c *PermissionController) GetPerInfoBySysCode() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, _, _ := tool.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	sysCode := c.Ctx.Input.Param(":sysCode")
	if data, err := models.GetPerInfoBySysCode(sysCode, tenantID); len(data) > 0 {
		permissionList := out.ParsePermissionDataForCheckbox(data)
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

// GetPerInfoBySysCodeUpdate 根据套餐编号和系统编号获取对应的权限
// @Title GetPerInfoBySysCode
// @Description  根据套餐编号和系统编号获取对应的权限
// @Param	sysCode		query 	string 	true		"系统编号"
// @Param	setMealCode		query 	string 	true	"套餐编号"
// @Success 200  result:1(success)  0(false)
// @router /getPerInfoBySysCodeUpdate [get]
func (c *PermissionController) GetPerInfoBySysCodeUpdate() {
	result := &out.OperResult{}
	sysCode := c.GetString("sysCode")
	setMealCode := c.GetString("setMealCode")
	if data, err := models.GetPermBySetMealCode(setMealCode, sysCode); err == nil {
		permissionList := out.ParsePermissionDataForCheckboxUpdate(data)
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

// GetPerInfoByRoleID 根据角色id获取权限
// @Title GetPerInfoByRoleID
// @Description  根据角色编号获取权限
// @Param	roleCode	 path 	string 	true		"系统编号"
// @Success 200  result:1(success)  0(false)
// @Failure 403 sysCode is empty
// @router /getPerInfoByRoleId/:roleId [get]
func (c *PermissionController) GetPerInfoByRoleID() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, userID, _ := tool.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	roleID := c.Ctx.Input.Param(":roleId")
	sysCode := c.GetString("sysCode")
	if data, err := models.GetPerInfoByRoleCode(roleID, sysCode, tenantID, userID); len(data) > 0 {
		permissionList := out.ParsePermissionDataForCheckboxUpdate(data)
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
