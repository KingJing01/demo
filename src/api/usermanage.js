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

// 获取租户表的信息
export function getUserInfo(tenId) {
  return request({
    url: '/tenant/' + tenId,
    method: 'get'
  })
}

// 获取租户下对应的系统权限
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

// 更新租户信息
export function updateTenantInfo(info) {
  return request({
    url: '/tenant/' + info.Id,
    method: 'put',
    data: info
  })
}

// 保存租户信息
export function saveTenantInfo(info) {
  return request({
    url: '/tenant',
    method: 'post',
    data: info
  })
}
