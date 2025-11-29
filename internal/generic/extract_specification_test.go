package generic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestExtractSpecification(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    *Specification
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "happy path",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=middle&aggregation=average&seed=abc123abc",
			},
			want: &Specification{
				Rolls:           50,
				Keep:            3,
				KeepCriteria:    Middle,
				RollAggregation: Average,
				Seed:            46104984252,
			},
			wantErr: assert.NoError,
		},
		{
			name: "none aggregation",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=middle&aggregation=none&seed=abc123abc",
			},
			want: &Specification{
				Rolls:           50,
				Keep:            3,
				KeepCriteria:    Middle,
				RollAggregation: None,
				Seed:            46104984252,
			},
			wantErr: assert.NoError,
		},
		{
			name: "highest keep criteria",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=highest&aggregation=none&seed=abc123abc",
			},
			want: &Specification{
				Rolls:           50,
				Keep:            3,
				KeepCriteria:    Highest,
				RollAggregation: None,
				Seed:            46104984252,
			},
			wantErr: assert.NoError,
		},
		{
			name: "lowest keep criteria",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=lowest&aggregation=none&seed=abc123abc",
			},
			want: &Specification{
				Rolls:           50,
				Keep:            3,
				KeepCriteria:    Lowest,
				RollAggregation: None,
				Seed:            46104984252,
			},
			wantErr: assert.NoError,
		},
		{
			name: "non-numeric rolls",
			args: args{
				url: "/rollDice?rolls=hello&keep=3&keepCriteria=middle&aggregation=average&seed=abc123abc",
			},
			wantErr: assert.Error,
		},
		{
			name: "non-numeric keep",
			args: args{
				url: "/rollDice?rolls=50&keep=hiya&keepCriteria=middle&aggregation=average&seed=abc123abc",
			},
			wantErr: assert.Error,
		},
		{
			name: "bad keep criteria",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=everyOther&aggregation=average&seed=abc123abc",
			},
			wantErr: assert.Error,
		},
		{
			name: "bad aggregate",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=average&aggregation=mode&seed=abc123abc",
			},
			wantErr: assert.Error,
		},
		{
			name: "bad seed",
			args: args{
				url: "/rollDice?rolls=50&keep=3&keepCriteria=average&aggregation=average&seed=yuio",
			},
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gin.SetMode(gin.TestMode)
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			req, _ := http.NewRequest("GET", tt.args.url, nil)
			c.Request = req

			got, err := ExtractSpecification(c)
			tt.wantErr(t, err, fmt.Sprintf("ExtractSpecification() err = %v", err))
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractSpecification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
