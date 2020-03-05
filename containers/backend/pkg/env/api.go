package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// APIConfig - APIサーバのホストとポートのコンフィグ
type APIConfig struct {
	Host string `split_words:"true"`
	Port string `split_words:"true"`
}

// ReadAPIEnv - 指定したenvfileからAPIサーバに関する設定を読み込む
func ReadAPIEnv(envFile string) (APIConfig, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return APIConfig{}, err
	}
	var APICfg APIConfig
	err = envconfig.Process("API", &APICfg)
	if err != nil {
		return APIConfig{}, err
	}
	return APICfg, nil
}
