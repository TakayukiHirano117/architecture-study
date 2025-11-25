package userdm

import "github.com/cockroachdb/errors"

type Evaluation uint8

func NewEvaluation(value uint8) (Evaluation, error) {
	if value > 5 {
		return 0, errors.New("Evaluation is too large")
	}

	return Evaluation(value), nil
}

func NewEvaluationByVal(value uint8) (Evaluation, error) {
	return Evaluation(value), nil
}

func (e Evaluation) Uint8() uint8 {
	return uint8(e)
}

func (e Evaluation) Equal(e2 Evaluation) bool {
	return e == e2
}
