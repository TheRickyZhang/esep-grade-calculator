package esepunittests

type GradeType int

const (
	Assignment GradeType = iota
	Exam
	Essay
)

type grade struct {
	name  string
	score float64
	gt    GradeType
}

type GradeCalculator struct {
	assignments []grade
	exams       []grade
	essays      []grade
}

func NewGradeCalculator() *GradeCalculator { return &GradeCalculator{} }

func (g *GradeCalculator) AddGrade(name string, score int, gt GradeType) {
	gr := grade{name: name, score: float64(score), gt: gt}
	switch gt {
	case Assignment:
		g.assignments = append(g.assignments, gr)
	case Exam:
		g.exams = append(g.exams, gr)
	case Essay:
		g.essays = append(g.essays, gr)
	}
}

func avg(xs []grade) float64 {
	if len(xs) == 0 {
		return 0
	}
	var s float64
	for _, v := range xs {
		s += v.score
	}
	return s / float64(len(xs))
}

func (g *GradeCalculator) GetFinalGrade() string {
	final := 0.50*avg(g.assignments) + 0.35*avg(g.exams) + 0.15*avg(g.essays)

	switch {
	case final >= 90:
		return "A"
	case final >= 80:
		return "B"
	case final >= 70:
		return "C"
	case final >= 60:
		return "D"
	default:
		return "F"
	}
}

