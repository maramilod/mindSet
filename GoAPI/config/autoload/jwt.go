package autoload

import "time"

type JwtConfig struct {
	TTL time.Duration `mapstructure:"ttl"`
	RefreshTTL time.Duration `mapstructure:"refresh_ttl"`
	SecretKey  string        `mapstructure:"secret_key"`
}

var Jwt = JwtConfig{
	TTL:        7200 * 12 * 365,
	RefreshTTL: 7200 * 12 * 15,
	SecretKey:  "",
}
