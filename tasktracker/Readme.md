This is a task tracker

create a task list 
add task
update task
delete task

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

func main() {
	myList := &taskList{}

	addTask(myList,"This is my first task");
	addTask(myList,"This is my Second task");
	updateTask(myList,2,"The second is updated",1)
	deleteTask(myList,3)

	for _,v := range myList.taskList{
		fmt.Print(v)
	}
}

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

func updateTask(list *taskList, id int, description string, status status){
	task := getByID(id,list);
	if task ==nil{
		println("Task doesn't exist")
		return
	}
	task.description = description;
	task.status = status;
	now := time.Now()
	task.updatedAt = &now
	
}
func deleteTask(list *taskList, id int){
	for i:=range list.taskList{
		if(list.taskList[i].id==id){
			list.taskList = append(list.taskList[:i],list.taskList[i+1:]...)
			return
		}
	}
	println("Task doesn't exist")
}

//helper 
func getByID(id int, list *taskList) *task {
	for i := range list.taskList {
		if list.taskList[i].id == id {
			return &list.taskList[i] 
		}
	}
	return nil
}


var getId = func() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}()

func (s status) String() string{
	switch s{
	case todo:
		return "To do";
	case inProgress:
		return "In Progress";
	case done:
		return "Done"
	default:
		return "Not set"
	}
}

func (t task) String() string {
	var updatedAtStr string;
	if(t.updatedAt!=nil){
		updatedAtStr =t.updatedAt.Format(time.RFC3339);
	}else{
		updatedAtStr = "<nil>"
	}
	return fmt.Sprintf("task{\nid: %d\ndescription: %s\nstatus: %s\ncreatedAt: %s\nupdatedAt: %v\n}\n",
		t.id, t.description, t.status.String(), t.createdAt.Format(time.RFC3339), updatedAtStr)
}