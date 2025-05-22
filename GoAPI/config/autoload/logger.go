package autoload

type DivisionTime struct {
	MaxAge       int `mapstructure:"max_age"`       
	RotationTime int `mapstructure:"rotation_time"` 
}

type DivisionSize struct {
	MaxSize    int  `mapstructure:"max_size"`   
	MaxBackups int  `mapstructure:"max_backups"` 
	MaxAge     int  `mapstructure:"max_age"`     
	Compress   bool `mapstructure:"compress"`    
}

type LoggerConfig struct {
	DefaultDivision string       `mapstructure:"default_division"`
	Filename        string       `mapstructure:"file_name"`
	DivisionTime    DivisionTime `mapstructure:"division_time"`
	DivisionSize    DivisionSize `mapstructure:"division_size"`
}

var Logger = LoggerConfig{
	DefaultDivision: "time", 
	Filename:        "sys.log",
	DivisionTime: DivisionTime{
		MaxAge:       15,
		RotationTime: 24,
	},
	DivisionSize: DivisionSize{
		MaxSize:    2,
		MaxBackups: 2,
		MaxAge:     15,
		Compress:   false,
	},
}
