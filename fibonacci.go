/**
 * User: guoyao
 * Time: 06-20-2013 12:16
 * Blog: http://www.guoyao.me
 */

package main

import "fmt"

func main() {
	//fmt.Println(fibonacci(40))
	fmt.Println(fibonacci(45))
}

func fibonacci(n uint) (result uint) {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}
