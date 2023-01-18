package job

// AliasFuncMap 定时任务: 别名与方法的映射
var AliasFuncMap = make(map[string]func())

// 任务注册
func init() {
	AliasFuncMap["test1"] = TestJob1
}
