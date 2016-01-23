package dice

import (
	"math/rand"
	"fmt"
)

type die struct {
	Faces int
}

func (d die) Roll() int {
	return rand.Intn(d.Faces) + 1
}

func (d die) String() string {
	return fmt.Sprintf("d%d", d.Faces)
}

var defaultDie = die{Faces: 6}

func RollD6() int {
	return defaultDie.Roll()
}
