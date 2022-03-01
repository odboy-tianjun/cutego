```go
id := c.Query("id") // 查询请求URL后面拼接的参数

name := c.PostForm("name") // 从表单中查询参数

uuid := c.Param("uuid") // 取得URL中参数

s, _ := c.Get("current_manager") // 从用户上下文读取值    

page := c.DefaultQuery("page", "0") // 查询请求URL后面的参数，如果没有填写默认值
```