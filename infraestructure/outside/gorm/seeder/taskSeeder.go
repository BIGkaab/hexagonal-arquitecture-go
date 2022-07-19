package seeder

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/config"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
	"github.com/labstack/gommon/log"
)

func Execute() {
	//runSeeders, _ := strconv.ParseBool(os.Getenv("RUN_SEEDERS"))
	runSeeders := true

	if !runSeeders {
		log.Warn("Seeders disabled")
	} else {
		dbInstance := config.ConnInstance()
		task := []entity.Task{
			{Name: "task 1", Description: "Lorem Ipsum is simply dummy text.", Punctuation: 5},
			{Name: "task 2", Description: "Lorem Ipsum is simply dummy text.", Punctuation: 8},
		}
		for _, c := range task {
			err := dbInstance.Create(&c).Error
			if err != nil {
				log.Fatalf("cannot seed task table: %v", err)
			}
		}
		log.Info("Executing Seeder")
	}
}
