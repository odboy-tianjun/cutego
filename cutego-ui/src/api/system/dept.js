import request from '@/utils/request'

// 查询部门列表
export function listDept(query) {
  return request({
    url: '/api/v1/dept/list',
    method: 'get',
    params: query
  })
}

// 查询部门列表（排除节点）
export function listDeptExcludeChild(deptId) {
  return request({
    url: '/api/v1/dept/list/exclude/' + deptId,
    method: 'get'
  })
}

// 查询部门详细
export function getDept(deptId) {
  return request({
    url: '/api/v1/dept/' + deptId,
    method: 'get'
  })
}

// 查询部门下拉树结构
export function deptTreeSelect() {
  return request({
    url: '/api/v1/dept/treeSelect',
    method: 'get'
  })
}

// 根据角色ID查询部门树结构
export function roleDeptTreeSelect(roleId) {
  return request({
    url: '/api/v1/dept/roleDeptTreeSelect/' + roleId,
    method: 'get'
  })
}

// 新增部门
export function addDept(data) {
  return request({
    url: '/api/v1/dept/create',
    method: 'post',
    data: data
  })
}

// 修改部门
export function updateDept(data) {
  return request({
    url: '/api/v1/dept/modify',
    method: 'put',
    data: data
  })
}

// 删除部门
export function delDept(deptId) {
  return request({
    url: '/api/v1/dept/' + deptId,
    method: 'delete'
  })
}
