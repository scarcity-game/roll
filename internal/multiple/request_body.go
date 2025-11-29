package multiple

import "github.com/scarcity-game/roll/internal/weighted"

type RequestType string

const (
	Dice     RequestType = "dice"
	Gaussian RequestType = "gaussian"
	Uniform  RequestType = "uniform"
	Weighted RequestType = "weighted"
)

type Request struct {
	RequestType RequestType       `json:"requestType"`
	DiceString  string            `json:"dice"`
	Min         float64           `json:"min"`
	Max         float64           `json:"max"`
	Mean        float64           `json:"mean"`
	Stddev      float64           `json:"stddev"`
	Choices     []weighted.Choice `json:"choices"`
}

type Body struct {
	Seed     string `json:"seed"`
	Requests []Request
}
