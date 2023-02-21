package database

import (
	"database/sql"
	"fmt"
	"jump-backend-interview/src/config"
	"log"
	"os"
	"sync"

	_ "github.com/lib/pq"
)

var once sync.Once

var (
	instance *sql.DB
)

func GetDb() *sql.DB {

	once.Do(func() {
		cfg, err := config.GetConfig()
		if err != nil {
			log.Default().Println(err.Error())
			os.Exit(-1)
		}
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Ip, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
		instance, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Default().Println(err.Error())
			os.Exit(-1)
		}
		var version string
		instance.QueryRow("SELECT VERSION()").Scan(&version)
	})
	return instance
}
