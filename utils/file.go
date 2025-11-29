package utils

import (
	"encoding/json"
	"errors"
	"os"
	"todo-list-cli/model"
)

const TodosFilePath = "data/todos.json"

// EnsureTodosFile memastikan direktori dan file todos.json ada
func EnsureTodosFile() error {
	_, err := os.Stat(TodosFilePath)
	if errors.Is(err, os.ErrNotExist) {
		// Buat direktori data jika belum ada
		if err := os.MkdirAll("data", 0755); err != nil {
			return err
		}
		// Buat file dengan array kosong
		return os.WriteFile(TodosFilePath, []byte("[]"), 0644)
	}
	return nil
}

// ReadTodosFromFile membaca semua data todos dari file JSON
func ReadTodosFromFile() ([]model.Todo, error) {
	if err := EnsureTodosFile(); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(TodosFilePath)
	if err != nil {
		return nil, err
	}

	var todos []model.Todo
	if err := json.Unmarshal(bytes, &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

// WriteTodosToFile menulis semua data todos ke file JSON
func WriteTodosToFile(todos []model.Todo) error {
	bytes, err := json.MarshalIndent(todos, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(TodosFilePath, bytes, 0644)
}
