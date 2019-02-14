package controllers

import (
	input "demo/inputmodels"
	"demo/models"
	out "demo/outmodels"
	"demo/tools"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	jwt "github.com/dgrijalva/jwt-go"
)

//UserInfo  用户信息实体
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

//AuthorityManageController  登陆,账户和用户信息管理模块
type AuthorityManageController struct {
	beego.Controller
}

// token 生成
const (
	SecretKey = "sfljdsfjsljdslfdsfsdfjdsf"
)

//URLMapping  路径映射
func (c *AuthorityManageController) URLMapping() {
	c.Mapping("AuthorityError", c.AuthorityError)
	c.Mapping("xsunLogin", c.SysLogin)
}

//Options options请求的返回
func (c *AuthorityManageController) Options() {
	c.Data["json"] = map[string]interface{}{"status": 200, "message": "ok", "moreinfo": ""}
	c.ServeJSON()
}

// AuthorityError 认证错误
// @Title AuthorityError
// @Description  系统认证错误
// @Success 200
// @router /AuthorityError [get]
func (c *AuthorityManageController) AuthorityError() {
	result := out.OperResult{}
	result.Result = 0
	result.Message = "未登录"
	c.Data["json"] = result
	c.ServeJSON()
}

// SysLogin 系统登陆跳转
// @Title SysLogin
// @Description  系统登陆跳转
// @Success 200  login.html
// @Failure 404  异常信息
// @router /xsunLogin  [get]
func (c *AuthorityManageController) SysLogin() {
	/*returnUrl := tc.GetString("ReturnUrl")
	sysID := tc.GetString("SysID")
	tc.Data["ReturnUrl"] = returnUrl
	tc.Data["SysID"] = sysID*/
	c.TplName = "login.html"
}

