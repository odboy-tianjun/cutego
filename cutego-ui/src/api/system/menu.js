import request from '@/utils/request'

// 查询菜单列表
export function listMenu(query) {
  return request({
    url: '/api/v1/menu/list',
    method: 'get',
    params: query
  })
}

// 查询菜单详细
export function getMenu(menuId) {
  return request({
    url: '/api/v1/menu/' + menuId,
    method: 'get'
  })
}

// 查询菜单下拉树结构
export function menuTreeSelect() {
  return request({
    url: '/api/v1/menu/treeSelect',
    method: 'get'
  })
}

// 根据角色ID查询菜单下拉树结构
export function roleMenuTreeSelect(roleId) {
  return request({
    url: '/api/v1/menu/roleMenuTreeSelect/' + roleId,
    method: 'get'
  })
}

// 新增菜单
export function addMenu(data) {
  data.isCache = parseInt(data.isCache)
  data.isFrame = parseInt(data.isFrame)
  return request({
    url: '/api/v1/menu/create',
    method: 'post',
    data: data
  })
}

// 修改菜单
export function updateMenu(data) {
  data.isCache = parseInt(data.isCache)
  data.isFrame = parseInt(data.isFrame)
  return request({
    url: '/api/v1/menu/modify',
    method: 'put',
    data: data
  })
}

// 删除菜单
export function delMenu(menuId) {
  return request({
    url: '/api/v1/menu/' + menuId,
    method: 'delete'
  })
}
