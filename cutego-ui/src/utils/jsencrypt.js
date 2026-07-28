// 前端加密/解密工具
// 注意：前端加密仅用于传输层之外的额外混淆，不能替代 HTTPS
// 敏感操作应依赖后端安全机制

// 加密
export function encrypt(txt) {
  // 使用内置的 btoa 进行基础编码，实际加密由后端处理
  try {
    return btoa(encodeURIComponent(txt))
  } catch (e) {
    return txt
  }
}

// 解密
export function decrypt(txt) {
  try {
    return decodeURIComponent(atob(txt))
  } catch (e) {
    return txt
  }
}