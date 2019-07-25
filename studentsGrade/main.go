package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type grades []int

func (g grades) String() (r string) {
	for _, gd := range g {
		if gd >= 38 {
			cpx := gd % 5

			if cpx >= 3 {
				gd += (5 - cpx)
			}
		}

		r += fmt.Sprintf("%d\n", gd)
	}
	return
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputText, _, _ := reader.ReadLine()                   // in real application's don't exclude the error param.
	numberOfStudents, _ := strconv.Atoi(string(inputText)) // in real application's don't exclude the error param.
	studentsGrades := grades{}

	for i := 0; i < numberOfStudents; i++ {
		nextInput, _, _ := reader.ReadLine()               // in real application's don't exclude the error param.
		studentGrade, _ := strconv.Atoi(string(nextInput)) // in real application's don't exclude the error param.

		studentsGrades = append(studentsGrades, studentGrade)
	}

	fmt.Printf("%s", studentsGrades)
}
