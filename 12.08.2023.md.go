package main

	

// func main()  {
//     var a = add (5,4)
//     var b = add (10,4)

//     fmt.Println(a)
//     fmt.Println(b)
// }

// func add (x,y int) int{
//     return x + y
// }

// bu yerda ikkita sonni qoshib beryapti pastki qismida amal bajarilyapti tepa qismida esa ularga qiymat berilyapti

// package main

// import "fmt"

// func main()  {
//     var a = add(5,4)
//     var b = add(5,4)
//     fmt.Println(a)
//     fmt.Println(b)
// }

// func add(x,y int)int  {
//     var z int = x + y
//     return z
// }

// package main

// import "fmt"
// func add (x,y int) int{
//     return x + y
// }
// func main()  {
//     var f func(int,int)int = add
//     fmt.Println(f(3,4))
//     var x = f(4,5)

//	    fmt.Println(x)
//	}
// package main

// import "fmt"

// func add (x,y int) int{
//     return x + y
// }

// func multiply (x,y int) int{
//     return x * y
// }

// func action (n1,n2 int, operation func(int,int)int){
//     result := operation(n1,n2)
//     fmt.Println(result)
// }

// func main() {
//     action(10,25, add)
//     action(10,5, multiply)
// }

// func isEven(x int)bool  {
//     return x % 2 == 0
// }

// func isPositive(x int) bool  {
//     return x > 0
// }

// func sum (numbers []int, creative func(int)bool)int{
//     result := 0
//     for _,value := range numbers{
//         if(creative(value)){
//             result += value
//         }
//     }
//     return result
// }

// func main()  {
//     slice := []int {-2, 4, 3, -1, 7, -4, 23}

//     sumofEven := sum (slice,isEven)
//     fmt.Println(sumofEven)
//     sumofPositive := sum (slice,isPositive)
//     fmt.Println(sumofPositive)
// }

// ikkiga bo'linsin va 0dan katta son bo'lsa qo'shsin deyabdi va yangi o'zgaruvchi ya'ni creative ni olsin bu int uzatilgan raqamni boolean ga o'zgartirib bersin
// bu yerda ikkita o'zgaruvchi even (juft bo'lsa true qiymatni) positive (ijobiy bo'lsa true qiymatni qaytarsin)

// func add (x , y int )int{return x + y}
// func substract (x , y int)int {return x - y}
// func multiply (x , y int)int {return x * y}

// func selectFn(n int ) func (int,int)int{
//     if n == 1 {
//         return add
//     }else if n == 2 {
//         return substract
//     }else {
//        return multiply
//     }
// }

// func main () {
//     f := selectFn(1)
//     fmt.Println(f(3 , 4))

//     f = selectFn(2)
//     fmt.Println(f(4,3))
//     f = selectFn(3)
//     fmt.Println(f(2,5))
// }

// func action(x,y int , operation func(int,int) int ){
//     result := operation(x,y)
//     fmt.Println(result)
// }

// func main()  {
//     action(10,25, func(x,y int)int {return x + y })
//     action(10,25, func(x,y int)int {return x * y})
// }

// anonim funksiyalarga qaragnda argument sifatida ishlatish mumkin

// func selectFn(x int) (func(int,int) int){
//         if x == 1 {
//             return func (x,y int)int {return x + y}
//         }else if x == 2 {
//             return func (x,y int) int {return x - y}
//         }else {
//             return func(x,y int ) int {return x * y}
//         }
// }

// func main()  {
//     f := selectFn(1)
//     fmt.Println(f(3,4))
//     fmt.Println(f(3,4))
//     fmt.Println(f(3,4))
// }



// func main()  {
//     defer finish ()
//     fmt.Println("hello")
//     fmt.Println("world")
// }


// func finish () {
//     fmt.Println("Hayrullox")
// }


// bu yerda defer funksiyasi ishlamoqda .
// Deferfunksiyani vazifasi hamma funksiya ishlab bo'lgandan so'ng oxirida Defer funksiyasi ishlaydi




// func main()  {
//     fmt.Println(divide(15,5))
//     fmt.Println(divide(4,0))
// }


// func divide (x , y float64) float64 {
//     if y == 0{
//         panic("Division by zero!")
//     }
//     return  x / y
// }
















