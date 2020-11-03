package Base
// ConfInfluxDB 数据库配置
type ConfInfluxDB struct {
	Host string
	Token string
}

type InfuxDBClient struct {
	ConfInfluxDB
}