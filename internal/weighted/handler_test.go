package weighted

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/scarcity-game/roll/web/output"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMakeWeightedChoice(t *testing.T) {
	t.Run("pre-seeded", func(t *testing.T) {
		gin.SetMode(gin.TestMode)
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)
		choices := Specification{
			Choices: []Choice{
				{
					Weight: 1234,
					Value:  "option1",
				},
				{
					Weight: 1.234,
					Value:  "option2",
				},
				{
					Weight: 123,
					Value:  "option3",
				},
			},
			StringSeed: "deadbeef1234",
		}
		if body, err := json.Marshal(choices); err != nil {
			t.Fatal(err)
		} else if req, err := http.NewRequest("GET", "/weightedChoice", bytes.NewBuffer(body)); err != nil {
			t.Fatal(err)
		} else {
			c.Request = req
		}
		MakeWeightedChoice(c)
		var responseBody output.Outcome
		err := json.Unmarshal(w.Body.Bytes(), &responseBody)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equalf(t, "option3", responseBody.StringValue, "MakeWeightedChoice()")
		assert.Equalf(t, int64(244837814047284), responseBody.Seed, "MakeWeightedChoice()")
	})
}
