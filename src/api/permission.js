/*  基本权限
 *  引用配置的api
 */
import request from '@/utils/request'

// 获取菜单列表信息
export function getMenuList(info) {
  return request({
    url: '/permission',
    method: 'get',
    params: info
  })
}
// 根据系统编号获取对应的权限
export function getPerInfoBySysCode(sysCode) {
  return request({
    url: '/permission/getPerInfoBySysCode/' + sysCode,
    method: 'get'
  })
}

// 修改功能下，根据系统编号和套餐编号获取选中的信息状态
export function getPerInfoBySysCodeUpdate(sysCode, setMealCode) {
  return request({
    url: '/permission/getPerInfoBySysCodeUpdate',
    method: 'get',
    params: {
      sysCode, setMealCode
    }
  })
}

// 获取角色下的权限
export function getPerInfoByRoleId(roleId, sysCode) {
  return request({
    url: '/permission/getPerInfoByRoleId/' + roleId,
    method: 'get',
    params: {
      sysCode
    }
  })
}
// 保存菜单信息
export function addPerInfo(info) {
  return request({
    url: '/permission/',
    method: 'Post',
    data: info
  })
}
// 获取菜单下所属的权限  详情/修改 功能
export function getPerInfoByMenuId(id) {
  return request({
    url: '/permission/' + id,
    method: 'get'
  })
}

//
export function updatePerInfo(info) {
  return request({
    url: '/permission/' + info.Id,
    method: 'put',
    data: info
  })
}
