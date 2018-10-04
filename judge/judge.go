package judge

import "github.com/tomocy/reliever/grader"

type Judge interface {
	Judge() Judgement
}

type judgeOfLetters struct {
	grades grader.Grades
}

func NewJudgeOfLetters(grades grader.Grades) Judge {
	judge := &judgeOfLetters{}
	judge.grades = judge.modifyGradesIfMultipleCourse(grades)

	return judge
}

func (j judgeOfLetters) modifyGradesIfMultipleCourse(grades grader.Grades) grader.Grades {
	newGrades := make(grader.Grades)
	for course, classes := range grades {
		switch course {
		case "先端教養科目":
			newGrades["先端教養科目 + 国際教養１"] = append(classes, grades["国際教養１"]...)
			delete(grades, "国際教養１")
		case "国際教養１":
			newGrades["先端教養科目 + 国際教養１"] = append(grades["先端教養科目"], classes...)
			delete(grades, "先端教養科目")
		case "大学英語":
			newGrades["大学英語 + 実践英語・専門英語"] = append(classes, grades["実践英語・専門英語"]...)
			delete(grades, "実践英語・専門英語")
		case "実践英語・専門英語":
			newGrades["大学英語 + 実践英語・専門英語"] = append(grades["大学英語"], classes...)
			delete(grades, "大学英語")
		default:
			newGrades[course] = classes
		}
	}

	return newGrades
}

var criteriaOfLetters = map[string]int{
	"基礎教養1":            2,
	"基礎教養2":            2,
	"現代教養科目":           2,
	"先端教養科目 + 国際教養１":   2,
	"国際教養２":            8,
	"大学英語 + 実践英語・専門英語": 8,
	"第２外国語":            4,
	"情報処理教育科目":         2,
	"健康・スポーツ教育科目":      2,
	"専門基礎教育科目":         4,
}

func (j judgeOfLetters) Judge() Judgement {
	judgement := newJudgement(criteriaOfLetters)
	for course, classes := range j.grades {
		criterion, ok := criteriaOfLetters[course]
		if !ok {
			continue
		}
		credits := sumCredits(classes...)
		if criterion <= credits {
			judgement.EnoughCourses[course] = append(judgement.EnoughCourses[course], classes...)
			continue
		}
		judgement.NotEnoughCourses[course] = append(judgement.NotEnoughCourses[course], classes...)
	}

	return *judgement
}

func sumCredits(classes ...grader.Course) int {
	var credits int
	for _, class := range classes {
		credits += class.Credits
	}

	return credits
}
