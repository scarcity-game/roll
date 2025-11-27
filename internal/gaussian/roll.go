package gaussian

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type Specification struct {
	mean   float64
	stddev float64
	min    float64
	max    float64
	valid  bool
}

func defaultSpecification() *Specification {
	return &Specification{
		mean:   0,
		stddev: 1,
		min:    math.MinInt64,
		max:    math.MaxFloat64,
	}
}

func (s *Specification) Validate() error {
	if s.max <= s.min {
		return errors.New("max is greater than min")
	}
	if s.mean < s.min {
		//return errors.New("mean is less than min")
		fmt.Println("mean is less than min. multiple samples may be needed")
	}
	if s.mean > s.max {
		//return errors.New("mean is greater than max")
		fmt.Println("mean is greater than max. multiple samples may be needed")
	}
	minDiff := s.mean - s.min
	maxDiff := s.mean - s.max
	if math.Signbit(minDiff) == math.Signbit(maxDiff) && min(math.Abs(minDiff), math.Abs(maxDiff)) > s.stddev {
		//return errors.New("mean is greater than max")
		fmt.Println("acceptable range more than a stddev from mean. multiple samples may be needed")
	}
	s.valid = true
	return nil
}

const maxSamples = 100

func (s *Specification) Roll(random *rand.Rand) (float64, error) {
	if !s.valid {
		panic("valid == false but roll called")
	}
	for i := 0; i < maxSamples; i++ {
		potential := random.NormFloat64()*s.stddev + s.mean
		if potential <= s.max && potential >= s.min {
			return potential, nil
		}
	}
	return 0, errors.New("max samples exceeded")
}
