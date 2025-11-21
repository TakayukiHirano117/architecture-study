package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestEvaluation_NewEvaluation_Success(t *testing.T) {
	validEvaluation := 5

	evaluation, err := userdm.NewEvaluation(validEvaluation)
	if err != nil {
		t.Errorf("NewEvaluation() with valid evaluation should not return error, got: %v", err)
	}

	if evaluation.Int() != validEvaluation {
		t.Errorf("NewEvaluation() should return correct value, expected: %d, got: %d", validEvaluation, evaluation.Int())
	}
}

func TestEvaluation_NegativeEvaluationReturnError(t *testing.T) {
	negativeEvaluation := -1

	_, err := userdm.NewEvaluation(negativeEvaluation)
	if err == nil {
		t.Errorf("NewEvaluation() with negative evaluation must be positive")
	}
}

func TestEvaluation_TooLongEvaluationReturnError(t *testing.T) {
	tooLongEvaluation := 6

	_, err := userdm.NewEvaluation(tooLongEvaluation)
	if err == nil {
		t.Errorf("NewEvaluation() with too long evaluation must be less than 5")
	}
}