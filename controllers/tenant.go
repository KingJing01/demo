package controllers

import (
	"demo/models"
	out "demo/outmodels"
	tool "demo/tools"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

//TenantController  企业信息管理模块
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
// @Success 200  result:1(success)  0(false)
// @Failure 403 body is empty
// @router / [post]
func (c *TenantController) Post() {
	result := &out.OperResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tool.GetInfoFromToken(originToken)
	var v models.Tenant
	var mystruct map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct)
	sysCode := tool.ParseInterfaceArr(mystruct["sysCode"].([]interface{}))
	perID := tool.ParseInterfaceArr(mystruct["perId"].([]interface{}))
	perMenu := tool.ParseInterfaceArr(mystruct["perMenu"].([]interface{}))
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.LastModificationTime = time.Now()
		if err, tmsUser := models.AddTenant(&v, sysCode, perID, perMenu, userID); err == nil {
			tool.InitRedis()
			jsonBytes, _ := json.Marshal(v)
			tool.Globalcluster.Do("set", v.Id, string(jsonBytes))
			tool.Globalcluster.Close()
			respCode, _ := out.SendUserInfoToTms(tmsUser)
			fmt.Println("################接口返回的标记值################ ", respCode)
			if respCode != 200 {
				result.Message = "数据已入库,tms推送失败"
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
// @Description get Tenant by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Tenant
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TenantController) GetOne() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
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
// @Param	tenantName	query	string	false	"用户名称"
// @Param	sysName	query	string	false	"系统名称"
// @Param	pageSize	query	string	false	"一页显示数据量 后台默认为10 "
// @Param	offset	query	string	false	"数据下标"
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
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tool.GetInfoFromToken(originToken)
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	v := models.Tenant{Id: id}
	var mystruct map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct)
	sysCode := mystruct["sysCode"].(string)
	perIDStr := mystruct["perId"].(string)
	perMenu := mystruct["perMenu"].(string)
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.LastModificationTime = time.Now()
		if err := models.UpdateTenantById(&v, sysCode, perIDStr, perMenu, v.Id, userID); err == nil {
			tool.InitRedis()
			jsonBytes, _ := json.Marshal(v)
			tool.Globalcluster.Do("set", v.Id, string(jsonBytes))
			tool.Globalcluster.Close()
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
	originToken := c.Ctx.Request.Header.Get("Authorization")
	_, _, userID, _ := tool.GetInfoFromToken(originToken)
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	if err := models.DeleteTenant(id, userID); err == nil {
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
// @Param	sysCode query    	string	true		"系统编码"
// @Param	tenId  query    	string	true		"企业ID"
// @Success 200  result:1(success)  0(false)
// @router /getTenantPermission [get]
func (c *TenantController) GetTenantPermission() {
	result := &out.OperResult{}
	sysCode := c.GetString("sysCode")
	idStr := c.GetString("tenId")
	id, _ := strconv.ParseInt(idStr, 10, 64)
	v, err := models.GetPerInfoForTenant(sysCode, id)
	if err != nil {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	} else {
		permissionList := out.ParsePermissionDataForCheckboxUpdate(v)
		result.Result = 1
		result.Data = permissionList
		c.Data["json"] = result
	}
	c.ServeJSON()
}
