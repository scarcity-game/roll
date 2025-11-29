package weighted

import (
	"errors"
	"math/rand"
)

type Choice struct {
	Weight float64 `json:"weight"`
	Value  string  `json:"value"`
}

type Specification struct {
	Choices []Choice `json:"choices"`
	valid   bool
}

func (s *Specification) Validate() error {
	if len(s.Choices) == 0 {
		return errors.New("no choices specified")
	}
	totalWeight := 0.0

	for _, choice := range s.Choices {
		if choice.Weight < 0 {
			return errors.New("choices must have a non-negative weight")
		}
		totalWeight += choice.Weight
	}
	if totalWeight == 0 {
		return errors.New("no weights specified")
	}
	s.valid = true
	return nil
}

func (s *Specification) Roll(random *rand.Rand) (string, error) {
	if !s.valid {
		panic("valid == false but roll called")
	}
	if len(s.Choices) == 1 {
		return s.Choices[0].Value, nil
	}
	totalWeight := 0.0
	for _, choice := range s.Choices {
		totalWeight += choice.Weight
	}
	selection := random.Float64() * totalWeight
	for _, choice := range s.Choices {
		selection -= choice.Weight
		if selection <= 0 {
			return choice.Value, nil
		}
	}
	return "", errors.New("after traversing choices, no option was found")
}
