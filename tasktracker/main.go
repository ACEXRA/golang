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
	updatedAt time.Time
}


func main() {
	fmt.Println("Hello world")
}