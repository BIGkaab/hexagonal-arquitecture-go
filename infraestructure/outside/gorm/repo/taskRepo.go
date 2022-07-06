package repo

import (
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/config"
	"github.com/BIGKaab/hexagonal-arquitecture-go/infraestructure/outside/gorm/entity"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type TaskRepo struct {
	db *gorm.DB
}

func (repo *TaskRepo) OutGetAllTasks() ([]entity.Task, error) {
	var tasks []entity.Task
	if err := repo.db.Find(&tasks).Error; err != nil {
		log.Error(err)
		return nil, err
	}
	return tasks, nil
}

func (repo *TaskRepo) OutAddTask(task *entity.Task) (*entity.Task, error) {

	if err := repo.db.Model(entity.Task{}).Create(&task).Error; err != nil {
		log.Error(err)
		return task, err
	}
	return task, nil
}

func (repo *TaskRepo) OutFindTaskById(ID int) (entity.Task, error) {
	var task entity.Task
	if err := repo.db.Find(&task, ID).Error; err != nil {
		log.Error(err)
		return task, err
	}
	return task, nil
}

func (repo *TaskRepo) OutInUpdateTask(ID int, task *entity.Task) (*entity.Task, error) {
	err := repo.db.Model(&entity.Task{}).Where("id = ?", ID).UpdateColumns(
		map[string]interface{}{
			"name":        task.Name,
			"description": task.Description,
			"punctuation": task.Punctuation,
		},
	).Error
	if err != nil {
		log.Error(err)
		return task, err
	}
	task.ID = ID
	return task, nil
}

func (repo *TaskRepo) OutDeleteTask(ID int) error {
	if err := repo.db.Delete(&entity.Task{}, ID).Error; err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func NewTaskRepo() *TaskRepo {
	return &TaskRepo{
		db: config.ConnInstance(),
	}
}
