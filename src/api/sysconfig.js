/*
 * 系统配置
 *权限系统 引用配置的api
 */
import request from '@/utils/request'

//  界面搜索查询的方法(带分页) 支持模糊搜索
export function getListData(info) {
  return request({
    url: '/application',
    method: 'get',
    params: info
  })
}

//  新增系统的信息
export function saveSysInfo(info) {
  return request({
    url: '/application',
    method: 'post',
    data: {
      'SysName': info.sysName,
      'IsValid': info.IsValid === true ? 0 : 1,
      'SysUrl': info.sysUrl
    }
  })
}
// 系统名称验证重复
export function uniqueCheck(sysName, sysId) {
  return request({
    url: '/application/checkRepeat',
    method: 'get',
    params: {
      'SysName': sysName,
      'SysId': sysId
    }
  })
}

// 修改系统信息
export function updateSysInfo(info) {
  return request({
    url: '/application/' + info.id,
    method: 'put',
    data: {
      'SysName': info.sysName,
      'IsValid': info.IsValid === true ? 0 : 1,
      'SysCode': info.sysCode,
      'SysUrl': info.sysUrl
    }
  })
}

// select 系统下拉数据
export function sysDataSelect() {
  return request({
    url: '/application/getSelectData',
    method: 'post'
  })
}
