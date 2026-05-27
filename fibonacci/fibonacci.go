package fibonacci

import (
	"fmt"
)

func fibonacci(n int) {
	var (
		f1       int
		f2       int
		num_fibo int
	)
	fmt.Printf("quantos números da sequência de fibonacci você quer? \n")
	fmt.Scanf("%d", &n)
	f2 = 1
	fmt.Printf("os numeros da sequencia de fibonacci a %d posicoes sao: ", n)
	fmt.Printf("0, 1, ")
	for i := 1; i <= n; i++ {
		num_fibo = f1 + f2
		f1 = f2
		f2 = num_fibo
		fmt.Printf("%d, ", num_fibo)
	}
	fmt.Printf("\n")
}
