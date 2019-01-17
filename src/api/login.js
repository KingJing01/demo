import request from '@/utils/request'

export function login(username, password, sysCode) {
  return request({
    url: '/authoritymanage/authLogin',
    method: 'post',
    data: {
      username,
      password
    }
  })
}

export function getInfo(token) {
  return request({
    url: '/authoritymanage/GetUserPermission',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/authoritymanage/Logout',
    method: 'post'
  })
}
