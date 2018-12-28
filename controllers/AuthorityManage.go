package controllers

import (
	input "demo/inputmodels"
	"demo/models"
	out "demo/outmodels"
	"demo/tools"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	jwt "github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	out.OperResult
	UserId       int64
	Name         string
	UserName     string
	Gender       string
	Age          string
	Avatar       string
	EmailAddress string
	PhoneNumber  string
	Permissions  []string
}

// 登陆,账户和用户信息管理模块
type AuthorityManageController struct {
	beego.Controller
}

const (
	SecretKey = "sfljdsfjsljdslfdsfsdfjdsf"
)

func (c *AuthorityManageController) URLMapping() {
	c.Mapping("AuthorityError", c.AuthorityError)
	c.Mapping("xsunLogin", c.SysLogin)
}

func (c *AuthorityManageController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

// AuthorityError...
// @Title AuthorityError
// @Description  系统认证错误
// @Success 200
// @router /AuthorityError [get]
func (tc *AuthorityManageController) AuthorityError() {
	result := out.OperResult{}
	result.Result = 0
	result.Message = "未登录"
	tc.Data["json"] = result
	tc.ServeJSON()
}

// SysLogin...
// @Title SysLogin
// @Description  系统登陆跳转
// @Success 200  login.html
// @Failure 404  异常信息
// @router /xsunLogin  [get]
func (tc *AuthorityManageController) SysLogin() {
	/*returnUrl := tc.GetString("ReturnUrl")
	sysID := tc.GetString("SysID")
	tc.Data["ReturnUrl"] = returnUrl
	tc.Data["SysID"] = sysID*/
	tc.TplName = "login.html"
}

// Login...
// @Title Login
// @Description  系统登陆
// @Param   body     body    inputmodels.LoginInfo  true        "登陆信息  username password"
// @Param   Authorization     header    string  false        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @Failure 404 User not found
// @router /Login [post]
func (tc *AuthorityManageController) Login() {
	lresult := &out.LoginResult{}
	originToken := tc.Ctx.Request.Header.Get("Authorization")
	sysCode := tc.Ctx.Request.Header.Get("SysCode")
	// 判断 token 是否有值  token为空表示第一次登陆  不为空验证 token是否有效
	if originToken == "" {
		l := &input.LoginInfo{}
		json.Unmarshal(tc.Ctx.Input.RequestBody, l)
		valid := validation.Validation{}
		resultUserName := valid.Required(l.UserName, "username").Message("请输入用户名")
		if resultUserName.Ok == false {
			lresult.Result = 0
			lresult.Message = resultUserName.Error.Message
			tc.Data["json"] = lresult
			tc.ServeJSON()
			return
		}
		resultPass := valid.Required(l.Password, "password").Message("请输入密码")
		if resultPass.Ok == false {
			lresult.Result = 0
			lresult.Message = resultPass.Error.Message
			tc.Data["json"] = lresult
			tc.ServeJSON()
			return
		}
		resultSysID := valid.Required(sysCode, "sysCode").Message("系统号不能为空")
		if resultSysID.Ok == false {
			lresult.Result = 0
			lresult.Message = resultSysID.Error.Message
			tc.Data["json"] = lresult
			tc.ServeJSON()
			return
		}
		result, user, err := models.LoginCheck(l.UserName, l.Password, sysCode)
		respmessage := ""
		if result == false {
			if err == nil {
				respmessage = "用户名和密码不匹配，重新登陆"
			} else {
				if err == orm.ErrNoRows {
					respmessage = "用户名和密码不匹配，重新登陆"
				} else {
					respmessage = err.Error()
				}
			}
			lresult.Result = 0
			lresult.Message = respmessage
			tc.Data["json"] = lresult
			tc.ServeJSON()
			return
		}
		tc.SetSession("userId", user.Id)
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["jti"] = user.Id
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(10)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(SecretKey))
		//获取用户对应的系统权限
		permissions, _ := models.GetPermissionByUser(user.Id, sysCode)
		permissionData, err := json.Marshal(permissions)
		// 设置 user 信息
		var userOut out.UserInfoToken
		userOut.UserName = user.UserName
		userOut.Phone = user.PhoneNumber
		tokenMap := make(map[string]string)
		tokenMap["ssoId"] = string(user.SsoID)
		jsonUser, _ := json.Marshal(userOut)
		tokenMap["userInfo"] = string(jsonUser)
		tools.InitRedis()
		skey := fmt.Sprintf("%s%s", tokenString, sysCode)
		tools.Globalcluster.Do("set", skey, permissionData)
		tools.Globalcluster.Do("set", tokenString, user.SsoID)
		tools.Globalcluster.Do("EXPIRE", tokenString, 3600)
		tools.Globalcluster.Close()
		lresult.Result = 1
		lresult.Token = tokenString
		tc.Data["json"] = lresult
		tc.Ctx.SetCookie("xy_token", tokenString, 12*3600, "/", ".free.idcfengye.com")
		tc.ServeJSON()
	} else {
		respmessage := &out.OperResult{}
		tools.InitRedis()
		exists, _ := tools.Globalcluster.Do("EXISTS", originToken)
		if exists.(int64) != 0 {
			respmessage.Result = 1
			respmessage.Message = "token有效"
			tc.Data["json"] = respmessage
			tc.ServeJSON()
		} else {
			respmessage.Result = 0
			respmessage.Message = "token失效"
			tc.Data["json"] = respmessage
			tc.ServeJSON()
		}
		tools.Globalcluster.Close()
	}

}

