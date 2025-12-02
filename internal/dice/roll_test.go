package dice

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSpecification_Validate(t *testing.T) {
	type fields struct {
		dice       []Die
		additional int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid",
			fields: fields{
				dice: []Die{{
					number: 1,
					sides:  6,
				}},
				additional: 1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "zero dice",
			fields: fields{
				dice: []Die{{
					number: 0,
					sides:  6,
				}},
				additional: 1,
			},
			wantErr: assert.Error,
		},
		{
			name: "zero dice",
			fields: fields{
				dice: []Die{{
					number: 0,
					sides:  6,
				}},
				additional: 1,
			},
			wantErr: assert.Error,
		},
		{
			name: "zero sides",
			fields: fields{
				dice: []Die{{
					number: 1,
					sides:  0,
				}},
				additional: 1,
			},
			wantErr: assert.Error,
		},
		{
			name: "too many sides",
			fields: fields{
				dice: []Die{{
					number: 1,
					sides:  2 << 32,
				}},
				additional: 1,
			},
			wantErr: assert.Error,
		},
		{
			name: "too many dice",
			fields: fields{
				dice: []Die{{
					number: 101,
					sides:  6,
				}},
				additional: 1,
			},
			wantErr: assert.Error,
		},
		{
			name: "too many dice types",
			fields: fields{
				dice:       make([]Die, 101),
				additional: 1,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Specification{
				dice:       tt.fields.dice,
				additional: tt.fields.additional,
			}
			tt.wantErr(t, s.Validate(), fmt.Sprintf("Validate()"))
		})
	}
}

func TestSpecification_Roll(t *testing.T) {
	type fields struct {
		dice       []Die
		additional int
	}
	type args struct {
		random *rand.Rand
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   float64
		delta  float64
	}{
		{
			name: "100d1 dice",
			fields: fields{
				dice: []Die{
					{
						number: 100,
						sides:  1,
					},
				},
				additional: 0,
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:  100,
			delta: 0,
		},
		{
			name: "additional 1",
			fields: fields{
				dice:       []Die{},
				additional: 1,
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:  1,
			delta: 0,
		},
		{
			name: "1000d6 dice",
			fields: fields{
				dice: []Die{
					{
						number: 1000,
						sides:  6,
					},
				},
				additional: 0,
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:  3401,
			delta: 0,
		},
		{
			name: "100d2 dice",
			fields: fields{
				dice: []Die{
					{
						number: 100,
						sides:  2,
					},
				},
				additional: 0,
			},
			args: args{
				random: rand.New(rand.NewSource(0)),
			},
			want:  136,
			delta: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Specification{
				dice:       tt.fields.dice,
				additional: tt.fields.additional,
				valid:      true,
			}
			if got, err := s.Roll(tt.args.random); err != nil {
				t.Error(err)
			} else {
				assert.InDelta(t, tt.want, got, tt.delta)
			}
		})
	}
}
