package go_prob

func Prob() string {
	return "This is probability module."
}

type Probability struct {
	Value float64
}

func CalculateProbability(event *float64, sampleSpace *float64) *Probability {
	if *sampleSpace == 0 {
		panic("Sample space cannot be zero.")
	}

	value := *event / *sampleSpace
	return &Probability{Value: value}
}




