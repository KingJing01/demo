// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

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
	)
	beego.AddNamespace(ns)
}
