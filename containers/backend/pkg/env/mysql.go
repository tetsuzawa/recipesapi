package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/tetsuzawa/voyageapi/containers/backend/pkg/mysql"
)

// ReadMysqlEnv - 指定したenvfileからMysqlに関する設定を読み込む
func ReadMysqlEnv(envFile string) (mysql.Config, error) {
	err := godotenv.Load(envFile)
	if err != nil {
		return mysql.Config{}, err
	}
	var MysqlCfg mysql.Config
	err = envconfig.Process("MYSQL", &MysqlCfg)
	if err != nil {
		return mysql.Config{}, err
	}
	return MysqlCfg, nil
}
