package gaussian

import (
	"errors"
	"fmt"
	"math"
	"math/rand"
)

type Specification struct {
	Mean   float64
	Stddev float64
	Min    float64
	Max    float64
	valid  bool
}

func DefaultSpecification() *Specification {
	return &Specification{
		Mean:   0,
		Stddev: 1,
		Min:    math.MinInt64,
		Max:    math.MaxFloat64,
	}
}

func (s *Specification) Validate() error {
	if s.Max <= s.Min {
		return errors.New("max is greater than min")
	}
	if s.Mean < s.Min {
		//return errors.New("mean is less than min")
		fmt.Println("mean is less than min. multiple samples may be needed")
	}
	if s.Mean > s.Max {
		//return errors.New("mean is greater than max")
		fmt.Println("mean is greater than max. multiple samples may be needed")
	}
	minDiff := s.Mean - s.Min
	maxDiff := s.Mean - s.Max
	if math.Signbit(minDiff) == math.Signbit(maxDiff) && min(math.Abs(minDiff), math.Abs(maxDiff)) > s.Stddev {
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
		potential := random.NormFloat64()*s.Stddev + s.Mean
		if potential <= s.Max && potential >= s.Min {
			return potential, nil
		}
	}
	return 0, errors.New("max samples exceeded")
}
