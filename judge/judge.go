package judge

import (
	"log"

	"github.com/tomocy/reliever/grader"
)

type Judge interface {
	Judge() Judgement
}

type judgeOfLetters struct {
	grades grader.Grades
}

func NewJudgeOfLetters(grades grader.Grades) Judge {
	newGrades := make(grader.Grades)
	for k, v := range grades {
		newGrades[k] = v
	}
	judge := &judgeOfLetters{
		grades: newGrades,
	}
	judge.modifyGrades()

	return judge
}

func (j *judgeOfLetters) modifyGrades() {
	newGrades := make(grader.Grades)
	for course, classes := range j.grades {
		switch course {
		case "先端教養科目", "国際教養１":
			newGrades["先端教養科目 + 国際教養１"] = append(j.grades["先端教養科目"], j.grades["国際教養１"]...)
			delete(j.grades, "先端教養科目")
			delete(j.grades, "国際教養１")
		case "大学英語", "実践英語・専門英語":
			newGrades["大学英語 + 実践英語・専門英語"] = append(j.grades["大学英語"], j.grades["実践英語・専門英語"]...)
			delete(j.grades, "大学英語")
			delete(j.grades, "実践英語・専門英語")
		case "ドイツ語", "フランス語", "ロシア語", "中国語", "朝鮮語", "スペイン語", "イタリア語", "日本語":
			newGrades["第２外国語"] = append(j.grades["ドイツ語"], j.grades["フランス語"]...)
			newGrades["第２外国語"] = append(newGrades["第２外国語"], append(j.grades["ロシア語"], j.grades["中国語"]...)...)
			newGrades["第２外国語"] = append(newGrades["第２外国語"], append(j.grades["朝鮮語"], j.grades["スペイン語"]...)...)
			newGrades["第２外国語"] = append(newGrades["第２外国語"], append(j.grades["イタリア語"], j.grades["日本語"]...)...)
			delete(j.grades, "ドイツ語")
			delete(j.grades, "フランス語")
			delete(j.grades, "ロシア語")
			delete(j.grades, "中国語")
			delete(j.grades, "朝鮮語")
			delete(j.grades, "スペイン語")
			delete(j.grades, "イタリア語")
			delete(j.grades, "日本語")
		case "第１外国語", "第２外国語", "第３外国語":
			newGrades["選択外国語"] = append(j.grades["第１外国語"], append(j.grades["第２外国語"], j.grades["第３外国語"]...)...)
			delete(j.grades, "第１外国語")
			delete(j.grades, "第２外国語")
			delete(j.grades, "第３外国語")
		default:
			newGrades[course] = classes
		}
	}

	j.grades = newGrades
}

var criteriaOfLetters = map[string]int{
	"基礎教養1":            2,
	"基礎教養2":            2,
	"現代教養科目":           2,
	"先端教養科目 + 国際教養１":   2,
	"国際教養２":            8,
	"大学英語 + 実践英語・専門英語": 8,
	"第２外国語":            4,
	"選択外国語":            4,
	"情報処理教育科目":         2,
	"健康・スポーツ教育科目":      2,
	"専門基礎教育科目":         4,
}

func (j judgeOfLetters) Judge() Judgement {
	judgement := newJudgement(criteriaOfLetters)
	for course, classes := range j.grades {
		criterion, ok := criteriaOfLetters[course]
		if !ok {
			log.Println("unknown course: ", course)
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
