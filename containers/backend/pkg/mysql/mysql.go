package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Config TODO
type Config struct {
	User      string `split_words:"true"`
	Password  string `split_words:"true"`
	Protocol  string `split_words:"true"`
	Host      string `split_words:"true"`
	Port      string `split_words:"true"`
	DBName    string `split_words:"true"`
	Charset   string `split_words:"true"`
	ParseTime string `split_words:"true"`
	Loc       string `split_words:"true"`
}

func (c Config) build() Config {
	if c.User == "" {
		c.User = "root"
	}
	if c.Password == "" {
		c.Password = "root"
	}
	if c.Protocol == "" {
		c.Protocol = "tcp"
	}
	if c.Host == "" {
		c.Host = "127.0.0.1"
	}
	if c.Port == "" {
		c.Port = "3306"
	}
	if c.Charset == "" {
		c.Charset = "utf8mb4"
	}
	if c.ParseTime == "" {
		c.ParseTime = "True"
	}
	if c.Loc == "" {
		c.Loc = "Local"
	}
	return c
}

// Connect TODO
func Connect(c Config) (*gorm.DB, error) {
	c = c.build()

	const DBMS = "mysql"
	CONNECT := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		c.User,
		c.Password,
		c.Protocol,
		c.Host,
		c.Port,
		c.DBName,
		c.Charset,
		c.ParseTime,
		c.Loc,
	)

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		return nil, err
	}
	return db, nil
}
