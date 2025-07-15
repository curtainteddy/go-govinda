package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    args := os.Args

    if len(args) < 2 {
        fmt.Println("Please provide a command: add, list, complete, or remove")
        return
    }

    command := args[1]

    switch command {
    case "add":
        if len(args) < 3 {
            fmt.Println("Please provide the task name. Usage: add \"task description\"")
            return
        }
        taskName := args[2]
        handleAdd(taskName)

    case "list":
        handleList()

    case "complete":
        if len(args) < 3 {
            fmt.Println("Please provide the task number to mark as complete. Usage: complete 1")
            return
        }
        taskNum, err := strconv.Atoi(args[2])
        if err != nil {
            fmt.Println("Invalid task number.")
            return
        }
        handleComplete(taskNum)

    case "remove":
        if len(args) < 3 {
            fmt.Println("Please provide the task number to remove. Usage: remove 1")
            return
        }
        taskNum, err := strconv.Atoi(args[2])
        if err != nil {
            fmt.Println("Invalid task number.")
            return
        }
        handleRemove(taskNum)

    default:
        fmt.Println("Unknown command. Available commands: add, list, complete, remove")
    }
}
