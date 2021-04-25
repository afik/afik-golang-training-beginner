package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Employee struct {
	Name           string   `json:"name"`
	Entity         string   `json:"entity"`
	EmployeeNumber int      `json:"employee_number"`
	Salary         float32  `json:"salary"`
	Domicile       Domicile `json:"domicile"`
}

type Domicile struct {
	Country  string `json:"country"`
	IsRemote bool   `json:"is_remote"`
}

func main() {
	data := string(`
    {
        "name": "Golang",
        "entity": "Xendit",
        "employee_number": 10,
        "salary": 1.5,
        "domicile": {
            "country": "ID",
            "is_remote": true
        }
    }
    `)
	var employee Employee
	if err := json.Unmarshal([]byte(data), &employee); err != nil {
		fmt.Println("Task #1 failed!")
		fmt.Print(err)
		return
	}

	database := make(map[string]Employee)
	database["Golang"] = employee
	if !reflect.DeepEqual(database["Golang"], employee) {
		fmt.Println("Task #2 failed!")
	}

	fmt.Println("All good")
}
