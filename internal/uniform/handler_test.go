package uniform

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSampleUniform(t *testing.T) {
	t.Run("pre-seeded", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/sampleUniform?max=100&min=10&rolls=50&keep=3&keepCriteria=middle&aggregation=average&seed=abc123abc", nil)
		c.Request = req
		SampleUniform(c)
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.InDelta(t, 73.776952, responseBody["value"], 0.0001)
	})
}
