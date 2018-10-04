package main

import (
	"fmt"
	"strings"
)

type grades map[string][]course

func (g grades) String() string {
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

func sumCredits(courses ...course) int {
	var credits int
	for _, course := range courses {
		credits += course.credits
	}

	return credits
}

func joinCourseNames(courses []course, separator string) string {
	courseNames := make([]string, len(courses))
	for i, course := range courses {
		courseNames[i] = course.name
	}

	return strings.Join(courseNames, separator)
}

func bye() string {
	return "\n---------- Good luck\n"
}
