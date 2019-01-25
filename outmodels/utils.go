package outmodels

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

const (
	//TMSAddUser  tms 新增用户地址
	TMSAddUser = "httpEdi/user/addUser.do"
)

func GetServerUrl() (url string) {
	return beego.AppConfig.String("TMSAddUser")
}

//ParseUser 请求数据解析
func ParseUser(tmsUser TMSUser) (result *bytes.Buffer) {
	//json序列化
	post := "{\"userCode\":\"" + tmsUser.UserCode +
		"\",\"ssoUid\":\"" + tmsUser.SsoUID +
		"\",\"mobile\":\"" + tmsUser.Mobile +
		"\",\"email\":\"" + tmsUser.Email +
		"\",\"sysId\":\"" + tmsUser.SysID +
		"\",\"companyId\":\"" + tmsUser.CompanyID +
		"\",\"companyName\":\"" + tmsUser.CompanyName +
		"\",\"isAdmin\":\"" + tmsUser.IsAdmin +
		"\",\"contact\":\"" + tmsUser.Contact +
		"\",\"shortCompanyName\":\"" + tmsUser.ShortCompanyName +
		"\"}"
	var jsonStr = []byte(post)
	return bytes.NewBuffer(jsonStr)
}

//SendUserInfoToTms 向tms推送用用户数据
func SendUserInfoToTms(tmsUser TMSUser) (respCode int, err error) {
	jsonStr := ParseUser(tmsUser)
	url := fmt.Sprintf("%s%s", GetServerUrl(), TMSAddUser)
	req, err := http.NewRequest("POST", url, jsonStr)
	//设置请求头为 application/json
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var result TMSRespData
	if err := json.NewDecoder(resp.Body).Decode(&result); err == nil {
		if result.Success == false {
			return 111, errors.New(result.Msg)
		}
	} else {
		return 111, err
	}
	return resp.StatusCode, err
}

/* authData: [
   {
     'childrenList': [
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 'tms.bill.export', 'permissionName': '导出账单' },
       { 'mychecked': false, 'indeterminate': false, 'permissionId':  'tms.bill.peint', 'permissionName': '打印账单' }
     ],
     'permissionName': '财务系统', 'mychecked': false, 'indeterminate': false, 'permissionId': 'tms.bill'
   }]*/
// 将后台查询的数据进行格式转化方便前台使用 新增使用
func ParsePermissionDataForCheckbox(data []PermissionCheckInfo) (result []map[string]interface{}) {
	for _, x := range data {
		mapResult := make(map[string]interface{})
		mapResult["mychecked"] = false
		mapResult["indeterminate"] = false
		// 父节点 中文
		mapResult["permissionName"] = x.DisplayName
		//父节点 缩写
		mapResult["permissionId"] = x.Name
		// 中文权限拆为数组
		displayArr := strings.Split(x.CodeName, ",")
		// 缩写权限 拆为数组
		arr := strings.Split(x.Code, ",")
		var children []map[string]interface{}
		//子列数据拼凑
		for j, t := range displayArr {
			for k, z := range arr {
				perArr := make(map[string]interface{})
				if j == k {
					perArr["permissionName"] = t
					perArr["permissionId"] = z
					perArr["mychecked"] = false
					perArr["indeterminate"] = false
					children = append(children, perArr)
				}
			}
		}
		mapResult["childrenList"] = children
		result = append(result, mapResult)
	}
	return result
}

/*  修改模块时将套餐中已选择的数据设置为已勾选
 *	将后台查询的数据进行格式转化方便前台使用  修改使用
 */
func ParsePermissionDataForCheckboxUpdate(data []PermissionCheckInfo) (result []map[string]interface{}) {
	for _, x := range data {
		mapResult := make(map[string]interface{})
		mapResult["mychecked"] = false
		mapResult["indeterminate"] = false
		// 父节点 中文
		mapResult["permissionName"] = x.DisplayName
		//父节点 缩写
		mapResult["permissionId"] = x.Name
		// 中文权限拆为数组
		displayArr := strings.Split(x.CodeName, ",")
		// 缩写权限 拆为数组
		arr := strings.Split(x.Code, ",")
		// 选中判断拆为数组
		flag := strings.Split(x.Flag, ",")
		var children []map[string]interface{}
		//子列数据拼凑
		for j, t := range displayArr {
			for k, z := range arr {
				perArr := make(map[string]interface{})
				if j == k {
					perArr["permissionName"] = t
					perArr["permissionId"] = z
					if flag[j] == strconv.Itoa(1) {
						perArr["mychecked"] = true
					} else {
						perArr["mychecked"] = false
					}
					perArr["indeterminate"] = false
					children = append(children, perArr)
				}
			}
		}
		mapResult["childrenList"] = children
		result = append(result, mapResult)
	}
	return result
}

//ParseCheckRadioData  解析数据给前台组件使用
func ParseCheckRadioData(data []ComponentData) (result []map[string]interface{}) {
	for _, x := range data {
		mapResult := make(map[string]interface{})
		mapResult["name"] = x.ParentName
		mapResult["key"] = x.ParentKey
		// 中文权限拆为数组
		displayArr := strings.Split(x.ChildName, ",")
		// 缩写权限 拆为数组
		arr := strings.Split(x.ChildKey, ",")
		var children []map[string]interface{}
		//子列数据拼凑
		for j, t := range displayArr {
			for k, z := range arr {
				perArr := make(map[string]interface{})
				if j == k {
					perArr["childName"] = t
					perArr["childCode"] = z
					children = append(children, perArr)
				}
			}
		}
		mapResult["childrenList"] = children
		result = append(result, mapResult)
	}
	return result
}

//ParsePermissionData  解析数据
func ParsePermissionData(data []PerInfo) (result []map[string]interface{}) {
	for _, x := range data {
		mapResult := make(map[string]interface{})
		mapResult["name"] = x.Name
		// 权限拆为数组
		displayArr := strings.Split(x.DisplayName, ",")
		mapResult["childrenList"] = displayArr
		result = append(result, mapResult)
	}
	return result
}

//ParsePermissionList 将数据返回为list数组
func ParsePermissionList(data []PerInfo) (result []string) {
	for _, x := range data {
		result = append(result, x.Name)
		// 权限拆为数组
		displayArr := strings.Split(x.DisplayName, ",")
		for _, y := range displayArr {
			result = append(result, y)
		}
	}
	return result
}
