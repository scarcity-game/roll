package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSampleGaussian(t *testing.T) {
	t.Run("pre-seeded", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/sampleGaussian?mean=100&stddev=10&rolls=50&keep=3&keepCriteria=middle&aggregation=average&seed=abc123abc", nil)
		c.Request = req
		SampleGaussian(c)
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.InDelta(t, 98.810335, responseBody["value"], 0.0001)
	})
}
