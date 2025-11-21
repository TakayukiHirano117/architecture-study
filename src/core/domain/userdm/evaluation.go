package userdm

import "github.com/cockroachdb/errors"

type Evaluation int

func NewEvaluation(value int) (Evaluation, error) {
	if value < 0 {
		return 0, errors.New("Evaluation is invalid")
	}

	if value > 5 {
		return 0, errors.New("Evaluation is too large")
	}

	return Evaluation(value), nil
}

func NewEvaluationByVal(value int) (Evaluation, error) {
	return Evaluation(value), nil
}

func (e Evaluation) Int() int {
	return int(e)
}

func (e Evaluation) Equal(e2 Evaluation) bool {
	return e == e2
}