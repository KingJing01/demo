package routers

// @APIVersion 1.0.0
// @Title 权限系统Swagger
// @Description 欣阳权限系统对接接口api
import (
	"demo/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/application",
			beego.NSInclude(
				&controllers.ApplicationController{},
			),
		),

		beego.NSNamespace("/permission",
			beego.NSInclude(
				&controllers.PermissionController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/ssouser",
			beego.NSInclude(
				&controllers.SsouserController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),

		beego.NSNamespace("/userrole",
			beego.NSInclude(
				&controllers.UserroleController{},
			),
		),

		beego.NSNamespace("/authoritymanage",
			beego.NSInclude(
				&controllers.AuthorityManageController{},
			),
		),

		beego.NSNamespace("/tenant",
			beego.NSInclude(
				&controllers.TenantController{},
			),
		),

		beego.NSNamespace("/setmeal",
			beego.NSInclude(
				&controllers.SetMealController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
