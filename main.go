package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	unitsList, err := os.Open("grades.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer unitsList.Close()

	grader := newCSVGrader(unitsList)
	grades, err := grader.grade()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(grades)
}
