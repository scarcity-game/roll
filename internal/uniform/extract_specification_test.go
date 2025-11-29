package uniform

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestExtractUniformSpecification(t *testing.T) {
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
				url: "/sampleGaussian?min=1&max=2",
			},
			want: &Specification{
				Min: 1,
				Max: 2,
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
				url: "/sampleGaussian?min=0&max=anything&mean=0&stddev=0",
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

			got, err := ExtractUniformSpecification(c)
			tt.wantErr(t, err, fmt.Sprintf("ExtractUniformSpecification() err = %v", err))
			if err != nil {
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractUniformSpecification() got = %v, want %v", got, tt.want)
			}
		})
	}
}
