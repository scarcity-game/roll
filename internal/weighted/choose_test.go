package weighted

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
			name: "valid",
			specification: &Specification{
				Choices: []Choice{{
					Weight: 100,
					Value:  "test",
				}},
			},
			wantErr: assert.NoError,
		},
		{
			name: "zero weight option",
			specification: &Specification{
				Choices: []Choice{{
					Weight: 100,
					Value:  "test",
				},
					{
						Weight: 0,
						Value:  "zero weight",
					}},
			},
			wantErr: assert.NoError,
		},
		{
			name: "zero choices ",
			specification: &Specification{
				Choices: []Choice{},
			},
			wantErr: assert.Error,
		},
		{
			name: "zero total weight",
			specification: &Specification{
				Choices: []Choice{{
					Weight: 0,
					Value:  "test",
				}},
			},
			wantErr: assert.Error,
		},
		{
			name: "negative weight",
			specification: &Specification{
				Choices: []Choice{{
					Weight: -100,
					Value:  "test",
				}},
			},
			wantErr: assert.Error,
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
		want          string
		wantErr       assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			specification: &Specification{
				Choices: []Choice{{
					Weight: 100,
					Value:  "test",
				}},
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:    "test",
			wantErr: assert.NoError,
		},
		{
			name: "happy path multiple options",
			specification: &Specification{
				Choices: []Choice{{
					Weight: 100,
					Value:  "option1",
				}, {
					Weight: 101,
					Value:  "option2",
				}, {
					Weight: 5.12345,
					Value:  "option3",
				}, {
					Weight: 10.5,
					Value:  "option4",
				}, {
					Weight: 1.2,
					Value:  "option5",
				}},
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:    "option3",
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.specification.valid = true
			err := tt.specification.Validate()
			if err != nil {
				panic(err)
			}
			got, err := tt.specification.Roll(tt.args.random)
			tt.wantErr(t, err, fmt.Sprintf("Roll()"))
			if err != nil {
				return
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
