package fibonacci

import (
	"encoding/json"
	"github.com/ananasjesus/fibonacci/pkg/cache"
	"log"
)

func CalcFibonacci(n int) []int64 {
	fibonacciNumbers := make([]int64, n+1)
	fibonacciNumbers[0] = 0
	fibonacciNumbers[1] = 1

	for i := 2; i < n; i++ {
		fibonacciNumbers[i] = fibonacciNumbers[i-1] + fibonacciNumbers[i-2]
	}

	return fibonacciNumbers[0:n]
}

// Так как из-за размера переменной установлено ограничение в 90 чисел, оптимально хранить всё под одним ключом
func GetCachedFibonacci(cache cache.Cache) ([]int64, error) {
	var slice []int64
	err := json.Unmarshal([]byte(cache.Get("fibonacci")), &slice)
	return slice, err
}

func SaveFibbonacciToCache(cache cache.Cache, slice []int64) {
	jsonFibonacci, err := json.Marshal(slice)
	if err != nil {
		log.Printf("error occured while marshal: %s", err.Error())
		return
	}
	err = cache.Set("fibonacci", jsonFibonacci, 0)
	if err != nil {
		log.Printf("error occured while set cache: %s", err.Error())
	}
}
