package migration

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/config"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
)

func Execute() {

	runMigrations, _ := strconv.ParseBool(os.Getenv("DATABASE_RUN_MIGRATIONS"))
	//runMigrations := true
	dropTableIfExists, _ := strconv.ParseBool(os.Getenv("DATABASE_DROP_TABLE_IF_EXISTS"))
	//dropTableIfExists := true

	if !runMigrations {
		log.Warn("Migrations disabled")
	} else {
		dbInstance := config.ConnInstance()

		if dropTableIfExists {
			//err := dbInstance.Migrator().DropTable(&models.Task{}).Error
			err := dbInstance.Migrator().DropTable(&entity.Task{})
			if err != nil {
				log.Panicf("cannot drop table: %v", err)
			}
		}
		//err := dbInstance.AutoMigrate(&models.Task{}).Error
		err := dbInstance.AutoMigrate(&entity.Task{})
		if err != nil {
			log.Panicf("cannot drop table: %v", err)
		}
		log.Info("Executing migrations")

	}

}
