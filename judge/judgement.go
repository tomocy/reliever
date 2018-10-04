package judge

import (
	"fmt"
	"strings"

	"github.com/tomocy/reliever/grader"
)

type Judgement struct {
	criteria         map[string]int
	EnoughCourses    grader.Grades
	NotEnoughCourses grader.Grades
}

func newJudgement(criteria map[string]int) *Judgement {
	return &Judgement{
		criteria:         criteria,
		EnoughCourses:    make(grader.Grades),
		NotEnoughCourses: make(grader.Grades),
	}
}

func (j Judgement) String() string {
	b := make([]byte, 0, 10)
	b = append(b, "\nEnough courses ----------\n"...)
	for course, classes := range j.EnoughCourses {
		criterion, ok := j.criteria[course]
		if !ok {
			continue
		}
		b = append(b, fmt.Sprintf("\t%s %d: %d (%s)\n", course, criterion, sumCredits(classes...), joinClassNames(classes, ", "))...)
	}
	b = append(b, "\nNot Enough courses ----------\n"...)
	for course, classes := range j.NotEnoughCourses {
		criterion, ok := j.criteria[course]
		if !ok {
			continue
		}
		b = append(b, fmt.Sprintf("\t%s %d: %d (%s)\n", course, criterion, sumCredits(classes...), joinClassNames(classes, ", "))...)
	}

	return string(b)
}

func joinClassNames(classes []grader.Course, separator string) string {
	classNames := make([]string, len(classes))
	for i, class := range classes {
		classNames[i] = class.Name
	}

	return strings.Join(classNames, separator)
}
