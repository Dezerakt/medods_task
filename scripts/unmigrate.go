package main

import (
	"github.com/fatih/color"
	"log"
	"medods_task/configs"
)

func init() {
	config := configs.NewConfigLoader()
	err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	log.Println(color.YellowString("Откат миграций..."))
	err := configs.DbObject.Migrator().DropTable(configs.TablesList...)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(color.GreenString("Откат миграций завершен"))
}
