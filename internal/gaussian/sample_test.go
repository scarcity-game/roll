package gaussian

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
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
			name:          "min>max",
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
			name:          "mean=0 stddev=1",
			specification: DefaultSpecification(),
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:    -0.28158587086436215,
			delta:   .000001,
			wantErr: assert.NoError,
		},
		{
			name: "outlier by 4 stddev",
			specification: &Specification{
				Mean:   0,
				Stddev: 1,
				Min:    4,
				Max:    math.MaxFloat64,
			},
			args: args{
				random: rand.New(rand.NewSource(2358639785634891464)),
			},
			want:    4.050253,
			delta:   .000001,
			wantErr: assert.NoError,
		},

		{
			name: "max exceeded",
			specification: &Specification{
				Mean:   0,
				Stddev: 1,
				Min:    100,
				Max:    math.MaxFloat64,
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			wantErr: assert.Error,
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
