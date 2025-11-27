package roll

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/scarcity-game/roll/web/json"
	"math/rand"
	"sort"
)

type Roller interface {
	Roll(*rand.Rand) (float64, error)
	Validate() error
}

type Specification struct {
	Rolls           int
	Keep            int
	KeepCriteria    RollKeepCriteria
	RollAggregation Aggregation
	Seed            int64
}

const MaxRolls = 100

func (s *Specification) Validate() error {
	if s.Rolls < s.Keep {
		return errors.New("unable to keep more rolls than the number rolled")
	}
	if s.Seed == 0 {
		return errors.New("invalid seed value")
	}
	if s.Rolls < 1 {
		return errors.New("rolls must be greater than zero")
	}
	if s.Rolls > MaxRolls {
		return errors.New("rolls is too large")
	}
	if s.Keep < 1 {
		return errors.New("unable to keep less than one roll")
	}
	if s.Keep > 1 && s.RollAggregation == None {
		return errors.New("rollAggregation must be provided when multiple values are kept")
	}
	return nil
}

func (s *Specification) Roll(roller Roller) (*json.Outcome, error) {
	random := rand.New(rand.NewSource(s.Seed))
	if err := s.Validate(); err != nil {
		return nil, err
	}
	if err := roller.Validate(); err != nil {
		return nil, err
	}
	outcome := &json.Outcome{}
	outcome.Seed = s.Seed
	outcome.RawValues = make([]float64, s.Rolls)
	for i := 0; i < s.Rolls; i++ {
		rawValue, err := roller.Roll(random)
		if err != nil {
			return nil, err
		}
		outcome.RawValues[i] = rawValue
	}
	s.applyKeep(outcome)
	switch s.RollAggregation {
	case None:
		outcome.Value = outcome.KeptValues[0]
	case Average:
		outcome.Value = average(outcome.KeptValues)
	}
	outcome.Ref = uuid.New().String()
	fmt.Println(fmt.Sprintf("ref created: %s. seed: %d. value: %f.", outcome.Ref, outcome.Seed, outcome.Value))
	return outcome, nil
}

func (s *Specification) applyKeep(outcome *json.Outcome) {
	outcome.KeptValues = make([]float64, s.Keep)
	if len(outcome.RawValues) == 1 && s.Keep == 1 {
		outcome.KeptValues[0] = outcome.RawValues[0]
		return
	}
	temp := make([]float64, len(outcome.RawValues)) // Create a destination slice of the same length
	copy(temp, outcome.RawValues)
	sort.Float64s(temp)
	switch s.KeepCriteria {
	case Highest:
		outcome.KeptValues = temp[len(temp)-s.Keep:]
	case Lowest:
		outcome.KeptValues = temp[:s.Keep]
	case Middle:
		start := len(temp)/2 - s.Keep/2
		outcome.KeptValues = temp[start : start+s.Keep]
	}
}

func average(values []float64) float64 {
	total := 0.0
	for _, val := range values {
		total += val
	}
	return total / float64(len(values))
}

func DefaultSpecification() *Specification {
	return &Specification{
		Rolls:           1,
		Keep:            1,
		RollAggregation: None,
		KeepCriteria:    Highest,
		Seed:            rand.Int63(),
	}
}
