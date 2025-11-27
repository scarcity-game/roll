package dice

import (
	"errors"
	"math/rand"
)

type Die struct {
	number int
	sides  int
}

type Specification struct {
	dice       []Die
	additional int
	valid      bool
}

const MaxDiceTypes = 100
const MaxDice = 100
const MaxSides = 2 << 31

func (s *Specification) Validate() error {
	if len(s.dice) > MaxDiceTypes {
		return errors.New("too many dice types")
	}
	for _, die := range s.dice {
		if die.number > MaxDice {
			return errors.New("too many dice")
		}
		if die.sides > MaxSides {
			return errors.New("too many dice sides")
		}
		if die.sides < 1 {
			return errors.New("too few dice sides")
		}
		if die.number < 1 {
			return errors.New("too few dice")
		}
	}
	s.valid = true
	return nil
}

func (s *Specification) Roll(random *rand.Rand) (float64, error) {
	if !s.valid {
		panic("valid == false but roll called")
	}
	ret := 0
	for _, die := range s.dice {
		for i := 0; i < die.number; i++ {
			ret += random.Intn(die.sides) + 1
		}
	}
	ret += s.additional
	return float64(ret), nil
}
