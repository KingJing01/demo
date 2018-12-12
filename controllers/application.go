package controllers

import (
	"demo/models"
	out "demo/outmodels"
	"encoding/json"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// ApplicationController operations for Application
type ApplicationController struct {
	beego.Controller
}

// URLMapping ...
func (c *ApplicationController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("Check", c.CheckRepeat)
	c.Mapping("GetSelectData", c.GetSelectData)
}

// Post ...
// @Title Post
// @Description create Application
// @Param	body		body 	models.Application	true		"body for Application content"
// @Success 201 {int} models.Application
// @Failure 403 body is empty
// @router / [post]
func (c *ApplicationController) Post() {
	result := &out.OperResult{}
	var v models.Application
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.CreationTime = time.Now()
		//生成系统编号
		v.SysCode = models.GenerateSysCode()
		if _, err := models.AddApplication(&v); err == nil {
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
// @Description get Application by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Application
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ApplicationController) GetOne() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetApplicationById(id)
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
// @Description get Application
// @Param	sysCode	query	string	false	""
// @Param	sysName	query	string	false	""
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Application
// @Failure 403
// @router / [get]
func (c *ApplicationController) GetAll() {
	var sysCode string
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
	// sysCode
	if v := c.GetString("sysCode"); v != "" {
		sysCode = v
	}
	// sysName
	if v := c.GetString("sysName"); v != "" {
		sysName = v
	}
	result := &out.OperResult{}
	data, err := models.GetApplicationList(sysCode, sysName, offset, limit)
	total := models.CountApplicationInfo(sysCode, sysName)
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

// Delete ...
// @Title Delete
// @Description delete the Application
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ApplicationController) Delete() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteApplication(id); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// check ...
// @router /checkRepeat [get]
func (c *ApplicationController) CheckRepeat() {
	result := &out.OperResult{}
	sysName := c.GetString("SysName")
	if total, _ := models.CheckRepeat(sysName); total > 0 {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// 获取下拉框的数据...
// @router /getSelectData [post]
func (c *ApplicationController) GetSelectData() {
	result := &out.OperResult{}
	if sysInfo := models.GetSelectData(); len(sysInfo) > 0 {
		result.Result = 1
		result.Data = sysInfo
		c.Data["json"] = result
	} else {
		result.Result = 0
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Application
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Application	true		"body for Application content"
// @Success 200 {object} models.Application
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ApplicationController) Put() {
	result := &out.OperResult{}
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.Application{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.LastModificationTime = time.Now()
		if err := models.UpdateApplicationById(&v); err == nil {
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
