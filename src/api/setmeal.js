/*  基本权限  用户套餐模块
 *  引用配置的api
 */
import request from '@/utils/request'

// 获取菜单列表信息
export function getSetMealList(info) {
  return request({
    url: '/setmeal',
    method: 'get',
    params: info
  })
}

export function addSetMealInfo(info) {
  return request({
    url: '/setmeal',
    method: 'post',
    data: info
  })
}

export function deleteSetMeal(ids) {
  return request({
    url: '/setmeal/' + ids,
    method: 'delete'
  })
}

export function updateSetMealInfo(info) {
  return request({
    url: '/setmeal/updateSetMealInfo',
    method: 'put',
    data: info
  })
}
