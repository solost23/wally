package configs

type ServerConfig struct {
	Version          string        `mapstructure:"version"`
	DebugMode        string        `mapstructure:"debug_mode"`
	TimeLocation     string        `mapstructure:"time_location"`
	Addr             string        `mapstructure:"addr"`
	PrometheusEnable bool          `mapstructure:"prometheus_enable"`
	ConfigPath       string        `mapstructure:"config_path"`
	MySQLConfig      MySQLConf     `mapstructure:"mysql"`
	RedisConfig      RedisConf     `mapstructure:"redis"`
	JWTConfig        JWTConf       `mapstructure:"jwt"`
	Zinc             Zinc          `mapstructure:"zinc"`
	ConsulConfig     ConsulConf    `mapstructure:"consul"`
	StaticOSS        StaticOSSConf `mapstructure:"static-oss"`
	PushSrvConfig    PushSrvConf   `mapstructure:"push"`
	OSSSrvConfig     OSSSrvConf    `mapstructure:"oss"`
}

type MySQLConf struct {
	Dsn             string `mapstructure:"dsn"`
	MaxOpenConn     int    `mapstructure:"max_open_conn"`
	MaxIdleConn     int    `mapstructure:"max_idle_conn"`
	MaxConnLifeTime int    `mapstructure:"max_conn_life_time"`
}

type RedisConf struct {
	Addr string `mapstructure:"addr"`
}

type JWTConf struct {
	Key      string `mapstructure:"key"`
	Duration int64  `mapstructure:"duration"`
}

type Zinc struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type ConsulConf struct {
	Host string `mapstructure:"host"`
	Port int    `mapstructure:"port"`
}

type StaticOSSConf struct {
	Domain string `mapstructure:"domain"`
}

type PushSrvConf struct {
	Name string `mapstructure:"name"`
}

type OSSSrvConf struct {
	Name string `mapstructure:"name"`
}
