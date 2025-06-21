package main

import (
	"fmt"
	"time"
)

type StatusCode int

const (
	todo StatusCode = iota
	progress
	done
)

var StatusMap = map[StatusCode]string{
	todo:     "todo",
	progress: "In-progress",
	done:     "done",
}

type Task struct {
	ID          int
	Description string
	Status      StatusCode
	CreatedAt   time.Time
	UpdatedAt   *time.Time
}

type TaskList struct {
	taskList []Task
}

func (t Task) String() string {
	var updatedAtStr string

	if t.UpdatedAt != nil {
		updatedAtStr = t.UpdatedAt.Format(time.RFC3339)
	} else {
		updatedAtStr = "<nil>"
	}

	return fmt.Sprintf("{\nid:%d,\ndescription:%s,\nstatus:%s,\ncreatedAt:%s,\nupdatedAt:%s\n}\n",
		t.ID, t.Description, StatusMap[t.Status], t.CreatedAt.Format(time.RFC3339), updatedAtStr)
}
