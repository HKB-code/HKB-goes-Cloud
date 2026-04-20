// PART-1
/*
// Every Go program is made up of packages.Programs start running in package main.
package main

// factored import statement.
import (
	"fmt"
	"math"
	"math/rand"
)
// Constants are flexible until used; variables are strict from the start.
const PI = 3.14
// The var statement declares a list of variables; as in function argument lists, the type is last.
// A var statement can be at package or function level.

var c,python,java bool
// A var declaration can include initializers, one per variable.If an initializer is present, the type can be omitted; the variable will take the type of the initializer.
var i,j = 1,2
func main(){
	fmt.Println(rand.Intn(10))
	//
	// When importing a package, you can refer only to its exported names. Any "unexported" names are not accessible from outside the package.
	//
	fmt.Println(math.Pi)
	fmt.Println(add(10,2))
	fmt.Println(swap("Hello","World"))
	fmt.Println(split(17))
	var i int
	fmt.Println(i, c, python, java)

	// short variable declaration:
	k:= 3
	fmt.Println(k)

	// Variables declared without an explicit initial value are given their zero value.
	var (
		j int
		l float64
		p string
	)
	fmt.Println(j,l,p)

	// Type conversion: Unlike in C, in Go assignment between items of different type requires an explicit conversion. Try removing the float64 or uint conversions in the example and see what happens.
	var f float64 = float64(i)
	fmt.Println(f)
}

func add(x,y int)int{
	z:=x+y
	return z
}

// A function can return any number of results.
func swap(x,y string)(string,string){
	return y,x
}
// A return statement without arguments returns the named return values. This is known as a "naked" return.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}


*/

// Part-2
/*
package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
) 

func pow(x,n,lim float64)float64{
	// Variables declared inside an if short statement are also available inside any of the else blocks.
	if v:= math.Pow(x,n);v<lim{
		return  v 
	}else{
		fmt.Printf("%g >= %g\n",v,lim)
	}
	// can't use v here, though
	return  lim
}

func main() {
	for i:= 0;i<5;i++{
		fmt.Println(i)
	}
	sum := 1
	for sum<10{
		sum+=sum
	}
	fmt.Println(sum)

		fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 20),
	)

	// Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

	switch os:= runtime.GOOS;os{
		case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switch without a condition is the same as switch true.
// This construct can be a clean way to write long if-then-else chains.

t:= time.Now()
switch{
case t.Hour()<12:
	fmt.Println("Good morning!")
case t.Hour()<17:
	fmt.Println("Good afternoon.")
default:
	fmt.Println("Good evening.")

}

// A defer statement defers the execution of a function until the surrounding function returns.

// The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
defer fmt.Println("world")

	fmt.Println("hello")


	// Stacking defers
// Deferred function calls are pushed onto a stack. When a function returns, its deferred calls are executed in last-in-first-out order.

fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}

*/

package main  





