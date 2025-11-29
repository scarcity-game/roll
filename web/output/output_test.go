package output

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOutcome_LogRef(t *testing.T) {
	t.Run("sets ref", func(t *testing.T) {
		o := &Outcome{}
		o.LogRef()
		assert.NotNil(t, o.Ref, "ref is still nil")
	})
}
