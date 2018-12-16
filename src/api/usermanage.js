// 用户管理的js
import request from '@/utils/request'

// 获取菜单列表信息
export function getUserList(info) {
  return request({
    url: '/tenant',
    method: 'get',
    params: info
  })
}

export function getUserInfo(tenId) {
  return request({
    url: '/tenant/' + tenId,
    method: 'get'
  })
}

export function getUserPermission(info) {
  return request({
    url: '/tenant/getTenantPermission',
    method: 'get',
    params: {
      sysCode: info.sysCode,
      tenId: info.tenId
    }
  })
}

export function updateTenantInfo(info) {
  return request({
    url: '/tenant/' + info.Id,
    method: 'put',
    data: info
  })
}
