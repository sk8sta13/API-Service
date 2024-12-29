package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBdrive   string `mapstructure:"DB_DRIVE"`
	DBhost    string `mapstructure:"DB_HOST"`
	DBport    string `mapstructure:"DB_PORT"`
	DBuser    string `mapstructure:"DB_USER"`
	DBpass    string `mapstructure:"DB_PASS"`
	DBname    string `mapstructure:"DB_NAME"`
	Webport   string `mapstructure:"WEB_PORT"`
	JWTsecret string `mapstructure:"JWT_SECRET"`
	JWTttl    int    `mapstructure:"JWT_TTL"`
	TokenAuth *jwtauth.JWTAuth
}

func LoadConfigs(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	var cfg *conf
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTsecret), nil)

	return cfg, nil
}
