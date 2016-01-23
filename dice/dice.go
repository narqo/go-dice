package dice

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type dice struct {
	num int
	die die
}

func (d dice) Roll() int {
	log.Print("rol: ", d)
	r := 0
	for i := 0; i < d.num; i++ {
		val := d.die.Roll()
		log.Print("rol: ", d.die, " got: ", val)
		r += val
	}
	log.Print("got: ", r)
	return r
}

func (d dice) String() string {
	return fmt.Sprintf("%d%s", d.num, d.die)
}

func Parse(notation string) (d *dice, err error) {
	log.Printf("parse: notation=%s", notation)

	var num int
	var faces int

	dc := strings.Split(notation, "d")
	if len(dc) == 2 {
		s := dc[0]
		if s == "" {
			num = 1
		} else {
			num, err = strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
		}
		faces, err = strconv.Atoi(dc[1])
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("bad notation \"" + notation + "\"")
	}

	log.Printf("parse: size=%d, faces=%d", num, faces)

	d = &dice{
		num: num,
		die: die{Faces: faces},
	}
	return d, nil
}
