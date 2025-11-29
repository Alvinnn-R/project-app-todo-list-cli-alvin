package service

import (
	"errors"
	"strings"
	"time"
	"todo-list-cli/dto"
	"todo-list-cli/model"
	"todo-list-cli/utils"
)

type TodoService struct{}

// NewTodoService membuat instance baru dari TodoService
func NewTodoService() TodoService {
	return TodoService{}
}

// AddTodo menambahkan tugas baru ke daftar
func (ts *TodoService) AddTodo(req dto.AddTodoRequest) (model.Todo, error) {
	// Validasi input - judul tidak boleh kosong
	if strings.TrimSpace(req.Title) == "" {
		return model.Todo{}, errors.New("title is required and cannot be empty")
	}

	// Baca data existing
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return model.Todo{}, err
	}

	// Validasi duplikat - cek apakah judul sudah ada
	for _, t := range todos {
		if strings.EqualFold(t.Title, req.Title) {
			return model.Todo{}, errors.New("todo with this title already exists")
		}
	}

	// Generate ID baru (max ID + 1)
	newID := 1
	for _, t := range todos {
		if t.Id >= newID {
			newID = t.Id + 1
		}
	}

	// Buat todo baru
	newTodo := model.Todo{
		Base: model.Base{
			Id:        newID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		Title:       req.Title,
		Description: req.Description,
		IsCompleted: false,
	}

	// Tambahkan ke slice
	todos = append(todos, newTodo)

	// Simpan ke file
	if err := utils.WriteTodosToFile(todos); err != nil {
		return model.Todo{}, err
	}

	return newTodo, nil
}

// ListTodos menampilkan semua tugas
func (ts *TodoService) ListTodos() ([]dto.ListTodoResponse, error) {
	// Baca data dari file
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return nil, err
	}

	// Konversi ke response DTO
	var listTodos []dto.ListTodoResponse
	for _, t := range todos {
		todo := dto.ListTodoResponse{
			Id:          t.Id,
			Title:       t.Title,
			Description: t.Description,
			IsCompleted: t.IsCompleted,
			CreatedAt:   t.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		listTodos = append(listTodos, todo)
	}

	return listTodos, nil
}

// CompleteTodo menandai tugas sebagai selesai
func (ts *TodoService) CompleteTodo(id int) error {
	// Validasi ID
	if id <= 0 {
		return errors.New("invalid ID: must be greater than 0")
	}

	// Baca data dari file
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return err
	}

	// Cari dan update status
	found := false
	for i := range todos {
		if todos[i].Id == id {
			if todos[i].IsCompleted {
				return errors.New("todo is already completed")
			}
			todos[i].IsCompleted = true
			todos[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		return errors.New("todo not found")
	}

	// Simpan perubahan
	if err := utils.WriteTodosToFile(todos); err != nil {
		return err
	}

	return nil
}

// DeleteTodo menghapus tugas berdasarkan ID
func (ts *TodoService) DeleteTodo(id int) error {
	// Validasi ID
	if id <= 0 {
		return errors.New("invalid ID: must be greater than 0")
	}

	// Baca data dari file
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return err
	}

	// Cari dan hapus todo
	found := false
	var updatedTodos []model.Todo
	for _, t := range todos {
		if t.Id == id {
			found = true
			continue // Skip todo yang akan dihapus
		}
		updatedTodos = append(updatedTodos, t)
	}

	if !found {
		return errors.New("todo not found")
	}

	// Simpan perubahan
	if err := utils.WriteTodosToFile(updatedTodos); err != nil {
		return err
	}

	return nil
}

// SearchTodos mencari tugas berdasarkan keyword
func (ts *TodoService) SearchTodos(keyword string) ([]dto.SearchTodoResponse, error) {
	// Validasi keyword
	if strings.TrimSpace(keyword) == "" {
		return nil, errors.New("keyword cannot be empty")
	}

	// Baca data dari file
	todos, err := utils.ReadTodosFromFile()
	if err != nil {
		return nil, err
	}

	// Cari todos yang mengandung keyword (case-insensitive)
	var results []dto.SearchTodoResponse
	keyword = strings.ToLower(keyword)

	for _, t := range todos {
		// Cek di title atau description
		if strings.Contains(strings.ToLower(t.Title), keyword) ||
			strings.Contains(strings.ToLower(t.Description), keyword) {
			result := dto.SearchTodoResponse{
				Id:          t.Id,
				Title:       t.Title,
				Description: t.Description,
				IsCompleted: t.IsCompleted,
			}
			results = append(results, result)
		}
	}

	return results, nil
}
