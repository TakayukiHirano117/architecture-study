package userdm_test

import (
	"testing"

	"github.com/TakayukiHirano117/architecture-study/src/core/domain/userdm"
)

func TestEvaluation_NewEvaluation_Success(t *testing.T) {
	var validEvaluation uint8 = 5

	evaluation, err := userdm.NewEvaluation(validEvaluation)
	if err != nil {
		t.Errorf("NewEvaluation() with valid evaluation should not return error, got: %v", err)
	}

	if evaluation.Uint8() != validEvaluation {
		t.Errorf("NewEvaluation() should return correct value, expected: %d, got: %d", validEvaluation, evaluation.Uint8())
	}
}

func TestEvaluation_TooLongEvaluationReturnError(t *testing.T) {
	var tooLongEvaluation uint8 = 6

	_, err := userdm.NewEvaluation(tooLongEvaluation)
	if err == nil {
		t.Errorf("NewEvaluation() with too long evaluation must be less than 5")
	}
}
