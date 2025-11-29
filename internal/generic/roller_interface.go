package generic

import "math/rand"

type Roller interface {
	Roll(*rand.Rand) (float64, error)
	Validate() error
}