// GetUserInfo...
// @Title GetUserInfo
// @Description 根据TOKEN获取用户信息
// @Param   Authorization     header    string  true        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @Failure 404  User not found
// @router /GetUserInfo [get]
func (tc *AuthorityManageController) GetUserInfo() {
	token := tc.Ctx.Request.Header.Get("Authorization")
	sysCode := tc.Ctx.Request.Header.Get("SysCode")
	result := &out.OperResult{}
	ok, claims, err := tools.CheckLogin(token)
	if !ok {
		result.Result = 0
		result.Message = err.Error()
		tc.Data["json"] = result
		tc.ServeJSON()
		return
	}
	tmp := strconv.FormatFloat(claims["jti"].(float64), 'f', -1, 64)
	userid, _ := strconv.ParseInt(tmp, 10, 64)
	u, _ := models.GetUserById(userid)
	permissions, _ := models.GetPermissionByUser(userid, sysCode)
	var arrPermission []string
	for _, v := range permissions {
		arrPermission = append(arrPermission, v.Name)
	}
	data := make(map[string]interface{})
	data["userInfo"] = u
	data["permissions"] = arrPermission
	result.Result = 1
	result.Data = data
	tc.Data["json"] = result
	tc.ServeJSON()
}

// RegistUser...
// @Title 注册新用户 app
// @Description  APP注册用户
// @Param   body     body    inputmodels.LoginInfo  true        "登陆信息  useraname password"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @Failure 400
// @router /registUser [post]
func (tc *AuthorityManageController) RegistUser() {
	sysCode := tc.Ctx.Request.Header.Get("SysCode")
	l := &input.LoginInfo{}
	json.Unmarshal(tc.Ctx.Input.RequestBody, l)
	lresult := &out.OperResult{}
	if ssoID, err := models.RegistUser(l, sysCode); err == nil {
		lresult.Result = 1
		lresult.Message = "创建用户成功"
		lresult.Data = ssoID
		tc.Data["json"] = lresult
	} else {
		lresult.Result = 1
		lresult.Message = "创建用户失败"
		tc.Data["json"] = lresult
	}
	tc.ServeJSON()
}

// Logout...
// @Title Login
// @Description  退出系统 清除redis保存的token信息
// @Param   Authorization     header    string  true        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @router /Logout [post]
func (tc *AuthorityManageController) Logout() {
	lresult := &out.LoginResult{}
	sysCode := tc.Ctx.Request.Header.Get("SysCode")
	authorization := tc.Ctx.Request.Header.Get("Authorization")
	skey := fmt.Sprintf("%s%s", authorization, sysCode)
	tools.InitRedis()
	tools.Globalcluster.Do("DEL", authorization)
	tools.Globalcluster.Do("DEL", skey)
	tools.Globalcluster.Close()
	lresult.Result = 1
	lresult.Token = ""
	lresult.Message = "退出系统"
	tc.Data["json"] = lresult
	tc.ServeJSON()
}

// PasswdUpdate...
// @Title Login
// @Description 修改密码
// @Param   body     body    inputmodels.LoginInfo  true        "新用户信息  用户名和密码"
// @Param   Authorization     header    string  true        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @router /passwdUpdate [put]
func (tc *AuthorityManageController) PasswdUpdate() {
	sysCode := tc.Ctx.Request.Header.Get("SysCode")
	l := &input.LoginInfo{}
	json.Unmarshal(tc.Ctx.Input.RequestBody, l)
	lresult := &out.OperResult{}
	if err := models.PasswdUpdate(l, sysCode); err == nil {
		lresult.Result = 1
		lresult.Message = "修改用户密码成功"
		tc.Data["json"] = lresult
	} else {
		lresult.Result = 0
		lresult.Message = "修改用户密码"
		tc.Data["json"] = lresult
	}
	tc.ServeJSON()
}
