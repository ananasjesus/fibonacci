package proto

import "strconv"

func (r *FibonacciRequest) Validate() bool {
	from, err := strconv.Atoi(r.From)
	if err != nil {
		return false
	}

	to, err := strconv.Atoi(r.To)
	if err != nil {
		return false
	}

	if from > 0 && to >= from && to <= 90 {
		return true
	} else {
		return false
	}
}
