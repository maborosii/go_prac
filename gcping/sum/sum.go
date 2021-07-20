package sum

func Sum(a []float64) float64 {
	var summary float64 = 0.00
	for _, i := range a {
		summary += i
	}
	return summary
}

func Avg(a []float64) float64 {
	lens_a := len(a)
	summary_a := Sum(a)
	if lens_a == 0 {
		return 0.00
	}
	return summary_a / float64(lens_a)
}
