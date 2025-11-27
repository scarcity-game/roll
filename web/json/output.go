package json

type Outcome struct {
	Value      float64   `json:"value"`
	RawValues  []float64 `json:"rawValues"`
	KeptValues []float64 `json:"keptValues"`
	Seed       int64     `json:"seed"`
	Ref        string    `json:"ref"`
}
