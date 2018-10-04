package grader

import (
	"encoding/csv"
	"io"
	"log"
)

type Grader interface {
	Grade() (Grades, error)
}

type csvGrader struct {
	reader io.Reader
}

func NewCSVGrader(reader io.Reader) Grader {
	return &csvGrader{reader: reader}
}

func (g csvGrader) Grade() (Grades, error) {
	courseNames := make([]string, 0)
	scanner := csv.NewReader(g.reader)
	for {
		grade, err := scanner.Read()
		if err != nil {
			if err == io.EOF {
				break
			}

			return nil, err
		}

		// grade[6] is the name of course
		courseNames = append(courseNames, grade[6])
	}

	return grade(courseNames), nil
}

func grade(courseNames []string) Grades {
	grades := make(Grades)
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
