package main

import (
	"fmt"
	"math/rand"
)

type tasks struct {
	taskId      int
	taskName    string
	description string
	status      string
}


func (t tasks) newTask(taskName string, description string, status string) *tasks {

	taskId := rand.Intn(1000)

	task := tasks{
		taskId:      taskId,
		taskName:    taskName,
		description: description,
		status:      status,
	}

	return &task

}

func (t *tasks) changeTaskStatus(status string) {
	t.status = status
}

func (t *tasks) editTask(taskName string, description string,status string) {
	t.taskName = taskName
	t.description = description
	t.status = status
}

func taskList(tasks ...*tasks) {
	for _, task := range tasks {
		fmt.Println("ID:", task.taskId, "Name:", task.taskName, "Description:", task.description, "Status:", task.status)
	}
}

func main() {

	task1 := tasks{}.newTask("Sleep", "This is task 1", "pending")
	task1.editTask("Wake up", "This is task 1 edited", "completed")

	task2 := tasks{}.newTask("Breakfast", "Have breakfast", "pending")
	task3 := tasks{}.newTask("Learn go", "Learn Go programming", "completed")
	task4 := tasks{}.newTask("Playing games", "Play some games", "pending")

	taskList(task1, task2, task3, task4)


}