package fibonacci

type FibonacciResponse struct {
	From    int     `json:"from"`
	To      int     `json:"to"`
	Numbers []int64 `json:"numbers"`
}
