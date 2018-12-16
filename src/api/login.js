import request from '@/utils/request'

export function login(username, password, sysCode) {
  return request({
    url: '/authoritymanage/Login',
    method: 'post',
    data: {
      username,
      password,
      sysCode
    }
  })
}

export function getInfo(token) {
  return request({
    url: '/authoritymanage/GetUserInfo',
    method: 'get'
  })
}

export function logout() {
  return request({
    url: '/authoritymanage/Logout',
    method: 'post'
  })
}
