package tools

import (
	out "demo/outmodels"
	"strings"
)

/* authData: [
   {
     'childrenList': [
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 21, 'permissionName': '订单' },
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 32, 'permissionName': '导出用户12' }
     ],
     'permissionName': '用户服务', 'mychecked': false, 'indeterminate': false, 'permissionId': 15
   },
   {
     'childrenList': [
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 25, 'permissionName': '内容查看' },
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 28, 'permissionName': '我的商城' }
     ],
     'permissionName': '内容商城', 'mychecked': false, 'indeterminate': false, 'permissionId': 16, 'showFlag': '1'
   }],*/
// 将后台查询的数据进行格式转化方便前台使用
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
