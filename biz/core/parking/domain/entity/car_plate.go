package parking_entity

import (
	"code.byted.org/oec/status_code/Status"
	E "github.com/dirac-lee/ddd_demo_parking/biz/common/exception"
)

func NewCarPlate(value string) CarPlate {
	if len(value) == 0 {
		E.Throw(E.New(Status.InvalidParam, E.WithDetail("car plate should not be empty")))
	}
	return CarPlate{value: value}
}

type CarPlate struct {
	value string
}

func (cp CarPlate) String() string {
	return cp.value
}
