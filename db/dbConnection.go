package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/labstack/gommon/log"

	_ "github.com/lib/pq" // postgres golang driver
	// "github.com/tkanos/gonfig"
)

type DBConfig struct {
	HostDB     string `json:"host"`
	PortDB     string `json:"port"`
	UserDB     string `json:"user"`
	PasswordDB string `json:"password"`
	NameDB     string `json:"dbname"`
}

func DB() *sql.DB {
	var configuration *DBConfig

	plan, err := ioutil.ReadFile("config/config.json")
	if err != nil {
		log.Error(err)
	}
	json.Unmarshal(plan, &configuration)
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", configuration.HostDB, configuration.PortDB, configuration.UserDB, configuration.PasswordDB, configuration.NameDB)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak dapat konek ke database : %s", err)
	}
	return result
}
