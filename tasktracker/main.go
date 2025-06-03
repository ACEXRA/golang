package main

import (
	"fmt"
	"time"
)

type status int

const(
	todo status = iota
	inProgress
	done
)

type task struct{
	id int
	description string
	status status
	createdAt time.Time
	updatedAt *time.Time
}

type taskList struct{
	taskList []task
}

func (t task) String() string {
	return fmt.Sprintf("task{\nid: %d\ndescription: %s\nstatus: %d\ncreatedAt: %s\nupdatedAt: %v\n}",
		t.id, t.description, t.status, t.createdAt.Format(time.RFC3339), t.updatedAt)
}

func main() {
	myList := &taskList{}

	addTask(myList,"This is my first task");
	addTask(myList,"This is my Second task");
	addTask(myList,"This is my Second task");

	for _,v := range myList.taskList{
		fmt.Print(v)
	}
}

var getId = func() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}()


func createTask(description string) task{
	return task{
		id:getId(),
		description: description,
		status: 0,
		createdAt: time.Now(),
		updatedAt: nil,
	}
}

func addTask(list *taskList,description string){
	newTask:=createTask(description);
	list.taskList = append(list.taskList,newTask)
}