package main

import (
	"fmt"
	"log"
	"os"

	"github.com/tomocy/reliever/grader"
	"github.com/tomocy/reliever/judge"
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
	judge := judge.NewJudgeOfLetters(grades)
	judgement := judge.Judge()

	fmt.Println(judgement)
}
