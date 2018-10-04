package main

import (
	"encoding/csv"
	"io"
	"log"
)

type grader interface {
	grade() (grades, error)
}

type csvGrader struct {
	reader io.Reader
}

func newCSVGrader(reader io.Reader) grader {
	return &csvGrader{reader: reader}
}

func (g csvGrader) grade() (grades, error) {
	courseNames := make([]string, 0)
	scanner := csv.NewReader(g.reader)
	for {
		unit, err := scanner.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		courseNames = append(courseNames, unit[6])
	}

	return grade(courseNames), nil
}

func grade(courseNames []string) grades {
	grades := make(grades)
	for _, courseName := range courseNames {
		course, ok := courses[courseName]
		if !ok {
			log.Println("unknown course name: ", courseName)
			continue
		}
		grades[course.kind] = append(grades[course.kind], course)
	}

	return grades
}
