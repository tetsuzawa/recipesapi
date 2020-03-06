package env

import "github.com/joho/godotenv"

// ReadEnv - 指定したファイル名から環境変数を読み込む
func ReadEnv(envFile string) error {
	return godotenv.Load(envFile)
}
