package main

import (
	"fmt"
)

type Contancts struct {
	BaseList
}

type Contact struct {
	FirstName string
	LastName  string
}

func DisplayTasks() {
	display("===================== Same CRUD for Tasks List======================================")

	display("======================Task List======================================")
	tasks := Tasks{}

	display("=============Adding contact johh===============")
	task1 := tasks.Create(Task{
		Name:     "Sleep well",
		Priority: "High",
		State:    "Not started",
	})
	fmt.Printf("created task to Sleep well: %v\n", tasks.Get(task1))

	display("=============Adding contact Mister=============")
	task2 := tasks.Create(Task{
		Name:     "Code more",
		Priority: "Very High",
		State:    "Ongoing",
	})
	fmt.Printf("created task to Code: %v\n", tasks.Get(task2))

	display("=============Updating task Sleep Well to just sleep=============")
	tasks.Update(task1, Task{
		Name:     "Sleep",
		Priority: "Middle",
		State:    "Not started",
	})
	fmt.Printf("updated task Sleep: %v\n", tasks.Get(task1))

	display("=============Printing all=============")
	fmt.Printf("All tasks: %v\n", tasks.GetAll())

	display("=============Deleting task to Sleep =============")
	tasks.Delete(task1)
	fmt.Printf("Forget to sleep: %d\n", task1)

	display("=============Printing all again=============")
	fmt.Printf("All tasks: %v\n", tasks.GetAll())
}
