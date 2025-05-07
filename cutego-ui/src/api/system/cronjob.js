import request from '@/utils/request'

// 查询定时任务列表
export function listCronjob(query) {
  return request({
    url: '/api/v1/cronjob/list',
    method: 'get',
    params: query
  })
}

// 新增定时任务
export function addCronjob(data) {
  return request({
    url: '/api/v1/cronjob/create',
    method: 'post',
    data: data
  })
}

// 修改定时任务
export function updateCronjob(data) {
  return request({
    url: '/api/v1/cronjob/modify',
    method: 'put',
    data: data
  })
}

// 删除定时任务
export function delCronjob(deptId) {
  return request({
    url: '/api/v1/cronjob/' + deptId,
    method: 'delete'
  })
}
