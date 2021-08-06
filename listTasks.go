package main

import "fmt"

type Tasks struct {
	BaseList
}

type Task struct {
	Name     string
	Priority string
	State    string
}

func DisplayList() {
	display("======================Contact List======================================")
	newContancts := Contancts{}

	display("=============Adding contact johh===============")
	johnId := newContancts.Create(Contact{
		FirstName: "johh",
		LastName:  "Doe",
	})
	fmt.Printf("created object John: %v\n", newContancts.Get(johnId))

	display("=============Adding contact Mister=============")
	twisterId := newContancts.Create(Contact{
		FirstName: "Mister",
		LastName:  "Twister",
	})
	fmt.Printf("created object Twister: %v\n", newContancts.Get(twisterId))

	display("=============Updating contact johh to John=============")
	newContancts.Update(johnId, Contact{
		FirstName: "Jonh",
		LastName:  "Doe",
	})
	fmt.Printf("updated object John: %v\n", newContancts.Get(johnId))

	display("=============Printing all=============")
	fmt.Printf("All objects: %v\n", newContancts.GetAll())

	display("=============Deleting John=============")
	newContancts.Delete(johnId)
	fmt.Printf("John deleted: %d\n", johnId)

	display("=============Printing all again=============")
	fmt.Printf("All objects: %v\n", newContancts.GetAll())
}
