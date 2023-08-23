package model

import (
	"fmt"
	"time"
)

type Todo struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

func GetAllTodos() ([]Todo, error) {
	var todos []Todo
	tx := db.Find(&todos).Order("created_at desc")

	if tx.Error != nil {
		return []Todo{}, tx.Error
	}

	return todos, tx.Error
}

func CreateTodo(todo Todo) (Todo, error) {
	tx := db.Create(&todo)

	return todo, tx.Error
}

func DeleteTodo(id int) error {
	tx := db.Unscoped().Delete(&Todo{}, id)

	return tx.Error
}

func UpdateTodo(id int, todo Todo) (Todo, error) {
	fmt.Println("completed: ", todo)
	tx := db.Model(&todo).Select("*").Updates(&todo)

	if tx.Error != nil {
		fmt.Println("error: ", tx.Error)
		return Todo{}, tx.Error
	}

	return todo, nil
}

func ClearTodos(completed bool) error {
	tx := db.Delete(&[]Todo{}, "completed = ?", completed)

	return tx.Error
}
