package controllers

import (
	"demo/models"
	out "demo/outmodels"
	tool "demo/tools"
	tools "demo/tools"
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

//UserController  用户基本信息管理
type UserController struct {
	beego.Controller
}

// URLMapping 路径映射
func (c *UserController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("updateUserValidStatus", c.UpdateUserValidStatus)
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 403 body is empty
// @router / [post]
func (c *UserController) Post() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, userID, _ := tool.GetInfoFromToken(originToken)
	var v models.User
	var mystruct map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct)
	roleIds := tool.ParseInterfaceArr(mystruct["RoleIds"].([]interface{}))
	sysCodes := tool.ParseInterfaceArr(mystruct["SysCodes"].([]interface{}))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.CreatorUserId = userID
		v.CreationTime = time.Now()
		if tmsUser, err := models.AddUser(&v, roleIds, sysCodes, tenantID, userID); err == nil {
			_, err := out.SendUserInfoToTms(tmsUser)
			result.Result = 1
			if err != nil {
				result.Result = 0
				result.Message = err.Error()
			}
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
// @Description get User by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :id is empty
// @router /:id [get]
func (c *UserController) GetOne() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	v, err := models.GetUserByID(id)
	if err == nil {
		result.Data = v
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description 获取用户信息
// @Param	roleName 	query	string	false	"角色名称名称"
// @Param	sysName	query	string	false	"系统名称"
// @Param	userName	query	string	false	"登录名"
// @Param	pageSize	query	string	false	 "一页显示数据量 后台默认为10 "
// @Param	offset	query	string	false	"数据下标"
// @Success 200 [object] models.SetMeal
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, tenantID, _, _ := tools.GetInfoFromToken(originToken)
	var roleName string
	var sysName string
	var userName string
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
	// userName
	if v := c.GetString("userName"); v != "" {
		userName = v
	}

	data, err := models.GetUserList(roleName, sysName, userName, offset, limit, tenantID)
	total := models.CountUserInfo(roleName, sysName, userName, tenantID)
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
// @Description update the User
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.User	true		"body for User content"
// @Success 200 {object} models.User
// @Failure 403 :id is not int
// @router /:id [put]
func (c *UserController) Put() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	var mystruct map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct)
	roleCode := mystruct["RoleCode"].(string)
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tools.GetInfoFromToken(originToken)
	v := models.User{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateUserByID(&v, roleCode, userID); err == nil {
			result.Data = v
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
// @Description delete the USer
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *UserController) Delete() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tools.GetInfoFromToken(originToken)
	ids := c.Ctx.Input.Param(":id")
	if err := models.DeleteUser(ids, userID); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// UpdateUserValidStatus 列表修改用户有效状态
// @Title Delete
// @Description 列表修改角色有效状态
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /updateUserValidStatus/:id [put]
func (c *UserController) UpdateUserValidStatus() {
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tools.GetInfoFromToken(originToken)
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	isValid, _ := c.GetInt64("IsValid")
	if err := models.UpdateUserValidStatus(id, isValid, userID); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}
