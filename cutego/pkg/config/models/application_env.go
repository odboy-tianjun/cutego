package config

type ApplicationEnvStruct struct {
	Server     ServerConfig     `yaml:"server"`
	DataSource DataSourceConfig `yaml:"datasource"`
	Redis      RedisConfig      `yaml:"redis"`
	Login      LoginConfig      `yaml:"login"`
	Jwt        JwtConfig        `yaml:"jwt"`
	Logger     LoggerConfig     `yaml:"logger"`
}

// ServerConfig web服务
type ServerConfig struct {
	// Running in "debug" mode. Switch to "release" mode in production
	RunMode string `yaml:"run-mode"`
	// web服务监听端口(生产的端口有可能不一样,所以根据环境隔离开)
	Port int `yaml:"port"`
}

// DataSourceConfig 数据源
type DataSourceConfig struct {
	// 数据库类型
	DbType string `yaml:"db-type" default:"mysql"`
	// 服务地址
	Host string `yaml:"host"`
	// 服务端口
	Port int `yaml:"port"`
	// 用户名称
	Username string `yaml:"username"`
	// 用户密码
	Password string `yaml:"password"`
	// 数据库名称
	Database string `yaml:"database"`
	// 编码
	Charset string `yaml:"charset"`
	// 空闲时的最大连接数
	MaxIdleSize int `yaml:"max-idle-size"`
	// 数据库的最大打开连接数
	MaxOpenSize int `yaml:"max-open-size"`
}

// RedisConfig 缓存
type RedisConfig struct {
	// 数据库索引
	Database int `yaml:"database"`
	// 服务地址
	Host string `yaml:"host"`
	// 服务端口
	Port int `yaml:"port"`
	// 服务密码
	Password string `yaml:"password"`
	// 连接超时时间
	Timeout int             `yaml:"timeout"`
	Pool    RedisPoolConfig `yaml:"pool"`
}

// RedisPoolConfig redis连接池配置
type RedisPoolConfig struct {
	// 连接池最大连接数（使用负值表示没有限制, 最佳配置为cpu核数+1）
	MaxActive int `yaml:"max-active"`
	// 连接池中的最大空闲连接
	MaxIdle int `yaml:"max-idle"`
	// 连接池最大阻塞等待时间（使用负值表示没有限制）
	MaxWait int `yaml:"max-wait"`
}

// LoginConfig 登录相关
type LoginConfig struct {
	// 是否限制单用户登录
	Single bool `yaml:"single"`
}

// JwtConfig jwt参数
type JwtConfig struct {
	// 请求头前缀
	Header string `yaml:"header" json:"header,omitempty"`
	// 令牌前缀
	TokenStartWith string `yaml:"token-start-with" json:"token_start_with,omitempty"`
	// 加密密钥
	TokenSecret string `yaml:"token-secret" json:"token_secret,omitempty"`
	// 令牌过期时间 此处单位: h , 默认1小时
	TokenExpired int `yaml:"token-expired" json:"token_expired,omitempty"`
	// token在cookie中的别称
	CookieKey string `yaml:"cookie-key"`
	// 在线用户key
	OnlineKey string `yaml:"online-key" json:"online_key,omitempty"`
	// token 续期检查时间范围（默认30分钟, 单位毫秒）, 在token即将过期的一段时间内用户操作了, 则给用户的token续期
	Detect int `yaml:"detect" json:"detect,omitempty"`
	// 续期时间范围, 默认1小时, 单位毫秒
	Renew int `yaml:"renew" json:"renew,omitempty"`
}

// LoggerConfig 日志文件
type LoggerConfig struct {
	// 最大保存时间(单位: d)
	MaxSaveAge int `yaml:"max-save-age"`
	// 日志切割时间间隔(单位: d)
	RotationTime int `yaml:"rotation-time"`
	// 日志级别
	Level string `yaml:"level"`
}
