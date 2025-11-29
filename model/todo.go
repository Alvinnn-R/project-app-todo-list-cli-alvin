package model

// Todo struct untuk menyimpan data tugas
type Todo struct {
	Base
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
