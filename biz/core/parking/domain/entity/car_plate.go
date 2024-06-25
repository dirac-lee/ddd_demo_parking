package parking_entity

type CarPlate string

func (cp CarPlate) String() string {
	return string(cp)
}
