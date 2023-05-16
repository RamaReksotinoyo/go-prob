package go_prob

func Prob() string {
	return "This is probability module."
}

type Probability struct {
	Value float64
}

// instance dari object probability
func InstProbability(value float64) *Probability {
	if value < 0 || value > 1 {
		panic("Probability value must be between 0 and 1.")
	}
	return &Probability{Value: value}
}

func CalculateProbability(event *float64, sampleSpace *float64) *Probability {
	if *sampleSpace == 0 {
		panic("Sample space cannot be zero.")
	}

	value := *event / *sampleSpace
	return &Probability{Value: value}
}