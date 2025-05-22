package autoload

import "time"

type MysqlConfig struct {
	Enable       bool          `mapstructure:"enable"`
	Host         string        `mapstructure:"host"`
	Username     string        `mapstructure:"username"`
	Password     string        `mapstructure:"password"`
	Port         uint16        `mapstructure:"port"`
	Database     string        `mapstructure:"database"`
	Charset      string        `mapstructure:"charset"`
	TablePrefix  string        `mapstructure:"table_prefix"`
	MaxIdleConns int           `mapstructure:"max_idle_conns"`
	MaxOpenConns int           `mapstructure:"max_open_conns"`
	MaxLifetime  time.Duration `mapstructure:"max_lifetime"`
	LogLevel     int           `mapstructure:"log_level"`
	PrintSql     bool          `mapstructure:"print_sql"`
}

var Mysql = MysqlConfig{
	Enable:       true,
	Host:         "",
	Username:     "",
	Password:     "",
	Port:         3306,
	Database:     "",
	Charset:      "utf8mb4",
	TablePrefix:  "",
	MaxIdleConns: 10,
	MaxOpenConns: 100,
	MaxLifetime:  time.Hour,
	LogLevel:     4,
	PrintSql:     false,
}
