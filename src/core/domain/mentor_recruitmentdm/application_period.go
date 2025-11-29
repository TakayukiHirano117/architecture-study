package mentor_recruitmentdm

import (
	"time"

	"github.com/cockroachdb/errors"
)

type ApplicationPeriod time.Time

func NewApplicationPeriod() ApplicationPeriod {
	now := time.Now()
	deadline := now.AddDate(0, 0, 14)

	// 時刻を00:00:00にして日付のみを保持
	date := time.Date(deadline.Year(), deadline.Month(), deadline.Day(), 0, 0, 0, 0, deadline.Location())

	return ApplicationPeriod(date)
}

func NewApplicationPeriodByVal(value time.Time) (ApplicationPeriod, error) {
	if value.IsZero() {
		return ApplicationPeriod{}, errors.New("application period must not be empty")
	}

	date := time.Date(value.Year(), value.Month(), value.Day(), 0, 0, 0, 0, value.Location())

	return ApplicationPeriod(date), nil
}

func (ap ApplicationPeriod) Time() time.Time {
	return time.Time(ap)
}

func (ap ApplicationPeriod) String() string {
	return time.Time(ap).Format("1月02日")
}

func (ap ApplicationPeriod) Equal(ap2 ApplicationPeriod) bool {
	return time.Time(ap).Equal(time.Time(ap2))
}
