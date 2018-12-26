package controllers

import (
	input "demo/inputmodels"
	"demo/models"
	out "demo/outmodels"
	"encoding/json"

	"github.com/astaxie/beego"
)

// 套餐信息管理
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
	c.Mapping("GetSetMealRadio", c.GetSetMealRadio)
	//c.Mapping("check", c.CheckRepeat)
}

// GetAll ...
// @Title Get All
// @Description 获取套餐信息
// @Param	setMealName 	query	string	false	"套餐名称"
// @Param	sysName	query	string	false	"系统名称"
// @Param	pageSize	query	string	false	 "一页显示数据量 后台默认为10 "
// @Param	offset	query	string	false	"数据下标"
// @Success 200 [object] models.SetMeal
// @Failure 403
// @router / [get]
func (c *SetMealController) GetAll() {
	var setMealName string
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
	// setMealName
	if v := c.GetString("setMealName"); v != "" {
		setMealName = v
	}
	// sysName
	if v := c.GetString("sysName"); v != "" {
		sysName = v
	}
	result := &out.OperResult{}
	data, err := models.GetSetMealList(setMealName, sysName, offset, limit)
	total := models.CountSetMealInfo(setMealName, sysName)
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
// @Param	body		body 	models.SetMeal	true		"body for SetMeal content"
// @Success 201 {int} models.SetMeal
// @Failure 403 body is empty
// @router / [post]
func (c *SetMealController) Post() {
	result := &out.OperResult{}
	userID := c.GetSession("userId")
	if userID == nil {
		result.Result = 0
		result.Message = "seesion失效"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	var v input.SetMeatInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddSetMeal(&v, userID.(int64)); err == nil {
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
// @Description  禁用套餐
// @Param	id 	path 	string	true	"需要禁用的套餐ID 例a,b,c 或 a"
// @Success 200  result:1(success)  0(false)
// @Failure 403 ids is empty
// @router /:id [delete]
func (c *SetMealController) Delete() {
	result := &out.OperResult{}
	userID := c.GetSession("userId")
	if userID == nil {
		result.Result = 0
		result.Message = "seesion失效"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	ids := c.Ctx.Input.Param(":id")
	if err := models.DeleteSetMeal(ids, userID.(int64)); err == nil {
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
// @Param   body     body    inputmodels.SetMeatInput  true       "套餐实体  "
// @Success 200  result:1(success)  0(false)
// @router /updateSetMealInfo [put]
func (c *SetMealController) UpdateSetMealInfo() {
	result := &out.OperResult{}
	userID := c.GetSession("userId")
	if userID == nil {
		result.Result = 0
		result.Message = "seesion失效"
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	var v input.SetMeatInput
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.UpdateSetMeal(&v, userID.(int64)); err == nil {
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

// GetSetMealRadio ...
// @Title GetSetMealRadio
// @Description 获取套餐单选数据
// @Param   sysCode  path 	string	true	"根据系统编号获取单选按钮数据"
// @Success 200  result:1(success)  0(false)
// @router /getSetMealRadio/:sysCode [get]
func (c *SetMealController) GetSetMealRadio() {
	sysCodeStr := c.Ctx.Input.Param(":sysCode")
	result := &out.OperResult{}
	if data, err := models.GetSetMealRadio(sysCodeStr); err == nil {
		result.Result = 1
		result.Data = data
		c.Data["json"] = result
	} else {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
	}
	c.ServeJSON()
}
