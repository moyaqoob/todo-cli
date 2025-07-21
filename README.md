# 📝 Todo CLI App (Go)

A simple Command Line Interface (CLI) Todo application written in Go. Add, list, mark, and delete your tasks directly from the terminal.

---

## 🚀 Features

- Add todos
- List current todos
- Mark todos as completed
- Delete todos

---

## 📦 Requirements

- Go 1.18 or later

---

## 🛠️ Installation & Running

### 1. Clone the repository

```bash
git clone https://github.com/your-username/todo-cli.git
cd todo-cli
```

### 2. Clone the repository
```bash
go run main.go
```

### 3. Example Usage

```bash
1. Add Todo
2. List Todos
3. Mark Todo as Completed
4. Delete Todo
5. Exit
Enter your choice:
```
### 4. Example Usage

```bash

todo-cli/
│
├── main.go        # Entry point of the CLI application
├── todo.go        # Contains Todo struct and business logic
└── README.md
```

🧠 Notes
All data is stored in-memory (no persistence). Exiting the app clears all todos.

Extendable for file/database persistence or sync with an API.