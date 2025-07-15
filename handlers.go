package main

import (
	"fmt"
)

func getNextID(tasks []Task) int {
    maxID := 0
    for _, task := range tasks {
        if task.ID > maxID {
            maxID = task.ID
        }
    }
    return maxID + 1
}


// Adds a new task
func handleAdd(name string) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    newTask := Task{
        ID:        getNextID(tasks),
        Name:      name,
        Completed: false,
    }

    tasks = append(tasks, newTask)

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving task:", err)
        return
    }

    fmt.Println("Task added successfully:", name)
}

// Lists all tasks
func handleList() {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }

    for _, task := range tasks {
        status := " "
        if task.Completed {
            status = "✓"
        }
        fmt.Printf("[%s] %d: %s\n", status, task.ID, task.Name)
    }
}

// Marks a task as complete
func handleComplete(id int) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    found := false
    for i, task := range tasks {
        if task.ID == id {
            tasks[i].Completed = true
            found = true
            break
        }
    }

    if !found {
        fmt.Println("Task not found.")
        return
    }

    err = saveTasks(tasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }

    fmt.Println("Task marked as complete.")
}

// Removes a task
func handleRemove(id int) {
    tasks, err := loadTasks()
    if err != nil {
        fmt.Println("Error loading tasks:", err)
        return
    }

    newTasks := []Task{}
    found := false
    for _, task := range tasks {
        if task.ID != id {
            newTasks = append(newTasks, task)
        } else {
            found = true
        }
    }

    if !found {
        fmt.Println("Task not found.")
        return
    }

    err = saveTasks(newTasks)
    if err != nil {
        fmt.Println("Error saving tasks:", err)
        return
    }

    fmt.Println("Task removed.")
}
