package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["demo/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["demo/controllers:ApplicationController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["demo/controllers:ApplicationController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["demo/controllers:ApplicationController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["demo/controllers:ApplicationController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["demo/controllers:ApplicationController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:ApplicationController"] = append(beego.GlobalControllerRouter["demo/controllers:ApplicationController"],
        beego.ControllerComments{
            Method: "CheckRepeat",
            Router: `/checkRepeat`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"] = append(beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"],
        beego.ControllerComments{
            Method: "AuthorityError",
            Router: `/AuthorityError`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"] = append(beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"],
        beego.ControllerComments{
            Method: "GetUserInfo",
            Router: `/GetUserInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"] = append(beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/Login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"] = append(beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/Logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"] = append(beego.GlobalControllerRouter["demo/controllers:AuthorityManageController"],
        beego.ControllerComments{
            Method: "SysLogin",
            Router: `/xsunLogin`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:PermissionController"] = append(beego.GlobalControllerRouter["demo/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:PermissionController"] = append(beego.GlobalControllerRouter["demo/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:PermissionController"] = append(beego.GlobalControllerRouter["demo/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:PermissionController"] = append(beego.GlobalControllerRouter["demo/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:PermissionController"] = append(beego.GlobalControllerRouter["demo/controllers:PermissionController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:RoleController"] = append(beego.GlobalControllerRouter["demo/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:RoleController"] = append(beego.GlobalControllerRouter["demo/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:RoleController"] = append(beego.GlobalControllerRouter["demo/controllers:RoleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:RoleController"] = append(beego.GlobalControllerRouter["demo/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:RoleController"] = append(beego.GlobalControllerRouter["demo/controllers:RoleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:SsouserController"] = append(beego.GlobalControllerRouter["demo/controllers:SsouserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:SsouserController"] = append(beego.GlobalControllerRouter["demo/controllers:SsouserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:SsouserController"] = append(beego.GlobalControllerRouter["demo/controllers:SsouserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:SsouserController"] = append(beego.GlobalControllerRouter["demo/controllers:SsouserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:SsouserController"] = append(beego.GlobalControllerRouter["demo/controllers:SsouserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:TenantController"] = append(beego.GlobalControllerRouter["demo/controllers:TenantController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:TenantController"] = append(beego.GlobalControllerRouter["demo/controllers:TenantController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:TenantController"] = append(beego.GlobalControllerRouter["demo/controllers:TenantController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:TenantController"] = append(beego.GlobalControllerRouter["demo/controllers:TenantController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:TenantController"] = append(beego.GlobalControllerRouter["demo/controllers:TenantController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserController"] = append(beego.GlobalControllerRouter["demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserController"] = append(beego.GlobalControllerRouter["demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserController"] = append(beego.GlobalControllerRouter["demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserController"] = append(beego.GlobalControllerRouter["demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserController"] = append(beego.GlobalControllerRouter["demo/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserroleController"] = append(beego.GlobalControllerRouter["demo/controllers:UserroleController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserroleController"] = append(beego.GlobalControllerRouter["demo/controllers:UserroleController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserroleController"] = append(beego.GlobalControllerRouter["demo/controllers:UserroleController"],
        beego.ControllerComments{
            Method: "GetOne",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserroleController"] = append(beego.GlobalControllerRouter["demo/controllers:UserroleController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["demo/controllers:UserroleController"] = append(beego.GlobalControllerRouter["demo/controllers:UserroleController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
