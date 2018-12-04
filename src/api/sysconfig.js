/*
 *权限系统 引用配置的api
 */
import request from '@/utils/request'

//  界面搜索查询的方法(带分页) 支持模糊搜索
export function getListData(sysCode, sysName, pageSize, offset) {
  return request({
    url: '/application',
    method: 'get',
    params: {
      sysCode,
      sysName,
      pageSize,
      offset
    }
  })
}
