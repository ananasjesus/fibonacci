package handler

import (
	fibonacciHelper "github.com/ananasjesus/fibonacci/pkg/fibonacci"
	requests "github.com/ananasjesus/fibonacci/pkg/fibonacci/request"
	responses "github.com/ananasjesus/fibonacci/pkg/fibonacci/response"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

const validationErrorMessage = "Validation error: 'from' must be greater then 0 and 'to' must be between 'from' and 90"

func (h *Handler) fibonacci(c *gin.Context) {
	request := &requests.FibonacciRequest{
		From: c.Query("from"),
		To:   c.Query("to"),
	}

	if !request.Validate() {
		c.JSON(http.StatusOK, struct {
			Error string `json:"error"`
		}{validationErrorMessage})
		return
	}

	from, _ := strconv.Atoi(request.From)
	to, _ := strconv.Atoi(request.To)

	fibonacci, err := fibonacciHelper.GetCachedFibonacci(h.cache)

	if err != nil || len(fibonacci) < to {
		fibonacci = fibonacciHelper.CalcFibonacci(to)
		fibonacciHelper.SaveFibbonacciToCache(h.cache, fibonacci)
	}

	c.JSON(http.StatusOK, responses.FibonacciResponse{
		From:    from,
		To:      to,
		Numbers: fibonacci[from-1 : to],
	})
}
