import request from '@/utils/request'

// 登录方法
export function login(username, password, code, uuid) {
  const data = {
    username,
    password,
    code,
    uuid
  }
  return request({
    url: '/api/v1/login',
    method: 'post',
    data: data
  })
}

// 获取用户详细信息
export function getInfo() {
  return request({
    url: '/api/v1/getInfo',
    method: 'get'
  })
}

// 退出方法
export function logout() {
  return request({
    url: '/api/v1/logout',
    method: 'post'
  })
}

// 获取验证码
export function getCodeImg() {
  return request({
    url: '/captchaImage',
    method: 'get'
  })
}

// 查询用户登录记录列表
export function getLoginHistory(query) {
  return request({
    url: '/api/v1/getLoginHistory',
    method: 'get',
    params: query
  })
}
