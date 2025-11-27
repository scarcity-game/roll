package dice

import (
	"errors"
	a "github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestParseDiceString(t *testing.T) {
	type args struct {
		diceString string
	}
	tests := []struct {
		name string
		args args
		err  error
		want *Specification
	}{
		{
			name: "complex happy path",
			args: args{"1d4+2d6-3d10+42"},
			err:  nil,
			want: &Specification{
				dice: []Die{
					{
						number: 1,
						sides:  4,
					},
					{
						number: 2,
						sides:  6,
					},
					{
						number: -3,
						sides:  10,
					},
				},
				additional: 42,
			},
		},
		{
			name: "simple happy path",
			args: args{"1d4+1"},
			err:  nil,
			want: &Specification{
				dice: []Die{
					{
						number: 1,
						sides:  4,
					},
				},
				additional: 1,
			},
		},
		{
			name: "add additional",
			args: args{"1-2+4-10+20"},
			err:  nil,
			want: &Specification{
				dice:       []Die{},
				additional: 13,
			},
		},
		{
			name: "ignore whitespace on additional ",
			args: args{"1 -  2 +   4 - 10+ 20"},
			err:  nil,
			want: &Specification{
				dice:       []Die{},
				additional: 13,
			},
		},
		{
			name: "combine dice",
			args: args{"1d4+2d4+3d4"},
			err:  nil,
			want: &Specification{
				dice: []Die{
					{
						number: 6,
						sides:  4,
					},
				},
				additional: 0,
			},
		},
		{
			name: "reject non-dice",
			args: args{"hello"},
			err:  errors.New("unable to parse token: hello"),
			want: nil,
		},
		{
			name: "reject empty dice",
			args: args{""},
			err:  errors.New("no dice tokens found in query string"),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseDiceString(tt.args.diceString)
			if !reflect.DeepEqual(err, tt.err) {
				t.Errorf("ParseDiceString() err = %v, want %v", err, tt.err)
			}
			if err != nil {
				return
			}
			a.ElementsMatch(t, tt.want.dice, got.dice)
			a.Equal(t, tt.want.additional, got.additional)
		})
	}
}
