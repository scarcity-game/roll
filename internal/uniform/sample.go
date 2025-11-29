package uniform

import (
	"errors"
	"math/rand"
)

type Specification struct {
	Min   float64
	Max   float64
	valid bool
}

func DefaultSpecification() *Specification {
	return &Specification{
		Min: 0,
		Max: 1,
	}
}

func (s *Specification) Validate() error {
	if s.Max <= s.Min {
		return errors.New("max is greater than min")
	}
	s.valid = true
	return nil
}

func (s *Specification) Roll(random *rand.Rand) (float64, error) {
	if !s.valid {
		panic("valid == false but roll called")
	}
	sampleRange := s.Max - s.Min
	return random.Float64()*sampleRange + s.Min, nil
}