// Login PC端系统登陆
// @Title Login
// @Description  PC端系统系统登陆
// @Param   body     body    inputmodels.LoginInfo  true        "登陆信息  username password"
// @Param   Authorization     header    string  false        "Token信息"
// @Success 200  result:1(success)  0(false)
// @Failure 404 User not found
// @router /Login [post]
func (c *AuthorityManageController) Login() {
	lresult := &out.LoginResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	sysCode := c.Ctx.Request.Header.Get("SysCode")
	// 判断 token 是否有值  token为空表示第一次登陆  不为空验证 token是否有效
	if originToken == "" {
		l := &input.LoginInfo{}
		json.Unmarshal(c.Ctx.Input.RequestBody, l)
		valid := validation.Validation{}
		resultUserName := valid.Required(l.UserName, "username").Message("请输入用户名")
		if resultUserName.Ok == false {
			lresult.Result = 0
			lresult.Message = resultUserName.Error.Message
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		resultPass := valid.Required(l.Password, "password").Message("请输入密码")
		if resultPass.Ok == false {
			lresult.Result = 0
			lresult.Message = resultPass.Error.Message
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		resultSysID := valid.Required(sysCode, "sysCode").Message("身份不能为空")
		if resultSysID.Ok == false {
			lresult.Result = 0
			lresult.Message = resultSysID.Error.Message
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		userTotal, _ := models.LoginValidUser(l.UserName)
		if userTotal == 0 {
			lresult.Result = 0
			lresult.Message = "用户信息不存在"
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		// 获取登陆信息有哪些系统
		sysCodeStr, _ := models.LoginValidSys(l.UserName)
		if sysCodeStr != "" {
			if flag := strings.Contains(sysCodeStr, sysCode); flag == false {
				lresult.Result = 0
				lresult.Message = "系统身份选择错误"
				c.Data["json"] = lresult
				c.ServeJSON()
				return
			}
		}

		//验证用户登陆信息
		result, user, err := models.LoginCheck(l.UserName, l.Password, sysCode)
		respmessage := ""
		if result == false {
			if err == nil {
				respmessage = "用户名密码不匹配,请重新登陆"
			} else {
				if err == orm.ErrNoRows {
					respmessage = "用户名密码不匹配,请重新登陆"
				} else {
					respmessage = err.Error()
				}
			}
			lresult.Result = 0
			lresult.Message = respmessage
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		} else {
			if user.IsValid == 1 {
				respmessage = "用户被禁用，请联系管理员"
				lresult.Result = 0
				lresult.Message = respmessage
				c.Data["json"] = lresult
				c.ServeJSON()
				return
			}
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		// jwt 唯一标识存放userId
		claims["jti"] = user.Id
		// jwt 有效时间
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(60)).Unix()
		// jwt 发布时间
		claims["iat"] = time.Now().Unix()
		// jwt 发布者 存放用户
		claims["iss"] = user.TenantId
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(SecretKey))
		//获取用户对应的系统权限
		permissions := models.GetPermissionByUser(user.Id, sysCode)
		arrPermission := out.ParsePermissionList(permissions)
		authData, _ := json.Marshal(arrPermission)
		// 设置 user 信息
		var userOut out.UserInfoToken
		userOut.UserName = user.UserName
		userOut.Phone = user.PhoneNumber
		tokenMap := make(map[string]string)
		tokenMap["ssoId"] = string(user.SsoID)
		jsonUser, _ := json.Marshal(userOut)
		tokenMap["userInfo"] = string(jsonUser)
		tools.InitRedis()
		skey := fmt.Sprintf("%s%s", strconv.FormatInt(user.SsoID, 10), sysCode)
		tools.Globalcluster.Do("set", skey, authData)
		tools.Globalcluster.Do("set", tokenString, user.SsoID)
		tools.Globalcluster.Do("EXPIRE", tokenString, 3600)
		tools.Globalcluster.Close()
		lresult.Result = 1
		lresult.Token = tokenString
		c.Data["json"] = lresult
		c.Ctx.SetCookie("xy_token", tokenString, 12*3600, "/", ".free.idcfengye.com")
		c.ServeJSON()
	} else {
		respmessage := &out.OperResult{}
		tools.InitRedis()
		exists, _ := tools.Globalcluster.Do("EXISTS", originToken)
		if exists.(int64) != 0 {
			respmessage.Result = 1
			respmessage.Message = "token有效"
			c.Data["json"] = respmessage
			c.ServeJSON()
		} else {
			respmessage.Result = 0
			respmessage.Message = "token失效"
			c.Data["json"] = respmessage
			c.ServeJSON()
		}
		tools.Globalcluster.Close()
	}

}

// AuthLogin 权限端系统登陆
// @Title Login
// @Description  权限端系统登陆
// @Param   body     body    inputmodels.LoginInfo  true        "登陆信息  username password"
// @Param   Authorization     header    string  false        "Token信息"
// @Success 200  result:1(success)  0(false)
// @Failure 404 User not found
// @router /authLogin [post]
func (c *AuthorityManageController) AuthLogin() {
	lresult := &out.LoginResult{}
	originToken := c.Ctx.Request.Header.Get("Authorization")
	// 判断 token 是否有值  token为空表示第一次登陆  不为空验证 token是否有效
	if originToken == "" {
		l := &input.LoginInfo{}
		json.Unmarshal(c.Ctx.Input.RequestBody, l)
		valid := validation.Validation{}
		resultUserName := valid.Required(l.UserName, "username").Message("请输入用户名")
		if resultUserName.Ok == false {
			lresult.Result = 0
			lresult.Message = resultUserName.Error.Message
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		resultPass := valid.Required(l.Password, "password").Message("请输入密码")
		if resultPass.Ok == false {
			lresult.Result = 0
			lresult.Message = resultPass.Error.Message
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		result, user, err := models.AuthLoginCheck(l.UserName, l.Password)
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
			c.Data["json"] = lresult
			c.ServeJSON()
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		// jwt 唯一标识存放userId
		claims["jti"] = user.Id
		// jwt 有效时间
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(60)).Unix()
		// jwt 发布时间
		claims["iat"] = time.Now().Unix()
		// jwt 发布者 存放用户
		claims["iss"] = user.TenantId
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(SecretKey))
		//获取用户对应的系统权限
		permissions := models.GetPermissionByUser(user.Id, "")
		arrPermission := out.ParsePermissionList(permissions)
		authData, _ := json.Marshal(arrPermission)
		// 设置 user 信息
		var userOut out.UserInfoToken
		userOut.UserName = user.UserName
		userOut.Phone = user.PhoneNumber
		tokenMap := make(map[string]string)
		tokenMap["ssoId"] = string(user.SsoID)
		jsonUser, _ := json.Marshal(userOut)
		tokenMap["userInfo"] = string(jsonUser)
		tools.InitRedis()
		skey := fmt.Sprintf("%s%s", strconv.FormatInt(user.SsoID, 10), "uam")
		tools.Globalcluster.Do("set", skey, authData)
		tools.Globalcluster.Do("set", tokenString, user.SsoID)
		tools.Globalcluster.Do("EXPIRE", tokenString, 3600)
		tools.Globalcluster.Close()
		lresult.Result = 1
		lresult.Token = tokenString
		c.Data["json"] = lresult
		c.ServeJSON()
	} else {
		respmessage := &out.OperResult{}
		tools.InitRedis()
		exists, _ := tools.Globalcluster.Do("EXISTS", originToken)
		if exists.(int64) != 0 {
			respmessage.Result = 1
			respmessage.Message = "token有效"
			c.Data["json"] = respmessage
			c.ServeJSON()
		} else {
			respmessage.Result = 0
			respmessage.Message = "token失效"
			c.Data["json"] = respmessage
			c.ServeJSON()
		}
		tools.Globalcluster.Close()
	}
}

// GetUserPermission 根据TOKEN获取用户信息 已废弃
// @Title GetUserPermission
// @Description 根据TOKEN获取用户信息
// @Param   Authorization     header    string  true        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @Failure 404  User not found
// @router /GetUserPermission [get]
func (c *AuthorityManageController) GetUserPermission() {
	token := c.Ctx.Request.Header.Get("Authorization")
	sysCode := c.Ctx.Request.Header.Get("SysCode")
	result := &out.OperResult{}
	ok, claims, err := tools.CheckLogin(token)
	if !ok {
		result.Result = 0
		result.Message = err.Error()
		c.Data["json"] = result
		c.ServeJSON()
		return
	}
	tmp := strconv.FormatFloat(claims["jti"].(float64), 'f', -1, 64)
	userid, _ := strconv.ParseInt(tmp, 10, 64)
	permissions := models.GetPermissionByUser(userid, sysCode)
	arrPermission := out.ParsePermissionList(permissions)
	data := make(map[string]interface{})
	data["permissions"] = arrPermission
	result.Result = 1
	result.Data = data
	c.Data["json"] = result
	c.ServeJSON()
}

// RegistUser APP注册用户
// @Title 注册新用户 app
// @Description  APP注册用户
// @Param   body     body    inputmodels.LoginInfo  true        "登陆信息  useraname password"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @Failure 400
// @router /registUser [post]
func (c *AuthorityManageController) RegistUser() {
	sysCode := c.Ctx.Request.Header.Get("SysCode")
	l := &input.LoginInfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, l)
	lresult := &out.OperResult{}
	if ssoID, err := models.RegistUser(l, sysCode); err == nil {
		lresult.Result = 1
		lresult.Message = "创建用户成功"
		lresult.Data = ssoID
		c.Data["json"] = lresult
	} else {
		lresult.Result = 1
		lresult.Message = "创建用户失败"
		c.Data["json"] = lresult
	}
	c.ServeJSON()
}

// Logout 退出系统
// @Title Login
// @Description  退出系统 清除redis保存的token信息
// @Param   Authorization     header    string  true        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @router /Logout [post]
func (c *AuthorityManageController) Logout() {
	lresult := &out.LoginResult{}
	authorization := c.Ctx.Request.Header.Get("Authorization")
	tools.InitRedis()
	tools.Globalcluster.Do("DEL", authorization)
	tools.Globalcluster.Close()
	lresult.Result = 1
	lresult.Token = ""
	lresult.Message = "退出系统"
	c.Data["json"] = lresult
	c.ServeJSON()
}

// PasswdUpdate 修改密码
// @Title Login
// @Description 修改密码
// @Param   body     body    inputmodels.LoginInfo  true        "新用户信息  用户名和密码"
// @Param   Authorization     header    string  true        "Token信息"
// @Param   SysCode     header    string  true        "系统编码"
// @Success 200  result:1(success)  0(false)
// @router /passwdUpdate [put]
func (c *AuthorityManageController) PasswdUpdate() {
	sysCode := c.Ctx.Request.Header.Get("SysCode")
	l := &input.LoginInfo{}
	json.Unmarshal(c.Ctx.Input.RequestBody, l)
	lresult := &out.OperResult{}
	if err := models.PasswdUpdate(l, sysCode); err == nil {
		lresult.Result = 1
		lresult.Message = "修改用户密码成功"
		c.Data["json"] = lresult
	} else {
		lresult.Result = 0
		lresult.Message = "修改用户密码"
		c.Data["json"] = lresult
	}
	c.ServeJSON()
}

// ValidUserActPermission 验证用户的操作权限是否有效
// @Title Login
// @Description 验证用户的操作权限是否有效
// @Param   Authorization     header    string  true        "Token信息"
// @Param   menuCode     query    string  true        "权限编码"
// @Success 200  result:1(success)  0(false)
// @router /validUserActPermission [post]
func (c *AuthorityManageController) ValidUserActPermission() {
	token := c.Ctx.Request.Header.Get("Authorization")
	var mystruct map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &mystruct)
	menuCode := mystruct["menuCode"].(string)
	lresult := &out.OperResult{}
	if flag, _, _ := tools.CheckAuthority(token, menuCode); flag == true {
		lresult.Result = 1
		c.Data["json"] = lresult
	} else {
		lresult.Result = 0
		c.Data["json"] = lresult
	}
	c.ServeJSON()
}
