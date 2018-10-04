package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tomocy/reliever/grader"
)

func main() {
	gradesFile, err := os.Open("grades.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer gradesFile.Close()

	grader := grader.NewCSVGrader(gradesFile)
	grades, err := grader.Grade()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(grades)
}
