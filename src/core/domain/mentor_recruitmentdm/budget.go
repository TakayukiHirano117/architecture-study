package mentor_recruitmentdm

import "github.com/cockroachdb/errors"

type Budget struct {
	min uint32
	max uint32
}

const (
	MinBudgetValue uint32 = 1000
	MaxBudgetValue uint32 = 1000000
)

func NewBudget(min, max uint32) (Budget, error) {
	if min < MinBudgetValue {
		return Budget{}, errors.Newf("minimum budget must be at least %d", MinBudgetValue)
	}

	if max > MaxBudgetValue {
		return Budget{}, errors.Newf("maximum budget must be at most %d", MaxBudgetValue)
	}

	if min > max {
		return Budget{}, errors.New("minimum budget must be less than or equal to maximum budget")
	}

	return Budget{min: min, max: max}, nil
}

func (b Budget) Min() uint32 {
	return b.min
}

func (b Budget) Max() uint32 {
	return b.max
}

func (b Budget) Equal(b2 Budget) bool {
	return b.min == b2.min && b.max == b2.max
}
