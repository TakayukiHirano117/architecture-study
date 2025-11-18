package userdm

import "github.com/cockroachdb/errors"

type Evaluation int

func NewEvaluation(value int) (*Evaluation, error) {
	if value < 0 {
		return nil, errors.New("Evaluation is invalid")
	}

	if value > 5 {
		return nil, errors.New("Evaluation is too large")
	}

	evaluation := Evaluation(value)

	return &evaluation, nil
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