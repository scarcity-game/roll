package uniform

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSpecification_Validate(t *testing.T) {
	tests := []struct {
		name          string
		specification *Specification
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name:          "valid",
			specification: DefaultSpecification(),
			wantErr:       assert.NoError,
		},
		{
			name:          "min > max",
			specification: &Specification{Min: 3, Max: 2},
			wantErr:       assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.wantErr(t, tt.specification.Validate(), fmt.Sprintf("Validate()"))
		})
	}
}

func TestSpecification_Roll(t *testing.T) {
	type args struct {
		random *rand.Rand
	}
	tests := []struct {
		name          string
		specification *Specification
		args          args
		want          float64
		delta         float64
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			specification: &Specification{
				Min: 4,
				Max: 10,
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:    9.671176895764699,
			delta:   .00000001,
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.specification.valid = true
			got, err := tt.specification.Roll(tt.args.random)
			tt.wantErr(t, err, fmt.Sprintf("Roll()"))
			if err != nil {
				return
			}
			assert.InDelta(t, tt.want, got, tt.delta)
		})
	}
}
