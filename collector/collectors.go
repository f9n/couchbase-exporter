package collector

const namespace = "couchbase"

func fromBool(i bool) float64 {
	if i {
		return 1.0
	}
	return 0.0
}

func last(ss []float64) float64 {
	return ss[len(ss)-1]
}
