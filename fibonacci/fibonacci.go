package fibonacci

import (
	"fmt"
	"net/http"
	"strconv"
)

// Calculate computes the nth Fibonacci number using an iterative approach
// This is much more efficient than the recursive approach with O(n) time complexity
// instead of O(2^n)

// func fibo(n int) int {
// 	if n <= 1 {
// 		return n
// 	}

// 	// Use iterative approach to avoid stack overflow and improve performance
// 	a, b := 0, 1
// 	for i := 2; i <= n; i++ {
// 		a, b = b, a+b
// 	}
// 	return b
// }

func Calculate(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Calculate(n-1) + Calculate(n-2)
	}
}

// Handler handles HTTP requests for Fibonacci calculations
func Handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	id := r.PathValue("id")

	if id == "" {
		// Handle the case where no ID is provided
		fmt.Fprintf(w, "Please provide a number. Fibonacci calculation request is rejected!\n")
		return
	}

	if n, err := strconv.Atoi(id); err != nil {
		fmt.Fprintf(w, "Please provide a valid number. Fibonacci calculation request is rejected!\n")
	} else {
		if n <= 45 && n >= 0 {
			fib := Calculate(n)
			fmt.Fprintf(w, "Fibonacci(%d)=%d\n", n, fib)
		} else {
			fmt.Fprintf(w, "%d is not valid. Please provide number <= 45, >= 0. Fibonacci calculation request is rejected!\n", n)
		}
	}
}
