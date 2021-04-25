package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// os.Setenv("task_1", "something")
	var task1 = os.Getenv("task_1")

	if !reflect.DeepEqual("", task1) {
		fmt.Println("Task #1 failed!")
	}

	var challenge1 interface{}
	challenge1 = "1"
	if !reflect.DeepEqual("1", challenge1) {
		fmt.Println("Challenge #1 failed!")
	}

	fmt.Println("All passed!")
}
