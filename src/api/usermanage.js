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

export function getUserInfo(info) {
  return request({
    url: '/tenant/' + info.tenId,
    method: 'get',
    params: {
      sysCode: info.sysCode
    }
  })
}
