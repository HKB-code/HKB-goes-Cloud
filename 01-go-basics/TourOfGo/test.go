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

// package main

// import (
// 	"fmt"
// 	"strings"
// )

// // A struct is a collection of fields.
// type Vertx struct{
// 	x int
// 	y int
// }

// func main() {
// 	// The type *T is a pointer to a T value. Its zero value is nil.
// 	var p *int
// 	fmt.Println(p)
// 	i:= 42
// 	// The & operator generates a pointer to its operand.
// 	p = &i
// fmt.Println(p)
// // The * operator denotes the pointer's underlying value.
// // This is known as "dereferencing" or "indirecting".
// 	fmt.Println(*p)

// 	*p = 21
// 	fmt.Println(*p)

// 	v:= Vertx{1,2}
// 	fmt.Println(v.x,v.y)

// 	z:= &v

// 	z.x = 5
// 	fmt.Println(v)

// 	(*z).y = 6
// 	fmt.Println(v)

// // A struct literal creates a new struct value by directly specifying its field values.

// // 👉 It both:
// // 1. Allocates memory
// // 2. Initializes fields

// // 🔥 Key Insight:
// // Struct literal = allocation + initialization in one step
// // Struct literal = “create + fill struct instantly”

// // 1.Positional Struct Literal:
// // 👉 Fields are filled in order:
// // ⚠️ Risk: Order must match struct definition
// v1 := Vertx{2,3}

// // 2. Named Struct Literal (Preferred ✅)
// // 👉 Benefits:
// // Order doesn’t matter
// // More readable
// // Safer for large structs
// v2 := Vertx{
// 	x:6,
// 	y:7,
// }

// //3. Partial Initialization
// // 👉 Missing fields get zero values:
// v3:= Vertx{x:9}

// // 4. Pointer to Struct (Important)
// // 👉 This creates:
// // A struct
// // Returns a pointer to it
// k:= &v1

// // 5. Anonymous Struct Literal
// // 👉 Useful for quick, one-time structures

// // 🧠 What’s happening here?

// // You are creating an anonymous struct:
// // - No type Person struct {...} defined globally
// // - Struct is defined inline
// // - Used only within this scope

// // 🔹 When should you use this?
// // ✅ Good use cases:
// // - Temporary data
// // - Quick testing
// // - Struct used only once
// // - Small programs / scripts

// // 🔴 When NOT to use
// // If you need:
// // - Reuse across functions
// // - Methods on struct
// // - Clean architecture
// l := struct{
// 	Name string
// 	Age int
// }{
// 	Name: "Jhon",
// 	Age : 44,
// }
// // ⚡ Important Limitation
// // - You cannot reuse that anonymous struct easily:
// // func printPerson(p ???) // ❌ can't reference its type, 👉 Because it has no name

// // 🔥🔥 Advanced Insight (Interview Level)
// // Even though it’s anonymous, Go still:
// // - Creates a real type
// // - But it's unique and unnamed
// // 👉 So two identical anonymous structs are considered different types

// fmt.Println(v1,v2,v3,k,l)
// p1 := struct{ Age int }{Age: 20}
// p2 := struct{ Age int }{Age: 25}

// fmt.Println(p1 == p2)

// // 🧠 Mental Model:
// //Anonymous struct = quick, one-time struct inside a function

// // Array:
// // An array's length is part of its type, so arrays cannot be resized
// var a[2] string
// a[0] = "Hello"
// a[1] = "World"
// fmt.Println(a[0],a[1])
// fmt.Println(a)

// //Slice:An array has a fixed size. A slice, on the other hand, is a dynamically-sized, flexible view into the elements of an array. In practice, slices are much more common than arrays.

// // A slice is formed by specifying two indices, a low and high bound, separated by a colon:
// b:=a[0:1]
// fmt.Println(b)

// // Slices are like references to arrays:
// // A slice does not store any data, it just describes a section of an underlying array.

// // Changing the elements of a slice modifies the corresponding elements of its underlying array.
// names := [4]string{
// 	"Jhon","Paul","George","Ringo",
// }
// fmt.Println(names)

// c:= names[0:2]
// d:= names[2:3]
// fmt.Println(c,d)
// d[0]= "bbbbb"
// fmt.Println(c,d)
// fmt.Println(names)

// // Slice literals: A slice literal is like an array literal without the length.

// r:= []bool{true,false,true,false}

// fmt.Println(r)

// s:= []struct{i int
// 	b bool}{{2,true},{3,false}}

// 	fmt.Println(s)

// 	f:= []int{2,3,4,5,6,7}
// 	 f = f[1:4]
// 	 fmt.Println(f)
// 	 fmt.Println(f[:])

// 	//  The length of a slice is the number of elements it contains.

// // The capacity of a slice is the number of elements in the underlying array, counting from the first element in the slice.

// fmt.Println(len(f[1:4]),cap(f[1:4]))

// // A nil slice has a length and capacity of 0 and has no underlying array.
// var w []int
// fmt.Println(w,len(w),cap(w))
// if w==nil{
// 	fmt.Println("nil")
// }

// // Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
// // The make function allocates a zeroed array and returns a slice that refers to that array:

// m:= make([]int,5)
// printSlice("m",m)

