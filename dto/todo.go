package dto

// AddTodoRequest untuk menambahkan tugas baru
type AddTodoRequest struct {
	Title       string
	Description string
}

// UpdateTodoRequest untuk mengupdate tugas
type UpdateTodoRequest struct {
	Id          int
	Title       string
	Description string
}

// ListTodoResponse untuk menampilkan daftar tugas
type ListTodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
	CreatedAt   string `json:"created_at"`
}

// SearchTodoResponse untuk hasil pencarian tugas
type SearchTodoResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}
