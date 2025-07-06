package init

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type mysqlOption struct {
	Username string
	Password string
	Host     string
	Name     string
	Debug    bool
}

func NewMysqlOption(conf *viper.Viper) mysqlOption {
	return mysqlOption{
		Username: conf.GetString("mysql.username"),
		Password: conf.GetString("mysql.password"),
		Host:     conf.GetString("mysql.host"),
		Name:     conf.GetString("mysql.database"),
		Debug:    conf.GetBool("mysql.debug"),
	}
}

func NewMySQL(option mysqlOption) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", option.Username, option.Password, option.Host, option.Name)

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
