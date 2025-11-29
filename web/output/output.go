package output

import (
	"fmt"
	"github.com/google/uuid"
)

type Outcome struct {
	FloatValue  float64   `json:"floatValue"`
	StringValue string    `json:"stringValue"`
	RawValues   []float64 `json:"rawValues"`
	KeptValues  []float64 `json:"keptValues"`
	Seed        int64     `json:"seed"`
	Ref         string    `json:"ref"`
}

func (o *Outcome) LogRef() {
	o.Ref = uuid.New().String()
	fmt.Println(fmt.Sprintf(
		"ref created: %s. seed: %d. value: %f. choice: %s.",
		o.Ref,
		o.Seed,
		o.FloatValue,
		o.StringValue,
	))
}
