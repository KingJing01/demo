
/* 转换权限配置界面下拉框的数据
  数据格式为：
   authData: [
   {
     'childrenList': [
       { 'mychecked': false, 'indeterminate': false, 'permissionId': 'tms.bill.export', 'permissionName': '导出账单' },
       { 'mychecked': false, 'indeterminate': false, 'permissionId':  'tms.bill.peint', 'permissionName': '打印账单' }
     ],
     'permissionName': '财务系统', 'mychecked': false, 'indeterminate': false, 'permissionId': 'tms.bill'
   }]
   * 返回数据 permission{perName：xxx ,perId:xxx }
   */
export function transPermissionCheckedData(data) {
  var permission = {}
  var permissionNameStr = ''
  var permissionIdStr = ''
  var permissionMenu = ''
  for (var i = 0; i < data.length; i++) {
    var children = data[i].childrenList
    for (var j = 0; j < children.length; j++) {
      if (children[j].mychecked === true) {
        permissionNameStr += children[j].permissionName + ','
        permissionIdStr += children[j].permissionId + ','
        if (permissionMenu.indexOf(data[i].permissionName) < 0) {
          permissionMenu += data[i].permissionName + ','
        }
      }
    }
  }
  permission.perName = permissionNameStr.substring(0, permissionNameStr.length - 1)
  permission.perId = permissionIdStr.substring(0, permissionIdStr.length - 1)
  permission.perMenu = permissionMenu.substring(0, permissionMenu.length - 1)
  return permission
}

