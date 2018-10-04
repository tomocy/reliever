package grader

import (
	"fmt"
	"strings"
)

type Grades map[string][]Course

func (g Grades) String() string {
	b := make([]byte, 0, 10)
	b = append(b, hello()...)
	for kind, courses := range g {
		b = append(b, fmt.Sprintf("%s: %d (%s)\n", kind, sumCredits(courses...), joinCourseNames(courses, ", "))...)
	}
	b = append(b, bye()...)

	return string(b)
}

func hello() string {
	return "\nYour grade ----------\n\n"
}

func sumCredits(courses ...Course) int {
	var credits int
	for _, course := range courses {
		credits += course.Credits
	}

	return credits
}

func joinCourseNames(courses []Course, separator string) string {
	courseNames := make([]string, len(courses))
	for i, course := range courses {
		courseNames[i] = course.Name
	}

	return strings.Join(courseNames, separator)
}

func bye() string {
	return "\n---------- Good luck\n"
}
