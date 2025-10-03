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

func TestGetGradeC(t *testing.T) {
	expected_value := "C"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 78, Assignment)
	gradeCalculator.AddGrade("exam 1", 75, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 73, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeD(t *testing.T) {
	expected_value := "D"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 65, Assignment)
	gradeCalculator.AddGrade("exam 1", 66, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 68, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGetGradeF(t *testing.T) {
	expected_value := "F"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 42, Assignment)
	gradeCalculator.AddGrade("exam 1", 60, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetFinalGrade()

	if expected_value != actual_value {
		t.Errorf("Expected GetGrade to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestGradeTypeString(t *testing.T) {
	if Assignment.String() != "assignment" {
		t.Errorf("Expected GradeType.String to return '%s'; got '%s' instead", "assignment", Assignment.String())
	}
	if Exam.String() != "exam" {
		t.Errorf("Expected GradeType.String to return '%s'; got '%s' instead", "exam", Exam.String())
	}
	if Essay.String() != "essay" {
		t.Errorf("Expected GradeType.String to return '%s'; got '%s' instead", "essay", Essay.String())
	}
}

func TestPass(t *testing.T) {
	expected_value := "Pass"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 90, Assignment)
	gradeCalculator.AddGrade("exam 1", 83, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 78, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}

func TestFail(t *testing.T) {
	expected_value := "Fail"

	gradeCalculator := NewGradeCalculator()

	gradeCalculator.AddGrade("open source assignment", 43, Assignment)
	gradeCalculator.AddGrade("exam 1", 24, Exam)
	gradeCalculator.AddGrade("essay on ai ethics", 55, Essay)

	actual_value := gradeCalculator.GetPassFail()

	if expected_value != actual_value {
		t.Errorf("Expected GetPassFail to return '%s'; got '%s' instead", expected_value, actual_value)
	}
}
