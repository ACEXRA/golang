package main

func main() {
	taskList := createTaskList();

	addTask(taskList,"This is my first task")
	addTask(taskList,"This is my second task")

	for _, v := range taskList.taskList {
		println(v.String())
	}
}

//create struct of task
//create tasklist, add task, update task, delete task
