/*  基本权限  用户套餐模块
 *  引用配置的api
 */
import request from '@/utils/request'

// 获取套餐列表信息
export function getSetMealList(info) {
  return request({
    url: '/setmeal',
    method: 'get',
    params: info
  })
}
// 新增套餐信息
export function addSetMealInfo(info) {
  return request({
    url: '/setmeal',
    method: 'post',
    data: {
      'setMealName': info.setMealName,
      'setMealCode': info.setMealCode,
      'perId': info.perId,
      'perName': info.perName,
      'sysCode': info.sysCode
    }
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

export function getSetMealRadio(info) {
  return request({
    url: '/setmeal/getSetMealRadio',
    method: 'post',
    data: {
      sysCodes: info
    }
  })
}
