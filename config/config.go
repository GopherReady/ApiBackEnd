package config

// Config 组合全部配置模型
type Config struct {
	Server Server `mapstructure:"server"`
	Mysql  Mysql  `mapstructure:"mysql"`
	Jwt    Jwt    `mapstructure:"jwt"`
}

// Server 服务启动端口号配置
type Server struct {
	Addr string `mapstructure:"addr"`
}

// Mysql MySQL数据源配置
type Mysql struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Url      string `mapstructure:"url"`
}

// Jwt 用户认证配置
type Jwt struct {
	SigningKey string `mapstructure:"signingKey"`
}
