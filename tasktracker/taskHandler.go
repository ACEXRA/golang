package main

import "time"

var getId = id()

func createTaskList() *TaskList{
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

func addTask(tasklist *TaskList,description string){
	tasklist.taskList = append(tasklist.taskList, createTask(description))
}

func updateTaskStatus(taskList *TaskList,id int, status StatusCode){
	
}