// n:= make([]int,0,5)
// printSlice("n",n)

// x := n[:2]
// printSlice("x",x)

// g:= x[2:5]
// printSlice("g",g)

// // slice of slices
// board := [][]string{
// 	[]string{"_","_","_"},
// 	[]string{"_","_","_"},
// 	[]string{"_","_","_"},
// }

// 	board[0][0] = "X"
// 	board[2][2] = "O"
// 	board[1][2] = "X"
// 	board[1][0] = "O"
// 	board[0][2] = "X"

// 	for i:=0;i<len(board);i++{
// 		fmt.Printf("%s\n",strings.Join(board[i]," ") )
// 	}

// // Appending to a slice

// var t []int
// printSlice("t",t)

// // append works on nil slices
// t = append(t, 0)

// // The slice grows as needed.
// t = append(t,1)

// // We can add more than one element at a time.
// t = append(t, 2,3,4,5)

// // Range
// var pow = []int{1,2,3,4,5}

// for i,v := range pow{
// 	fmt.Printf("2**%d = %d\n",i,v )
// }

// }

// func printSlice(s string, x []int) {
// 	fmt.Printf("%s len=%d cap=%d %v\n",
// 		s, len(x), cap(x), x)
// }

///////////////////////////////
/* package main

// import (
// 	"fmt"
// 	"math"
// )

// type Vertex struct{
// 	Lat,Long float64
// }
// // Maps
// // A map maps keys to values.

// // The zero value of a map is nil. A nil map has no keys, nor can keys be added.

// // The make function returns a map of the given type, initialized and ready for use.
// var m map[string]Vertex

// // Map literals:
// // Map literals are like struct literals, but the keys are required.
// var z = map[string]Vertex{
// 	"Bell Labs": Vertex{
// 		40.687,-74.399,
// 	},
// 	"Google": Vertex{
// 		37.42202 , -1223.645,
// 	},
// }
// // If the top-level type is just a type name, you can omit it from the elements of the literal.

// var k = map[string]Vertex{
// 	"Bell Labs": {
// 		40.687,-74.399,
// 	},
// 	"Google": {
// 		37.42202 , -1223.645,
// 	},
// }
// func main() {
// 	m = make(map[string]Vertex)
// 	m["Bell Labs"] = Vertex{
// 		40.687,-74.399,
// 	}
// 	fmt.Println(m["Bell Labs"])
// 	fmt.Println(z)
// 	fmt.Println(k)

// 	p:= make(map[string]int)

// 	p["answer"] = 42
// 	fmt.Println("The value: ", p["answer"])

// 	p["answer"] = 48
// 	fmt.Println("The value: ", p["answer"])

// 	delete(p,"answer")
// 	fmt.Println("The value: ", p["answer"])

// 	v,ok:= p["answer"]
// 	fmt.Println("The value: ", v,"Present?",ok)

// 	// functions values

// hypot:= func(x,y float64)float64{
// 	return math.Sqrt(x*x+y*y)
// }
// fmt.Println(hypot(5,12))
// fmt.Println(compute(hypot))
// fmt.Println(compute(math.Pow))

// // func closures:Go functions may be closures. A closure is a function value that references variables from outside its body. The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.

// pos,neg := adder(), adder()
// for i:=0;i<10;i++{
// 	fmt.Println(pos(i), neg(-2*i))
// }

// }

// func compute(fn func(float64,float64)float64)float64{
// 	return fn (3,4)
// }

// func adder() func(int)int{
// 	sum:=0
// 	return func(x int)int{
// 		sum+=x
// 		return  sum
// 	}
// } */

/*
// Part-3
package main

import (
	"fmt"
	"math"
) 


type Vertex struct{
	X,Y float64
}

// Methods:
// Go does not have classes. However, you can define methods on types.
// A method is a function with a special receiver argument.
// Remember: a method is just a function with a receiver argument.

func (v Vertex) Abs() float64{
	return  math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Myfloat float64
// You can declare a method on non-struct types, too.
func (f Myfloat)Abs()float64{
	if f<0{
		return  float64(-f)
	}
	return  float64(f)
}
// You can only declare a method with a receiver whose type is defined in the same package as the method. You cannot declare a method with a receiver whose type is defined in another package 

// pointer receiver: This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)
// 🧩 Mental model
// Think:
// Value receiver → “read-only copy”
// Pointer receiver → “edit original”
func (v *Vertex) Scale(f float64){
	v.X = v.X*f
	v.Y = v.Y*f
}
// 👉 If a method has a pointer receiver (*T), only *T satisfies the interface; if it has a value receiver (T), both T and *T satisfy it.

func ScaleFunc(v *Vertex,  f float64){
	(*v).X = v.X*f
	v.Y = v.Y*f

}
// There are two reasons to use a pointer receiver.

// The first is so that the method can modify the value that its receiver points to.

// The second is to avoid copying the value on each method call. This can be more efficient if the receiver is a large struct, for example.

func main() {
	v:= Vertex{3,4}
	fmt.Println(v.Abs())
    
	f:= Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())
	v.Scale(3)
	ScaleFunc(&v,10)

	p := &Vertex{4,3}
	p.Scale(6)
	ScaleFunc(p,8)
	fmt.Println(v,p)
}
	*/