package handler

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
	"todo-list-cli/dto"
	"todo-list-cli/service"
)

type TodoHandler struct {
	TodoService service.TodoService
}

// NewTodoHandler membuat instance baru dari TodoHandler
func NewTodoHandler(todoService service.TodoService) TodoHandler {
	return TodoHandler{
		TodoService: todoService,
	}
}

// AddTodo handler untuk menambahkan tugas baru
func (th *TodoHandler) AddTodo(title, description string) {
	req := dto.AddTodoRequest{
		Title:       title,
		Description: description,
	}

	todo, err := th.TodoService.AddTodo(req)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Println("\nTodo added successfully!")
	fmt.Printf("ID: %d\n", todo.Id)
	fmt.Printf("Title: %s\n", todo.Title)
	fmt.Printf("Description: %s\n", todo.Description)
	fmt.Printf("Created at: %s\n", todo.CreatedAt.Format("2025-11-27 15:04:05"))
}

// ListTodos handler untuk menampilkan semua tugas dalam format tabel
func (th *TodoHandler) ListTodos() {
	todos, err := th.TodoService.ListTodos()
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	if len(todos) == 0 {
		fmt.Println("\nNo todos found. Add your first todo!")
		return
	}

	fmt.Println("\nTODO LIST")
	fmt.Println(strings.Repeat("=", 100))

	// Gunakan tabwriter untuk format tabel
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tDescription\tStatus\tCreated At")
	fmt.Fprintln(w, strings.Repeat("-", 5)+"\t"+strings.Repeat("-", 25)+"\t"+strings.Repeat("-", 35)+"\t"+strings.Repeat("-", 12)+"\t"+strings.Repeat("-", 20))

	for _, todo := range todos {
		status := "Pending"
		if todo.IsCompleted {
			status = "Done"
		}

		// Batasi panjang title dan description untuk tampilan tabel
		title := todo.Title
		if len(title) > 25 {
			title = title[:22] + "..."
		}

		desc := todo.Description
		if len(desc) > 35 {
			desc = desc[:32] + "..."
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n",
			todo.Id,
			title,
			desc,
			status,
			todo.CreatedAt,
		)
	}
	w.Flush()
	fmt.Println()
}

// CompleteTodo handler untuk menandai tugas sebagai selesai
func (th *TodoHandler) CompleteTodo(id int) {
	err := th.TodoService.CompleteTodo(id)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("\nTodo #%d marked as completed!\n", id)
}

// DeleteTodo handler untuk menghapus tugas
func (th *TodoHandler) DeleteTodo(id int) {
	err := th.TodoService.DeleteTodo(id)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	fmt.Printf("\nTodo #%d deleted successfully!\n", id)
}

// SearchTodos handler untuk mencari tugas berdasarkan keyword
func (th *TodoHandler) SearchTodos(keyword string) {
	results, err := th.TodoService.SearchTodos(keyword)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	if len(results) == 0 {
		fmt.Printf("\nNo todos found with keyword: '%s'\n", keyword)
		return
	}

	fmt.Printf("\nSearch results for '%s':\n", keyword)
	fmt.Println(strings.Repeat("=", 100))

	// Gunakan tabwriter untuk format tabel
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tTitle\tDescription\tStatus")
	fmt.Fprintln(w, strings.Repeat("-", 5)+"\t"+strings.Repeat("-", 25)+"\t"+strings.Repeat("-", 35)+"\t"+strings.Repeat("-", 12))

	for _, todo := range results {
		status := "Pending"
		if todo.IsCompleted {
			status = "Done"
		}

		// Batasi panjang title dan description
		title := todo.Title
		if len(title) > 25 {
			title = title[:22] + "..."
		}

		desc := todo.Description
		if len(desc) > 35 {
			desc = desc[:32] + "..."
		}

		fmt.Fprintf(w, "%d\t%s\t%s\t%s\n",
			todo.Id,
			title,
			desc,
			status,
		)
	}
	w.Flush()
	fmt.Println()
}
