/*  基本权限 用户管理模块
 *  引用配置的api
 */
import request from '@/utils/request'

// 获取用户列表信息
export function getUserList(info) {
  return request({
    url: '/user',
    method: 'get',
    params: info
  })
}
// 新增角色信息
export function addUserInfo(info) {
  return request({
    url: '/user',
    method: 'post',
    data: {
      'roleName': info.roleName,
      'perId': info.perId,
      'perName': info.perName,
      'sysCode': info.sysCode
    }
  })
}
// 删除角色信息
export function deleteUser(ids) {
  return request({
    url: '/user/' + ids,
    method: 'delete'
  })
}
// 更新角色信息
export function updateUserInfo(info) {
  return request({
    url: '/user',
    method: 'put',
    data: info
  })
}

// 更新用户有效状态 updateUserValidStatus
export function updateUserValidStatus(info) {
  debugger
  return request({
    url: '/user/updateUserValidStatus/' + info.Id,
    method: 'put',
    data: info.IsValid
  })
}

