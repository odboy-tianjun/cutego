package config

type ApplicationCoreStruct struct {
	CuteGoConfig CuteGoConfig `yaml:"cutego"`
}

// CuteGoConfig 总配置
type CuteGoConfig struct {
	// 默认激活dev配置
	Active string `yaml:"active" default:"dev"`
	// 开启演示模式
	DemoMode bool           `yaml:"demo-mode"`
	Mail     MailConfig     `yaml:"mail"`
	TaskPool TaskPoolConfig `yaml:"task-pool"`
	Captcha  CaptchaConfig  `yaml:"captcha"`
	File     FileConfig     `yaml:"file"`
}

// MailConfig 邮件
type MailConfig struct {
	// 服务地址
	Host string `yaml:"host"`
	// 服务端口
	Port int `yaml:"port"`
	// 用户名
	Username string `yaml:"username"`
	// 密码
	Password string `yaml:"password"`
	// 默认编码
	DefaultEncoding string `yaml:"default-encoding"`
}

// TaskPoolConfig 线程池
type TaskPoolConfig struct {
	// 核心线程池大小
	CorePoolSize int `yaml:"core-pool-size"`
	// 最大线程数(尽可能的大)
	MaxPoolSize int `yaml:"max-pool-size"`
	// 活跃时间(单位: s)
	KeepAliveSeconds int `yaml:"keep-alive-seconds"`
	// 队列容量
	QueueCapacity int `yaml:"queue-capacity"`
}

// CaptchaConfig 验证码有效时间(单位: s)
type CaptchaConfig struct {
	// 邮箱
	Email int `yaml:"email"`
	// 手机短信
	Sms int `yaml:"sms"`
}

// FileConfig 文件上传
type FileConfig struct {
	// 文件大小(单位: mb)
	FileMaxSize int `yaml:"file-max-size"`
	// 头像大小(单位: mb)
	AvatarMaxSize int      `yaml:"avatar-max-size"`
	Mac           FilePath `yaml:"mac"`
	Linux         FilePath `yaml:"linux"`
	Windows       FilePath `yaml:"windows"`
}
type FilePath struct {
	Path   string `yaml:"path"`
	Avatar string `yaml:"avatar"`
	Logs   string `yaml:"logs"`
}
