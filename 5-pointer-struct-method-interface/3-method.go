package main

import "fmt"

type Employee struct {
	Firstname string
	Lastname  string
}

type Manager struct {
	Firstname string
	Lastname  string
	Title     string
}

//function
func fullName(firstName string, lastName string) (fName string) {
	fName = firstName + "-" + lastName
	return
}

//method
func (emp Employee) fullName() string {
	return emp.Firstname + "-" + emp.Lastname
}

func (mgr Manager) fullName() string {
	return mgr.Firstname + "-" + mgr.Lastname + "-" + mgr.Title
}

func main() {
	emp1 := Employee{
		Firstname: "john",
		Lastname:  "doe",
	}

	emp2 := Employee{
		Firstname: "budi",
		Lastname:  "sudarsono",
	}

	mgr1 := Manager{
		Firstname: "budi",
		Lastname:  "sudarsono",
		Title:     "Ph.D",
	}

	fmt.Println(fullName(emp1.Firstname, emp1.Lastname))
	fmt.Println("emp 1", emp1.fullName())
	fmt.Println("emp 2", emp2.fullName())
	fmt.Println(mgr1.fullName())

}
