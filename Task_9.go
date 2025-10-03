package main

import (
 "fmt"
)

type employee struct {
 id     int
 name   string
 salary float64
 dept   int
}

type department struct {
 id        int
 name      string
 employees []*employee
}

func (d *department) totalSalary() float64 {
 var totalSum = 0.0
 for _, emp := range d.employees {
  totalSum += emp.salary
 }
 fmt.Printf("Фонд зарплаты отдела '%s': %.2f руб.\n", d.name, totalSum)
 return totalSum
}

func (d *department) addEmployee(emp *employee) {
 d.employees = append(d.employees, emp)
 fmt.Printf("Сотрудник '%s' добавлен в отдел '%s'\n", emp.name, d.name)
}

func (d *department) removeEmployee(idEmp int) {
 for i, emp := range d.employees {
  if emp.id == idEmp {
   d.employees = append(d.employees[:i], d.employees[i+1:]...)
   fmt.Printf("Сотрудник с ID-%d удалён из отдела '%s'\n", idEmp, d.name)
   break
  }
 }
}

func (d *department) listEmployees() []*employee {
 fmt.Printf("Список сотрудников отдела '%s':\n", d.name)
 for _, emp := range d.employees {
  fmt.Printf("- %s\n", emp.name)
 }
 return d.employees
}

func main() {
 marketingDept := department{
  id:        2,
  name:      "Отдел маркетинга",
  employees: []*employee{},
 }

 emp1 := &employee{id: 1001, name: "Антонов Антон Антонович", salary: 90000.0, dept: 2}
 emp2 := &employee{id: 1002, name: "Петрова Мария Сергеевна", salary: 80000.0, dept: 2}

 marketingDept.addEmployee(emp1)
 marketingDept.addEmployee(emp2)
 marketingDept.totalSalary()
 marketingDept.listEmployees()
 marketingDept.removeEmployee(1001)
 marketingDept.listEmployees()
}
