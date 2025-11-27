package dice

import (
	"errors"
	"regexp"
	"strconv"
)

func ParseDiceString(diceString string) (*Specification, error) {
	ret := &Specification{}
	tokenRegex := regexp.MustCompile("[+-]?[^+-]+")
	tokens := tokenRegex.FindAllString(diceString, -1)
	if len(tokens) == 0 {
		return nil, errors.New("no dice tokens found in query string")
	}
	diceRegex := regexp.MustCompile("^([+-]?)\\W*(\\d+)(d(\\d+))?\\W*$")

	seenDiceSides := make(map[int]*Die)
	for _, token := range tokens {
		matches := diceRegex.FindStringSubmatch(token)
		if matches == nil {
			return nil, errors.New("unable to parse token: " + token)
		}
		number, err := strconv.Atoi(matches[2])
		if err != nil {
			return nil, err
		}
		if matches[1] == "-" {
			number = number * -1
		}
		if len(matches[3]) > 0 {
			sides, err := strconv.Atoi(matches[4])
			if err != nil {
				return nil, err
			}
			die, exists := seenDiceSides[sides]
			if !exists {
				die = &Die{}
				die.sides = sides
				die.number = number
				seenDiceSides[sides] = die
			} else {
				die.number += number
			}
		} else {
			ret.additional += number
		}
	}
	for _, die := range seenDiceSides {
		ret.dice = append(ret.dice, *die)
	}
	return ret, nil
}
