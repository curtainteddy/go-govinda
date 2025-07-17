# go-govinda 🎭

![Demo](demo.gif)

A straightforward CLI Todo app built in Go — no drama, just action. Like Govinda in his prime.

## Features
- ✅ **Add tasks**: Create new todo items with descriptive names
- 📋 **List tasks**: Display all tasks with completion status
- ✓ **Complete tasks**: Mark tasks as done with visual indicators
- 🗑️ **Remove tasks**: Delete tasks from your list
- 💾 **Persistent storage**: Tasks saved locally in JSON format
- 🎯 **Simple CLI**: Clean command interface with helpful error messages

## Installation

### Option 1: Build from source
```bash
git clone https://github.com/curtainteddy/go-govinda.git
cd go-govinda
go build -o govinda
```

### Option 2: Cross-platform builds
Use the provided build scripts:
- **Windows**: Run `build.bat` to create `govWin.exe`
- **Linux/Mac**: Run `build.sh` to create `govLin` and `govMac`

## Usage

### Basic Commands

```bash
# Add a new task
./govinda add "Buy groceries"
./govinda add "Learn Go programming"

# List all tasks
./govinda list

# Mark a task as complete (use task number from list)
./govinda complete 1

# Remove a task (use task number from list)
./govinda remove 2
```

### Example Session

```bash
$ ./govinda add "Buy groceries"
Task added successfully: Buy groceries

$ ./govinda add "Learn Go programming"
Task added successfully: Learn Go programming

$ ./govinda add "Write documentation"
Task added successfully: Write documentation

$ ./govinda list
[ ] 1: Buy groceries
[ ] 2: Learn Go programming
[ ] 3: Write documentation

$ ./govinda complete 1
Task marked as complete.

$ ./govinda complete 3
Task marked as complete.

$ ./govinda list
[✓] 1: Buy groceries
[ ] 2: Learn Go programming
[✓] 3: Write documentation

$ ./govinda remove 2
Task removed.

$ ./govinda list
[✓] 1: Buy groceries
[✓] 3: Write documentation
```

### Error Handling Examples

```bash
# Missing command
$ ./govinda
Please provide a command: add, list, complete, or remove

# Missing task description
$ ./govinda add
Please provide the task name. Usage: add "task description"

# Invalid task number
$ ./govinda complete abc
Invalid task number.

# Task not found
$ ./govinda complete 999
Task not found.
```

## Clean Code Principles Applied

This project demonstrates several clean code principles:

### 🎯 Single Responsibility Principle
- **`main.go`**: Handles command parsing and routing
- **`handlers.go`**: Contains business logic for each command
- **`storage.go`**: Manages file I/O operations
- **`task.go`**: Defines the Task data structure

### 📝 Meaningful Names
- Functions: `handleAdd()`, `loadTasks()`, `saveTasks()`, `getNextID()`
- Variables: `taskName`, `taskNum`, `newTask`, `fileName`
- Constants: Clear file naming (`tasks.json`)

### 🔧 Small, Focused Functions
- Each handler function does exactly one thing
- File operations are separated into dedicated functions
- ID generation logic is extracted into its own function

### 🚫 Avoiding Deep Nesting
- Clean switch statement for command routing
- Early returns to avoid deep if-else nesting
- Simple conditional structures throughout

### ⚠️ Graceful Error Handling
- File operation errors are caught and reported
- User input validation with helpful messages
- Proper error propagation from storage layer

### 📖 Readable and Consistent Code
- Consistent indentation and formatting
- Clear variable naming conventions
- Logical code organization and structure
- Helpful comments where needed

## File Structure

```
go-govinda/
├── main.go          # Entry point and command routing
├── handlers.go      # Business logic for each command
├── storage.go       # File I/O operations
├── task.go         # Task data structure
├── tasks.json      # Persistent storage (created automatically)
├── build.bat       # Windows build script
├── build.sh        # Linux/Mac build script
├── go.mod          # Go module definition
└── README.md       # This file
```

## Technical Details

- **Language**: Go 1.24.3
- **Storage**: JSON file (`tasks.json`)
- **Architecture**: Simple modular design with separation of concerns
- **Error Handling**: Comprehensive validation and user-friendly messages
- **Build**: Cross-platform support (Windows, Linux, macOS)
