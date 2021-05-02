package main

import (
	"fmt"
	"strconv"
)

func main() {
	tests := []struct {
		name string
		in   int
		want string
	}{
		{
			name: "Should return number 1",
			in:   1,
			want: "1",
		}, {
			name: "Should return Xen",
			in:   9,
			want: "Xen",
		}, {
			name: "Should return dit",
			in:   24,
			want: "dit",
		}, {
			name: "Should return Xendit",
			in:   30,
			want: "Xendit",
		},
		{
			name: "Should return Xendit",
			in:   20,
			want: "Xendit",
		}
	}
	for _, tt := range tests {
		result := xenDit(tt.in)
		if result != tt.want {
			fmt.Println("Task #1 - Failed! ", tt.in, tt.name)
			panic(1)
		}
	}
	fmt.Println("Task #1 - Passed!")
}

func xenDit(i int) string {
	if multiplicationOf(i, 10) {
		return "Xendit"
	}

	if multiplicationOf(i, 3) {
		if multiplicationOf(i, 2) {
			return "dit"
		}
		return "Xen"
	}

	return strconv.Itoa(i)
}

func multiplicationOf(number int, divider int) bool {
	return (number%divider == 0)
}
