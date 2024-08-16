package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func InitDatabase() (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASS"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)

	db, err := sql.Open(viper.GetString("DB_DRIVER"), dsn)
	if err != nil {
		return nil, err
	}
	// defer db.Close()

	return db, err
}
