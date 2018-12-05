/*
 *权限系统 引用配置的api
 */
import request from '@/utils/request'

//  界面搜索查询的方法(带分页) 支持模糊搜索
export function getListData(info) {
  const query = 'SysName:' + info.sysName + ',SysCode:' + info.sysCode
  return request({
    url: '/application',
    method: 'get',
    params: {
      'query': query,
      'fields': 'SysName,SysCode,Id,IsValid',
      'limit': info.pageSize,
      'offset': info.offset
    }
  })
}

//  新增系统的信息
export function saveSysInfo(info) {
  return request({
    url: '/application',
    method: 'post',
    data: {
      'SysName': info.SysName,
      'IsValid': info.IsValid
    }
  })
}
// 系统名称验证重复
export function uniqueCheck(sysName) {
  return request({
    url: '/application/checkRepeat',
    method: 'get',
    params: {
      'SysName': sysName
    }
  })
}
