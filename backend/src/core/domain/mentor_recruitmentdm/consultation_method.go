package mentor_recruitmentdm

import "github.com/cockroachdb/errors"

type ConsultationMethod string

const (
	Chat      ConsultationMethod = "チャット"
	VideoCall ConsultationMethod = "ビデオ通話"
)

func NewConsultationMethod(value string) (ConsultationMethod, error) {
	method := ConsultationMethod(value)
	if method != Chat && method != VideoCall {
		return "", errors.New("consultation method must be チャット or ビデオ通話")
	}

	return method, nil
}

func NewConsultationMethodByVal(value string) (ConsultationMethod, error) {
	if value == "" {
		return "", errors.New("consultation method must not be empty")
	}

	return ConsultationMethod(value), nil
}

func (c ConsultationMethod) String() (string, error) {
	switch c {
	case Chat:
		return string(Chat), nil
	case VideoCall:
		return string(VideoCall), nil
	}

	return "", errors.New("invalid consultation method")
}

func (c ConsultationMethod) Equal(c2 ConsultationMethod) bool {
	return c == c2
}
