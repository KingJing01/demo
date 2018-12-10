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

export function getPerInfoBySysCode(sysCode) {
  return request({
    url: '/permission/getPerInfoBySysCode/' + sysCode,
    method: 'get'
  })
}
