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

// @router /AuthorityError [get]
func (tc *AuthorityManageController) AuthorityError() {
	result := out.OperResult{}
	result.Result = 0
	result.Message = "未登录"
	tc.Data["json"] = result
	tc.ServeJSON()
}

// @router /xsunLogin  [get]
func (tc *AuthorityManageController) SysLogin() {
	/*returnUrl := tc.GetString("ReturnUrl")
	sysID := tc.GetString("SysID")
	tc.Data["ReturnUrl"] = returnUrl
	tc.Data["SysID"] = sysID*/
	tc.TplName = "login.html"
}

// @Title Login
// @Description 登入接口
// @Param   key     path    string  true        "The email for login"
// @Success 200 {object} controllers.LoginResult
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /Login [post]
func (tc *AuthorityManageController) Login() {
	lresult := &out.LoginResult{}
	originToken := tc.Ctx.Request.Header.Get("Authorization")
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
		resultSysID := valid.Required(l.SysID, "sysId").Message("系统号不能为空")
		if resultSysID.Ok == false {
			lresult.Result = 0
			lresult.Message = resultSysID.Error.Message
			tc.Data["json"] = lresult
			tc.ServeJSON()
			return
		}
		result, user, err := models.LoginCheck(l.UserName, l.Password, l.SysID)
		respmessage := ""
		if result == false {
			if err == nil {
				respmessage = "用户名和密码不匹配，重新登陆"
			} else {
				respmessage = err.Error()
			}
			lresult.Result = 0
			lresult.Message = respmessage
			tc.Data["json"] = lresult
			tc.ServeJSON()
			return
		}
		token := jwt.New(jwt.SigningMethodHS256)
		claims := make(jwt.MapClaims)
		claims["jti"] = user.Id
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(10)).Unix()
		claims["iat"] = time.Now().Unix()
		token.Claims = claims
		tokenString, err := token.SignedString([]byte(SecretKey))
		//获取用户对应的系统权限
		permissions, _ := models.GetPermissionByUser(user.Id, l.SysID)
		permissionData, err := json.Marshal(permissions)
		// 设置 user 信息
		/*var userOut out.UserInfoToken
		userOut.UserName = user.UserName
		userOut.Phone = user.PhoneNumber
		tokenMap := make(map[string]string)
		tokenMap["ssoId"] = string(user.SsoID)
		jsonUser, _ := json.Marshal(userOut)
		tokenMap["userInfo"] = string(jsonUser)*/
		tools.InitRedis()
		skey := fmt.Sprintf("%s_%s", tokenString, l.SysID)
		tools.Globalcluster.Do("set", skey, permissionData)
		//tokenInfo, _ := json.Marshal(tokenMap)
		tools.Globalcluster.Do("set", tokenString, user.SsoID)
		tools.Globalcluster.Do("EXPIRE", tokenString, 12*3600)
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

//获取用户信息
// @Title GetUserInfo
// @Description 根据TOKEN获取用户信息
// @Param   Authorization     header    string  true        "Token信息"
// @Success 200 {object} controllers.UserInfo
// @Failure 400 Invalid email supplied
// @Failure 404 User not found
// @router /GetUserInfo [get]
func (tc *AuthorityManageController) GetUserInfo() {
	token := tc.Ctx.Request.Header.Get("Authorization")
	oResult := UserInfo{}
	ok, claims, err := tools.CheckLogin(token)
	if !ok {
		oResult.Result = 0
		oResult.Message = err.Error()
		tc.Data["json"] = oResult
		tc.ServeJSON()
		return
	}
	tmp := strconv.FormatFloat(claims["jti"].(float64), 'f', -1, 64)
	userid, _ := strconv.ParseInt(tmp, 10, 64)
	u, _ := models.GetUserById(userid)
	if u != nil {
		oResult.UserName = u.UserName
		oResult.Name = u.Name
		oResult.EmailAddress = u.EmailAddress
		oResult.PhoneNumber = u.PhoneNumber
	}
	permissions, _ := models.GetPermissionByUser(userid, "0")
	var arrPermission []string
	for _, v := range permissions {
		arrPermission = append(arrPermission, v.Name)
	}
	oResult.Permissions = arrPermission
	oResult.Result = 1

	tc.Data["json"] = oResult
	tc.ServeJSON()
}

//登出
// @Title Logout
// @Description 登出
// @router /Logout [post]
func (tc *AuthorityManageController) Logout() {
	lresult := &out.LoginResult{}
	lresult.Result = 1
	lresult.Token = ""
	lresult.Message = "登出成功"
	tc.Data["json"] = lresult
	tc.ServeJSON()
}
