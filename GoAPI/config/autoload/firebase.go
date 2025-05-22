package autoload

type FirebaseConfig struct {
	Enable            bool   `mapstructure:"enable"`
	ProjectId         string `mapstructure:"project_id"`
	PrivateKeyId      int64  `mapstructure:"private_key_id"`
	ServiceAccount    string `mapstructure:"service_account"`
	ServiceAccountKey string `mapstructure:"service_account_key"`
}

var Firebase = FirebaseConfig{
	Enable:            true,
	ProjectId:         "",
	PrivateKeyId:      539554129699,
	ServiceAccount:    "",
	ServiceAccountKey: "",
}
