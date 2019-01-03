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
export function addRoleInfo(info) {
  return request({
    url: '/role',
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
export function deleteRole(ids) {
  return request({
    url: '/role/' + ids,
    method: 'delete'
  })
}
// 更新角色信息
export function updateRoleInfo(info) {
  return request({
    url: '/role',
    method: 'put',
    data: info
  })
}

// 更新角色有效状态 updateValidStatus
export function updateValidStatus(info) {
  return request({
    url: '/role/updateValidStatus/' + info.Id,
    method: 'put',
    data: info.IsValid
  })
}

