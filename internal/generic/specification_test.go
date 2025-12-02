package generic

import (
	"errors"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/scarcity-game/roll/web/output"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestSpecification_Validate(t *testing.T) {
	type fields struct {
		Rolls           int
		Keep            int
		KeepCriteria    RollKeepCriteria
		RollAggregation Aggregation
		Seed            int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "valid",
			fields: fields{
				Rolls:           1,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.NoError,
		},

		{
			name: "too many rolls",
			fields: fields{
				Rolls:           101,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.Error,
		},
		{
			name: "keep more than rolled",
			fields: fields{
				Rolls:           1,
				Keep:            2,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.Error,
		},
		{
			name: "invalid seed",
			fields: fields{
				Rolls:           1,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            0,
			},
			wantErr: assert.Error,
		},
		{
			name: "keep zero",
			fields: fields{
				Rolls:           1,
				Keep:            0,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.Error,
		},
		{
			name: "zero rolls",
			fields: fields{
				Rolls:           0,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.Error,
		},
		{
			name: "negative rolls",
			fields: fields{
				Rolls:           -1,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.Error,
		},
		{
			name: "keep multiple without aggregation",
			fields: fields{
				Rolls:           10,
				Keep:            2,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Specification{
				Rolls:           tt.fields.Rolls,
				Keep:            tt.fields.Keep,
				KeepCriteria:    tt.fields.KeepCriteria,
				RollAggregation: tt.fields.RollAggregation,
				Seed:            tt.fields.Seed,
			}
			tt.wantErr(t, s.Validate(), fmt.Sprintf("Validate()"))
		})
	}
}

type testRoller struct {
	rolls []float64
	err   error
	curr  int
}

func (t *testRoller) Roll(rand *rand.Rand) (float64, error) {
	val := t.rolls[t.curr]
	t.curr++
	return val, nil
}

func (t *testRoller) Validate() error {
	return t.err
}

func TestSpecification_Roll(t *testing.T) {
	type fields struct {
		Rolls           int
		Keep            int
		KeepCriteria    RollKeepCriteria
		RollAggregation Aggregation
		Seed            int64
	}
	type args struct {
		roller Roller
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *output.Outcome
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			fields: fields{
				Rolls:           1,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   nil,
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1},
				KeptValues: []float64{1},
				FloatValue: 1,
				Seed:       1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "unable to validate spec",
			fields: fields{
				Rolls:           0,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   nil,
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1},
				KeptValues: []float64{1},
				FloatValue: 1,
				Seed:       1,
			},
			wantErr: assert.Error,
		},
		{
			name: "unable to validate roller",
			fields: fields{
				Rolls:           1,
				Keep:            1,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   errors.New("error"),
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1},
				KeptValues: []float64{1},
				FloatValue: 1,
				Seed:       1,
			},
			wantErr: assert.Error,
		},
		{
			name: "multiple rolls",
			fields: fields{
				Rolls:           2,
				Keep:            1,
				KeepCriteria:    Lowest,
				RollAggregation: None,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   nil,
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1, 2},
				KeptValues: []float64{1},
				FloatValue: 1,
				Seed:       1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "multiple rolls, keep middle two",
			fields: fields{
				Rolls:           4,
				Keep:            2,
				KeepCriteria:    Middle,
				RollAggregation: Average,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   nil,
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1, 2, 3, 4},
				KeptValues: []float64{2, 3},
				FloatValue: 2.5,
				Seed:       1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "multiple rolls, keep top two",
			fields: fields{
				Rolls:           4,
				Keep:            2,
				KeepCriteria:    Highest,
				RollAggregation: Average,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   nil,
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1, 2, 3, 4},
				KeptValues: []float64{3, 4},
				FloatValue: 3.5,
				Seed:       1,
			},
			wantErr: assert.NoError,
		},
		{
			name: "multiple rolls, keep bottom two",
			fields: fields{
				Rolls:           4,
				Keep:            2,
				KeepCriteria:    Lowest,
				RollAggregation: Average,
				Seed:            1,
			},
			args: args{
				roller: &testRoller{
					rolls: []float64{1, 2, 3, 4},
					err:   nil,
				},
			},
			want: &output.Outcome{
				RawValues:  []float64{1, 2, 3, 4},
				KeptValues: []float64{1, 2},
				FloatValue: 1.5,
				Seed:       1,
			},
			wantErr: assert.NoError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Specification{
				Rolls:           tt.fields.Rolls,
				Keep:            tt.fields.Keep,
				KeepCriteria:    tt.fields.KeepCriteria,
				RollAggregation: tt.fields.RollAggregation,
				Seed:            tt.fields.Seed,
			}
			got, err := s.Roll(tt.args.roller)
			if !tt.wantErr(t, err, fmt.Sprintf("Roll(%v)", tt.args.roller)) {
				return
			}
			if err != nil {
				return
			}
			if !cmp.Equal(tt.want, got, cmpopts.IgnoreFields(output.Outcome{}, "Ref")) {
				t.Errorf("Roll() mismatch")
			}
		})
	}
}
