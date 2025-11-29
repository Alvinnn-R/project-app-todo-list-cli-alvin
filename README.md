# 📝 Project CLI To-Do List

Aplikasi To-Do List berbasis Command Line Interface (CLI) menggunakan bahasa pemrograman Golang. Aplikasi ini membantu pengguna mengelola daftar tugas mereka secara lokal melalui terminal, dengan menyimpan data ke dalam file berformat JSON.

## 🎥 Video Demo

[![Video Demo Penjelasan Project](https://img.shields.io/badge/YouTube-Demo_Video-red?style=for-the-badge&logo=youtube)](https://youtu.be/WIvoIXmEnqw)

## 🎯 Fitur Utama

- ✅ **Menambahkan Tugas** - Tambahkan tugas baru dengan title dan description
- 📋 **Menampilkan Daftar Tugas** - Lihat semua tugas dalam format tabel yang rapi
- ✔️ **Menyelesaikan Tugas** - Tandai tugas sebagai selesai
- 🗑️ **Menghapus Tugas** - Hapus tugas yang tidak diperlukan
- 🔍 **Mencari Tugas** - Cari tugas berdasarkan kata kunci
- ✨ **Validasi Data** - Tidak ada title kosong dan tidak ada duplikat title
- 📊 **Tampilan Tabel** - Menggunakan `text/tabwriter` untuk tampilan yang rapi

## 📁 Struktur Project

```
mini-project/
├── main.go              # Entry point aplikasi dengan CLI flag
├── go.mod              # Module dependencies
├── data/
│   └── todos.json      # File penyimpanan data (auto-generated)
├── model/
│   ├── base.go         # Base model dengan timestamp
│   └── todo.go         # Model Todo
├── dto/
│   └── todo.go         # Data Transfer Objects
├── service/
│   └── todo.go         # Business logic layer
├── handler/
│   └── todo.go         # Handler layer untuk CLI
└── utils/
    └── file.go         # File operations utilities
```

## 🛠️ Teknologi yang Digunakan

- **Go 1.25.3**
- **Package flag** - untuk CLI command parsing
- **Package encoding/json** - untuk JSON processing
- **Package text/tabwriter** - untuk tampilan tabel
- **Package os** - untuk file handling

## 📋 Implementasi Ketentuan

### Ketentuan Utama

- ✅ **Operator**: Menggunakan operator perbandingan (==, !=, <=, >=) dan logika (&&, ||)
- ✅ **Variable**: Digunakan untuk menyimpan data, input, dan status
- ✅ **Function**: Lebih dari 3 fungsi (AddTodo, ListTodos, CompleteTodo, DeleteTodo, SearchTodos)
- ✅ **Array dan Slice**: Menggunakan slice untuk menyimpan daftar todos
- ✅ **Layout dan Formatting**: Menggunakan fmt.Printf, fmt.Sprintf, dan tabwriter
- ✅ **Error Handling**: Implementasi error handling di semua fungsi
- ✅ **JSON Processing**: Menggunakan encoding/json untuk read/write
- ✅ **File Handling**: Menggunakan os.ReadFile dan os.WriteFile
- ✅ **CLI**: Menggunakan package flag untuk command line arguments

### Ketentuan Tambahan

- ✅ **Package flag**: Menggunakan flag untuk --add, --list, --done, --delete, --search
- ✅ **Fitur pencarian**: Cari tugas berdasarkan keyword di title atau description
- ✅ **Validasi data**: Title tidak boleh kosong dan tidak boleh duplikat
- ✅ **Tampilan tabel**: Menggunakan text/tabwriter untuk tampilan rapi

## 🚀 Cara Menjalankan

### 1. Clone atau Download Project

```bash
cd mini-project
```

### 2. Menambahkan Todo Baru

```bash
go run main.go add --title "Belajar Golang" --desc "Mempelajari Go untuk project CLI"
```

### 3. Melihat Semua Todo

```bash
go run main.go list
```

### 4. Menandai Todo Sebagai Selesai

```bash
go run main.go done --id 1
```

### 5. Menghapus Todo

```bash
go run main.go delete --id 2
```

### 6. Mencari Todo

```bash
go run main.go search --keyword "Golang"
```

### 7. Melihat Bantuan

```bash
go run main.go help
```

Output:

```
📋 TODO LIST
====================================================================================================
ID    Title                          Description                              Status      Created At
---------- ------------------------------ ---------------------------------------- ---------- --------------------
1    Belajar Golang                 Mempelajari dasar-dasar Golang           ⏳ Pending  2025-11-29 10:30:15
2    Membuat Project                Membuat CLI Todo List                    ⏳ Pending  2025-11-29 10:31:20
3    Code Review                    Review kode dengan mentor                ⏳ Pending  2025-11-29 10:32:10
```

### Menyelesaikan Todo

```bash
go run main.go done --id 1
```

Output:

```
✅ Todo #1 marked as completed!
```

### Mencari Todo

```bash
go run main.go search --keyword "Golang"
```

Output:

```
🔍 Search results for 'Golang':
====================================================================================================
ID    Title                          Description                              Status
---------- ------------------------------ ---------------------------------------- ----------
1    Belajar Golang                 Mempelajari dasar-dasar Golang           ✅ Done
```

## 🔍 Validasi dan Error Handling

### Validasi Title Kosong

```bash
go run main.go add --title "" --desc "Test"
```

Output: `❌ Error: title is required and cannot be empty`

### Validasi Duplikat Title

```bash
go run main.go add --title "Belajar Golang" --desc "Duplikat"
```

Output: `❌ Error: todo with this title already exists`

### Validasi ID Tidak Valid

```bash
go run main.go done --id 0
```

Output: `❌ Error: invalid ID: must be greater than 0`

### Validasi Todo Tidak Ditemukan

```bash
go run main.go delete --id 999
```

Output: `❌ Error: todo not found`

## 📝 Struktur Data JSON

Data disimpan di `data/todos.json` dengan format:

```json
[
  {
    "id": 1,
    "created_at": "2025-11-29T10:30:15.123456Z",
    "updated_at": "2025-11-29T10:30:15.123456Z",
    "deleted_at": "0001-01-01T00:00:00Z",
    "title": "Belajar Golang",
    "description": "Mempelajari dasar-dasar Golang",
    "is_completed": false
  }
]
```

## 👨‍💻 Author

**Nama**: Alvin Rama S  
**Project**: CLI To-Do List Application  
**Tanggal**: 29 November 2025

## 📄 License

Project ini dibuat untuk keperluan pembelajaran Bootcamp Golang.

---

**Catatan**: Pastikan Go sudah terinstall di sistem Anda sebelum menjalankan aplikasi ini.
