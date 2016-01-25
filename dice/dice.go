package dice

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/narqo/go-dice/Godeps/_workspace/src/github.com/apex/log"
)

type dice struct {
	num int
	die die
}

func (d dice) Roll() int {
	log.WithFields(log.Fields{
		"dice": d,
	}).Debug("rol")

	r := 0
	for i := 0; i < d.num; i++ {
		val := d.die.Roll()
		log.WithFields(log.Fields{
			"dice": d.die,
			"got":  val,
		}).Debug("rol")
		r += val
	}

	log.WithFields(log.Fields{
		"dice": d,
		"got":  r,
	}).Debug("rol")

	return r
}

func (d dice) String() string {
	return fmt.Sprintf("%d%s", d.num, d.die)
}

func Parse(notation string) (d *dice, err error) {
	log.WithFields(log.Fields{
		"notation": notation,
	}).Debug("parse")

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

	log.WithFields(log.Fields{
		"size":  num,
		"faces": faces,
	}).Debug("parse")

	d = &dice{
		num: num,
		die: die{Faces: faces},
	}
	return d, nil
}
