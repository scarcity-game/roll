package web

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRollDice(t *testing.T) {
	t.Run("pre-seeded", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		req, _ := http.NewRequest("GET", "/rollDice?dice=1d20&rolls=50&keep=3&keepCriteria=middle&aggregation=average&seed=abc123abc", nil)
		c.Request = req
		RollDice(c)
		var responseBody map[string]interface{}
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, 11.666666666666666, responseBody["value"])
	})
}
