// 1.Exercise: Fibonacci closure
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func Fibonacci()func()int{
// 	a,b:=0,1

// 	return  func()int{
// 		res:= a
// 		a,b = b ,a+b
// 		return  res
// 	}

// }

// func main() {
// f:= Fibonacci()
// 		for i := 0; i < 10; i++ {
// 		fmt.Println(f())
// 	}
// }

// 2.wordscount
package main

import (
	"fmt"
	"strings"
) 
func wordsCount(s string)map[string]int{
	a:= strings.Fields(s)
	z:= map[string]int{}
	for _,v:= range a{
		z[v]++
	}
	return z
}

func main() {
	g:=wordsCount("Hello Hi Hey Whoo Hello")
fmt.Println(g)
}