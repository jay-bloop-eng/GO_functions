package fibonacci

import (
	"fmt"
)

func Fibonacci(n int) {
	var (
		f1       int
		f2       int = 1
		num_fibo int
	)
	
	fmt.Printf("os numeros da sequencia de fibonacci a %d posicoes sao: ", n)
	fmt.Printf("%d, %d, ", f1, f2)
	for i := 1; i <= n; i++ {
		num_fibo = f1 + f2
		f1 = f2
		f2 = num_fibo
		fmt.Printf("%d, ", num_fibo)
	}
	fmt.Printf("\n")
}
