package controllers

import (
	"demo/models"
	out "demo/outmodels"

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
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Application
// @Failure 403
// @router / [get]
func (c *SetMealController) GetAll() {
	var setMealNAme string
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
