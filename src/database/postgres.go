package database

import (
	"database/sql"
	"fmt"
	"os"
	"sync"

	"gopkg.in/yaml.v2"

	_ "github.com/lib/pq"
)

var once sync.Once

var (
	instance *sql.DB
)

type DbConfig struct {
	Database struct {
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
		Ip       string `yaml:"ip"`
		Port     string `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func GetDb() *sql.DB {

	once.Do(func() {
		f, err := os.Open("config.yml")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		defer f.Close()
		var cfg DbConfig
		decoder := yaml.NewDecoder(f)
		err = decoder.Decode(&cfg)
		if err != nil {
			os.Exit(-1)
		}
		psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.Database.Ip, cfg.Database.Port, cfg.Database.Username, cfg.Database.Password, cfg.Database.Name)
		instance, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(-1)
		}
		var version string
		instance.QueryRow("SELECT VERSION()").Scan(&version)
	})
	return instance
}
