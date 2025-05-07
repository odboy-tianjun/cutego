import request from '@/utils/request'

// 查询定时任务调度列表
export function listJob(query) {
  return request({
    url: '/api/v1/monitor/cronJob/list',
    method: 'get',
    params: query
  })
}

// 查询定时任务调度详细
export function getJob(jobId) {
  return request({
    url: '/api/v1/monitor/cronJob/' + jobId,
    method: 'get'
  })
}

// 新增定时任务调度
export function addJob(data) {
  return request({
    url: '/api/v1/monitor/cronJob/create',
    method: 'post',
    data: data
  })
}

// 修改定时任务调度
export function updateJob(data) {
  return request({
    url: '/api/v1/monitor/cronJob/modify',
    method: 'put',
    data: data
  })
}

// 删除定时任务调度
export function delJob(jobId, funcAlias) {
  return request({
    url: `/api/v1/monitor/cronJob/${jobId}/${funcAlias}`,
    method: 'delete'
  })
}

// 任务状态修改
export function changeJobStatus(jobId, funcAlias, status) {
  const data = {
    jobId,
    funcAlias,
    status
  }
  return request({
    url: '/api/v1/monitor/cronJob/changeStatus',
    method: 'put',
    data: data
  })
}


// 定时任务立即执行一次
export function runJob(jobId, funcAlias) {
  const data = {
    jobId,
    funcAlias
  }
  return request({
    url: '/api/v1/monitor/cronJob/run',
    method: 'put',
    data: data
  })
}
