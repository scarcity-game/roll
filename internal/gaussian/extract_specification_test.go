package gaussian

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestExtractGaussianSpecification(t *testing.T) {
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
				url: "/sampleGaussian?min=1&max=2&mean=0&stddev=0",
			},
			want: &Specification{
				Min:    1,
				Max:    2,
				Mean:   0,
				Stddev: 0.0,
			},
			wantErr: assert.NoError,
		},
		{
			name: "non-numeric min",
			args: args{
				url: "/sampleGaussian?min=hello&max=2&mean=0&stddev=0",
			},
			wantErr: assert.Error,
		},
		{
			name: "non-numeric max",
			args: args{
				url: "/sampleGaussian?min=0&max=tyui&mean=0&stddev=0",
			},
			wantErr: assert.Error,
		},
		{
			name: "non-numeric mean",
			args: args{
				url: "/sampleGaussian?min=1&max=2&mean=qweqew&stddev=0",
			},
			wantErr: assert.Error,
		},
		{
			name: "non-numeric stddev",
			args: args{
				url: "/sampleGaussian?min=1&max=2&mean=0&stddev=thjk",
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

			got, err := ExtractGaussianSpecification(c)
			tt.wantErr(t, err, fmt.Sprintf("ExtractGaussianSpecification() err = %v", err))
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractGaussianSpecification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
