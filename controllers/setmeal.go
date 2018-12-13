package controllers

import (
	input "demo/inputmodels"
	"demo/models"
	out "demo/outmodels"
	"encoding/json"

	"github.com/astaxie/beego"
)

type SetMealController struct {
	beego.Controller
}

// URLMapping ...
func (c *SetMealController) URLMapping() {
	c.Mapping("Post", c.Post)
	//c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	//c.Mapping("check", c.CheckRepeat)
}

// GetAll ...
// @Title Get All
// @Description get Application
// @Param	setMealNAme	query	string	false	""
// @Param	sysName	query	string	false	""
// @Param	pageSize	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Application
// @Failure 403
// @router / [get]
func (c *SetMealController) GetAll() {
	var setMealNAme string
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
	// setMealNAme
	if v := c.GetString("setMealName"); v != "" {
		setMealNAme = v
	}
	// sysName
	if v := c.GetString("sysName"); v != "" {
		sysName = v
	}
	result := &out.OperResult{}
	data, err := models.GetSetMealList(setMealNAme, sysName, offset, limit)
	total := models.CountSetMealInfo(setMealNAme, sysName)
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

// Post ...
// @Title Post
// @Description create SetMeat
// @Param	body		body 	models.Application	true		"body for Application content"
// @Success 201 {int} models.Application
// @Failure 403 body is empty
// @router / [post]
func (c *SetMealController) Post() {
	result := &out.OperResult{}
	var v input.SetMeatInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSetMeal(&v); err == nil {
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
// @Description delete the setmeal
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *SetMealController) Delete() {
	result := &out.OperResult{}
	ids := c.Ctx.Input.Param(":id")
	if err := models.DeleteSetMeal(ids); err == nil {
		result.Result = 1
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}

// UpdateSetMealInfo ...
// @Title updateSetMealInfo
// @Description 修改套餐的信息
// @router /updateSetMealInfo [put]
func (c *SetMealController) UpdateSetMealInfo() {
	result := &out.OperResult{}
	var v input.SetMeatInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.UpdateSetMeal(&v); err == nil {
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
