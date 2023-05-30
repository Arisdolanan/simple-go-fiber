package repository

import (
	"errors"
	"gorm.io/gorm"
	"simple-go-fiber-crud/database"
	"simple-go-fiber-crud/models"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository() *TodoRepository {
	return &TodoRepository{database.DBConn}
}

func (u TodoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	u.db.Find(&todos)
	return todos, nil
}

func (u TodoRepository) FindById(id int) ([]models.Todo, error) {
	var todos []models.Todo
	err := u.db.Where("id = ?", id).Find(&todos).Error
	if errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return todos, err
	}
	return todos, nil
}

func (u TodoRepository) Create(todo models.Todo) *models.Todo {
	err := u.db.Create(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil
	}
	return &todo
}

func (u TodoRepository) Update(id int, updated models.Todo) *models.Todo {
	todo := new(models.Todo)
	err := u.db.Find(&todo, id).First(&todo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil
	}
	todo.Title = updated.Title
	todo.Completed = updated.Completed
	u.db.Save(&todo)
	return todo
}

func (u TodoRepository) Delete(id int) (*models.Todo, error) {
	var todos []models.Todo
	err := u.db.First(&todos, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, err
	}
	u.db.Delete(&todos)
	result, err := u.Delete(id)
	if errors.Is(err, gorm.ErrRecordNotFound) && err != nil {
		return nil, err
	}
	return result, nil
}
