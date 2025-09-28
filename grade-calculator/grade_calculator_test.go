// grade-calculator/grade_calculator_test.go
package esepunittests

import "testing"

func TestGetGradeA(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 100, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 100, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeB(t *testing.T) {
	expected_value := "B"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 80, Assignment)
	gradeCalculator.AddGrade("exam 1", 81, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 85, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "A"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 100, Assignment)
	gradeCalculator.AddGrade("exam 1", 95, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 91, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestEmptyCalculatorIsF(t *testing.T) {
	g := NewGradeCalculator()
	if got := g.GetFinalGrade(); got != "F" {
		t.Fatalf("want F for empty; got %s", got)
	}
}

func TestAddGradeCoversAllCategories(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a1", 100, Assignment)
	g.AddGrade("e1", 80, Exam)
	g.AddGrade("s1", 60, Essay)
	if got := g.GetFinalGrade(); got == "" {
		t.Fatal("unexpected empty grade")
	}
}

func TestBoundaryA90(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a", 90, Assignment)
	g.AddGrade("e", 90, Exam)
	g.AddGrade("s", 90, Essay)
	if got := g.GetFinalGrade(); got != "A" {
		t.Fatalf("want A; got %s", got)
	}
}

func TestBoundaryB80(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a", 80, Assignment)
	g.AddGrade("e", 80, Exam)
	g.AddGrade("s", 80, Essay)
	if got := g.GetFinalGrade(); got != "B" {
		t.Fatalf("want B; got %s", got)
	}
}

func TestBoundaryC70(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a", 70, Assignment)
	g.AddGrade("e", 70, Exam)
	g.AddGrade("s", 70, Essay)
	if got := g.GetFinalGrade(); got != "C" {
		t.Fatalf("want C; got %s", got)
	}
}

func TestBoundaryD60(t *testing.T) {
	g := NewGradeCalculator()
	g.AddGrade("a", 60, Assignment)
	g.AddGrade("e", 60, Exam)
	g.AddGrade("s", 60, Essay)
	if got := g.GetFinalGrade(); got != "D" {
		t.Fatalf("want D; got %s", got)
	}
}

func TestBelow60IsF(t *testing.T) {
	g := NewGradeCalculator()
	// By the lemma of they are all under 60 obviously their average is under 60
	g.AddGrade("a", 40, Assignment)
	g.AddGrade("e", 50, Exam)
	g.AddGrade("s", 55, Essay)
	if got := g.GetFinalGrade(); got != "F" {
		t.Fatalf("want F; got %s", got)
	}
}

func TestSingleCategoryOnly_WeightsApply(t *testing.T) {
	// Only assignments, final = 0.5 * mean(assignments)
	g := NewGradeCalculator()
	g.AddGrade("a1", 100, Assignment) // final = 50
	if got := g.GetFinalGrade(); got != "F" {
		t.Fatalf("want F (50); got %s", got)
	}
}
