package fibonacci

import (
	"context"
	"fmt"
	"github.com/ananasjesus/fibonacci/pkg/cache"
	fibonacciHelper "github.com/ananasjesus/fibonacci/pkg/fibonacci"
	pb "github.com/ananasjesus/fibonacci/proto"
	"strconv"
)

type GRPCServer struct {
	Cache cache.Cache
	pb.UnimplementedFibonacciServer
}

func (s *GRPCServer) Fibonacci(con context.Context, r *pb.FibonacciRequest) (*pb.FibonacciResponse, error) {
	if !r.Validate() {
		// Надо бы добавить отдельное поле под сообщение об ошибке
		// и отдельный валидатор, так как в сгенерированных файлах он затрётся, а в папке proto ему не место
		return &pb.FibonacciResponse{From: "validation error"}, nil
	}

	from, _ := strconv.Atoi(r.From)
	to, _ := strconv.Atoi(r.To)

	fibonacci, err := fibonacciHelper.GetCachedFibonacci(s.Cache)

	if err != nil || len(fibonacci) < to {
		fibonacci = fibonacciHelper.CalcFibonacci(to)
		fibonacciHelper.SaveFibbonacciToCache(s.Cache, fibonacci)
	}

	return &pb.FibonacciResponse{
		From:    fmt.Sprint(from),
		To:      fmt.Sprint(to),
		Numbers: fibonacci[from-1 : to],
	}, nil
}
