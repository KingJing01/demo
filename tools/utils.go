package tools

import (
	out "demo/outmodels"
	"strconv"
	"strings"
)

/* authData: [
   {
     'childrenList': [
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 'tms.bill.export', 'permissionName': '导出账单' },
       { 'mychecked': false, 'indeterminate': false, 'permissionId':  'tms.bill.peint', 'permissionName': '打印账单' }
     ],
     'permissionName': '财务系统', 'mychecked': false, 'indeterminate': false, 'permissionId': 'tms.bill'
   }]*/
// 将后台查询的数据进行格式转化方便前台使用 新增使用
func ParsePermissionDataForCheckbox(data []out.PermissionCheckInfo) (result []map[string]interface{}) {
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
func ParsePermissionDataForCheckboxUpdate(data []out.PermissionCheckInfo) (result []map[string]interface{}) {
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

// 将前台的字符串数组转化 []interface{} 转 []string
func ParseInterfaceArr(params []interface{}) (param []string) {
	strArray := make([]string, len(params))
	for i, arg := range params {
		strArray[i] = arg.(string)
	}
	return strArray
}

//ParseCheckRadioData  解析数据给前台组件使用
func ParseCheckRadioData(data []out.ComponentData) (result []map[string]interface{}) {
	for _, x := range data {
		mapResult := make(map[string]interface{})
		mapResult["name"] = x.ParentName
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
