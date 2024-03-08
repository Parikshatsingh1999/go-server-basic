package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	ID          uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `json:"-"`
	IsDone      bool           `json:"is_done" gorm:"column:is_done;default:false;not null;"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
}

var db *gorm.DB

func ConnectDB() {
	dsn := "user=postgres dbname=testdata sslmode=disable host=localhost port=5432 password=password"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// AutoMigrate will create the table if it does not exist
	db.AutoMigrate(&Todo{})

	// newRecord := Todo{Name: "first", Description: "test"}
	// db.Create(&newRecord)
	var data []Todo
	db.Find(&data)
	// router.Test = data

	fmt.Println("Connected to the PostgreSQL database using GORM! added")
}

func GetDb() *gorm.DB {
	return db
}

func GetAllTodos() ([]Todo, error) {
	var data []Todo
	result := db.Order("id").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func GetTodoById(id int) ([]Todo, error) {
	var data []Todo
	result := db.First(&data, id)
	if result.Error != nil {
		return []Todo{}, result.Error
	}
	return data, nil
}

func InsertTodo(data Todo) (Todo, error) {
	if len(data.Name) != 0 && len(data.Description) != 0 {
		result := db.Create(&Todo{Name: data.Name, Description: data.Description})
		// result := db.Save(&data), created when Id is not given, else will update
		if result.Error != nil {
			return Todo{}, result.Error
		}
		return data, nil
	}
	return Todo{}, errors.New("name and description cannot be empty")
}

func DeleteTodo(id int) (Todo, error) {
	var data Todo
	result := db.Delete(&data, id)
	if result.Error != nil {
		return Todo{}, result.Error
	}
	return data, nil
}

func UpdateTodo(updates map[string]interface{}, id int) (Todo, error) {
	var newTodo Todo
	result := db.Model(&newTodo).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return Todo{}, result.Error
	}
	return newTodo, nil
}
