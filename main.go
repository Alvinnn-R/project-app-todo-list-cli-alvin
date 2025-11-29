package main

import (
	"flag"
	"fmt"
	"os"
	"todo-list-cli/handler"
	"todo-list-cli/service"
)

func main() {
	// Inisialisasi service dan handler
	todoService := service.NewTodoService()
	todoHandler := handler.NewTodoHandler(todoService)

	// Define subcommands
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	searchCmd := flag.NewFlagSet("search", flag.ExitOnError)

	// Flags untuk add command
	addTitle := addCmd.String("title", "", "Title of the todo (required)")
	addDesc := addCmd.String("desc", "", "Description of the todo (optional)")

	// Flags untuk done command
	doneId := doneCmd.Int("id", 0, "ID of the todo to mark as done (required)")

	// Flags untuk delete command
	deleteId := deleteCmd.Int("id", 0, "ID of the todo to delete (required)")

	// Flags untuk search command
	searchKeyword := searchCmd.String("keyword", "", "Keyword to search (required)")

	// Cek apakah ada subcommand
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	// Parse subcommand
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		if *addTitle == "" {
			fmt.Println("Error: --title flag is required")
			addCmd.PrintDefaults()
			os.Exit(1)
		}
		todoHandler.AddTodo(*addTitle, *addDesc)

	case "list":
		listCmd.Parse(os.Args[2:])
		todoHandler.ListTodos()

	case "done":
		doneCmd.Parse(os.Args[2:])
		if *doneId <= 0 {
			fmt.Println("Error: --id flag is required and must be greater than 0")
			doneCmd.PrintDefaults()
			os.Exit(1)
		}
		todoHandler.CompleteTodo(*doneId)

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		if *deleteId <= 0 {
			fmt.Println("Error: --id flag is required and must be greater than 0")
			deleteCmd.PrintDefaults()
			os.Exit(1)
		}
		todoHandler.DeleteTodo(*deleteId)

	case "search":
		searchCmd.Parse(os.Args[2:])
		if *searchKeyword == "" {
			fmt.Println("Error: --keyword flag is required")
			searchCmd.PrintDefaults()
			os.Exit(1)
		}
		todoHandler.SearchTodos(*searchKeyword)

	case "help":
		printUsage()

	default:
		fmt.Printf("Unknown command: %s\n\n", os.Args[1])
		printUsage()
		os.Exit(1)
	}
}

// printUsage menampilkan panduan penggunaan aplikasi
func printUsage() {
	fmt.Println(`
==============================================================
|           📝 TODO LIST CLI APPLICATION 📝                  |
==============================================================

USAGE:
  todo-list-cli <command> [options]

COMMANDS:
  add       Add a new todo
  list      List all todos
  done      Mark a todo as completed
  delete    Delete a todo
  search    Search todos by keyword
  help      Show this help message

EXAMPLES:
  # Add a new todo
  go run main.go add --title "Buy groceries" --desc "Milk, bread, eggs"
  
  # List all todos
  go run main.go list
  
  # Mark todo as done
  go run main.go done --id 1
  
  # Delete a todo
  go run main.go delete --id 2
  
  # Search todos
  go run main.go search --keyword "groceries"

OPTIONS FOR EACH COMMAND:
  add:
    --title string    Title of the todo (required)
    --desc string     Description of the todo (optional)
  
  done:
    --id int          ID of the todo to mark as done (required)
  
  delete:
    --id int          ID of the todo to delete (required)
  
  search:
    --keyword string  Keyword to search in title or description (required)
`)
}
