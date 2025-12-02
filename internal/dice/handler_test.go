package dice

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/web/output"
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
		var outcome output.Outcome
		err := json.Unmarshal(w.Body.Bytes(), &outcome)
		if err != nil {
			t.Fatal(err)
		}
		assert.InDelta(t, 11.6666666, outcome.FloatValue, 0.00001)
	})
}
