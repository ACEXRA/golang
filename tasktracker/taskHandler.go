package main

import (
	"time"
)

var getId = generateId()

func createTaskList() *TaskList {
	return &TaskList{}
}

func createTask(description string) Task {

	return Task{
		ID:          getId(),
		Description: description,
		Status:      todo,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
	}
}

func addTask(tasklist *TaskList, description string) {
	tasklist.taskList = append(tasklist.taskList, createTask(description))
}

func updateTaskStatus(taskList *TaskList, id int, status StatusCode) {
	now := time.Now()
	for i := range taskList.taskList {
		if taskList.taskList[i].ID == id {
			taskList.taskList[i].Status = status
			taskList.taskList[i].UpdatedAt = &now
			break
		}
	}
}

func updateTaskDesc(taskList *TaskList, id int, description string) {
	now := time.Now()
	for i := range taskList.taskList {
		if taskList.taskList[i].ID == id {
			taskList.taskList[i].Description = description
			taskList.taskList[i].UpdatedAt = &now
			break
		}
	}
}

func deleteTask(taskList *TaskList, id int) {
	for i := range taskList.taskList {
		if taskList.taskList[i].ID == id {
			taskList.taskList = append(taskList.taskList[:i], taskList.taskList[i+1:]...)
			break
		}
	}
}
