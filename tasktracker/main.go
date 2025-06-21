package main

func main() {
	taskList := createTaskList()

	addTask(taskList, "This is my first task")
	addTask(taskList, "This is my second task")
	addTask(taskList, "This is my third task")
	addTask(taskList, "This is my fourth task")
	addTask(taskList, "This is my fifth task")
	updateTaskStatus(taskList, 2, done)
	updateTaskDesc(taskList, 1, "This is an edited task")
	deleteTask(taskList, 3)
	deleteTask(taskList,2)
	deleteTask(taskList,1)
	deleteTask(taskList,5)

	for _, v := range taskList.taskList {
		println(v.String())
	}
}

//create struct of task
//create tasklist, add task, update task, delete task
