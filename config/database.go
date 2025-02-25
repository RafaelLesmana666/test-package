package config

import (
	"fmt"
	"log"
	"os"

	jamethelper "github.com/RafaelLesmana666/jamet-helper"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var H23 *gorm.DB
var H1 *gorm.DB
var Jamet *jamethelper.Jamet

func Open() {

	check := godotenv.Load("./config/.env")
	if check != nil {
		log.Fatal("Error loading .env file") // Or handle more gracefully
	}

	var err error

	configH23 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("H23_USERNAME"), os.Getenv("H23_PASSWORD"), os.Getenv("H23_HOST"), os.Getenv("H23_PORT"), os.Getenv("H23_DB"))
	H23, err = gorm.Open(mysql.Open(configH23), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	configH1 := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("H1_USERNAME"), os.Getenv("H1_PASSWORD"), os.Getenv("H1_HOST"), os.Getenv("H1_PORT"), os.Getenv("H1_DB"))
	H1, err = gorm.Open(mysql.Open(configH1), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Jamet = jamethelper.NewJamet(map[string]*gorm.DB{
		"": H23,
		"H1":  H1,
	})

}
