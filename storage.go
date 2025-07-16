package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const fileName = "tasks.json"

func loadTasks() ([]Task, error) {
    var tasks []Task

    if _, err := os.Stat(fileName); os.IsNotExist(err) {
        return tasks, nil // no tasks yet
    }

    data, err := ioutil.ReadFile(fileName)
    if err != nil {
        return nil, err
    }

    err = json.Unmarshal(data, &tasks)
    if err != nil {
        return nil, err
    }

    return tasks, nil
}

func saveTasks(tasks []Task) error {
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return err
    }

    return ioutil.WriteFile(fileName, data, 0644)
}
