package main

// import "fmt"

// func main() {

//     var x int = 4
//     var p *int
//     p = &x
// 	    fmt.Println(p)
// }

// package main

// import "fmt"

// func main() {
// }
//     var x int = 4
//     var p *int  = &x
// 	fmt.Println("Address:", p)
// 	fmt.Println("Value:", *p)
// }

// package main
// import "fmt"

// func  main () {
// 	var (
// 		a float32 = 9
// 	b  float64 = 10
// 	d = (float64(a) + b)
// 	)

// 	fmt.Println(d)
// }

// package main

// import "fmt"

// func main () {
// 	var (
// 		a = 16
// 		b = 4
// 		d = (a + b)
// 		c = (a - b)
// 		e = (a * b)
// 		f = (a / b)
// 	)

// 	fmt.Println(d)
// 	fmt.Println(c)
// 	fmt.Println(e)
// 	fmt.Println(f)
// }

// bir nechta o'zgaruvchini berish

// package main
// import "fmt"

// func main () {

//    var (
// 	 	name = "Hayrullox"
// 	    age = 19
// 	)
// 	fmt.Println(name,age)

// }

// package  main
// import "fmt"

// func main () {
// 	var (
// 		a int64 = 15
// 		b int64 = 9
// 		d = (a + b)
// 	)

// 	fmt.Println(d)
// }

// package main
// import "fmt"

// func main () {
// 	var (
// 		a uint32 = 1
// 		b uint32 = 9
// 		d = (a + b)
// 	)
// 	fmt.Println(d)
// }

// package main
// import "fmt"

// func main () {
// 	var hello string
// 	hello = "Hello World"

// 	fmt.Println(hello)
// }

// package main
// import "fmt"

// func main () {
// 	var hello string ; hello = "Hello" ; fmt.Println(hello)
// }

// kodni bir qatorga yozib ketish

// package main
// import "fmt"

// func main () {
// 	var hello string = "World"

// 	fmt.Println(hello)
// }

// package main
// import "fmt"

// func main () {
// 	var hello string
// 	hello = "Hello world"
// 	fmt.Println(Hello)
// }

// o'zgaruvchini nomini kichik harfda e'lon qilinyapti Printlnda esa katta harf bilan berilyapti bu hato

// package main
// import "fmt"

// func main () {
// 	var hello string = "Hello "
// 	fmt.Println(hello)

// 	hello = "Hello go"
// 	fmt.Println(hello)

// 	hello = " Go Go Go Go "
// 	fmt.Println(hello)
// }

// o'zgaruvchilarning qiymati ko'p marta o'zgartirilishi mumkin

// package main
// import "fmt"

// func main () {
// const (
//     pi float64 = 3.1415
//     e float64 = 2.7182
// 	d = (e + pi)
// )

// fmt.Println(d)

// }

// package main

// import "fmt"

// func main () {
// 	const (
// 		a = 1
// 		b
// 		c = 2
// 		d
// 		f = 3
// 		g
// 	)
// 	fmt.Println(a, b, c, d, f, g)
// }

// package main
// import "fmt"

// func main () {
// 	var  (
// 		a = 6
// 		b = 8
// 		d = 10
// 		c = a * b * d
// 	)
// 	fmt.Println(c)
// }

// package main
//  import "fmt"

//  func main () {

// 	var (
// 		a int = 10
// 		b int = 4
// 		c = a * b
// 	)

// 	fmt.Println(c)
//  }

// package main

// import "fmt"

// func main ()  {
// 	var (
// 		a float32= 10
// 		b float32= 4.0
// 		c = (a / b)
// 	)

// 	fmt.Println(c)
// }

// package main
// import "fmt"

// func main () {
// 	var (
// 		a int = 35 % 3
// 	)
// 	fmt.Println(a)
// }

// bu yerda 35 % 3 bo'lganda 3 ni 33 qilib olyapti

// package main
// import "fmt"

// func main ()  {
// 	var a int = 8
// a--
// fmt.Println(a)
// }

//bu yerda bittaga kamaytirib beryapti

// package main
// import "fmt"

// func main ()  {
// 	var a int = 8
// a++
// fmt.Println(a)
// }

// bu yerda bittaga oshirib beryapti

// package main
// import "fmt"

// func main()  {
// 	var  (
// 		a int = 4
// 		b int = 2
// 		c bool = a == b
// 	)

// 	fmt.Println(c)
// }

// bu yerda ikkita amalni teng yoki teng emasligini tekshirib beryapti

// package main
// import "fmt"

// func main()  {
// 	var  (
// 		a int = 4
// 		b int = 2
// 		c bool = a > b
// 	)

// 	fmt.Println(c)
// }

// bu yerda ikkita amalni katta yoqmi shuni tekshiryapti

// package main
// import "fmt"

// func main()  {
// 	var  (
// 		a int = 4
// 		b int = 2
// 		c bool = a < b
// 	)

// 	fmt.Println(c)
// }

// bu yerda ikkita amalni kichik yoqmi shuni tekshiryapti

// package main
// import "fmt"

// func main ()  {
// 	var (
// 		a int = 6
// 		b int = 5
// 		c bool = a <= b
// 	)

// 	fmt.Println(c)
// }

// bu yerda ikkita sonning kichikmi yoki tengmi ekanligini tekshiryapti

// package main
// import "fmt"

// func main ()  {
// 	var (
// 		a int = 10
// 		b int = 9
// 		c bool = a >= b
// 	)
// 	fmt.Println(c)
// }

// bu yerda ikkita sonning tengmi yoki teng emasligini tekshiryapti

// package main
// import "fmt"

// func main()  {
// 	var a bool = false
//       var b bool = !a
//       var c bool = !b
// 	  fmt.Println(a,b,c)
// }

// package main
// import "fmt"

// func main ()  {
// 	var a bool = 4 > 5 && 6 > 8
// 	var b bool = 4 <= 5 && 10 > 8
// 	fmt.Println(a,b)
// }

// package main
// import "fmt"

// func main()  {
// 	var a  = [] int {1,2,3,4,5}
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	var a = [5]int{1, 2, 3, 4}

// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	var a = [...]int{1, 2, 3, 4, 5, 999, 0}
// 	fmt.Println(a)
// }

// package main

// import "fmt"

// func main() {
// 	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
// 	fmt.Println(numbers[0])
// 	fmt.Println(numbers[4])
// 	numbers[0] = 87
// 	fmt.Println(numbers[0])
// }

// package main

// import "fmt"

// func main() {
// 	colors := [3]string{2: "blue", 0: "red", 1: "black"}
// 	fmt.Println(colors[2])
// }

// package main

// import "fmt"

// func main() {
// 	a := 9
// 	b := 7
// 	if a < b {
// 		fmt.Println("a is less b")
// 	} else {
// 		fmt.Println("b is less a")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var a = 8
// 	if a == 9 {
// 		fmt.Println("a = 9")
// 	} else if a == 8 {
// 		fmt.Println("a = 8")
// 	} else if a == 7 {
// 		fmt.Println("a == 7")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	a := 7
// 	switch a {
// 	case 9:
// 		fmt.Println("a = 9")
// 	case 8:
// 		fmt.Println("a = 8")
// 	case 7:
// 		fmt.Println("a = 7")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	a := 7
// 	switch a {
// 	case 9:
// 		fmt.Println("a = 9")
// 	case 8:
// 		fmt.Println("a = 8")
// 	case 7:
// 		fmt.Println("a = 7")
// 	}

// }

// package main

// import "fmt"

// func main() {
// 	a := 5
// 	switch a {
// 	case 9:
// 		fmt.Println("a = 9")
// 	case 8:
// 		fmt.Println("a = 8")
// 	case 6, 5, 4:
// 		fmt.Println("a = 6 and 5 and 4")
// 	default:
// 		fmt.Println("bunday narsa yo'q")
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for i := 1; i < 10; i++ {
// 		fmt.Println(i * i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var i = 1
// 	for ; i < 10; i++ {
// 		fmt.Println(i + i)
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	var i = 1
// 	for i < 10 {
// 		fmt.Println(i)
// 		i++
// 	}
// }

// package main

// import "fmt"

// func main() {
// 	for i := 1; i < 10; i++ {
// 		for j := 1; j < 10; j++ {
// 			for g := 1; g < 10; g++ {

// 				fmt.Println(i * j * g)
// 			}
// 		}
// 	}
// }

// package main

// import "fmt"

//	func main() {
//		var users = [3]string{"Tom", "Alisa", "Xura"}
//		for index, value := range users {
//			fmt.Println(index, value)
//		}
//	}
// package main

// import "fmt"

// func main() {
// 	var users = [3]string{"Tom", "Alisa", "Xura"}
// 	for i := 0; i < len(users); i++ {
// 		fmt.Println(users[i])
// 	}

// }

// package main
// import "fmt"
// func main(){
// 	for i := 0 ; i < 10 ; i++{
// 		if i % 2 ==	0 {
// 			fmt.Println(i,"juft son")
// 		}else{
// 			fmt.Println(i,"toq son")
// 		}
// 	}
// }

// package main

// import "fmt"

// func main()  {
// 	var i = 0
// 	fmt.Scan("raqam kiriting: ")

// 	for i := 0; i < len(a); i++{
// 		fmt.Println(a[i])
// 	}
// }

// package main

// import "fmt"

// func main()  {
// 	hello()
// 	hello()
// 	hello()

// }
// func hello(){
// 	fmt.Println("hello world")
// }

// package main

// import "fmt"

// func main() {
// 	f := func(x , y int) int{return x + y}
// 	fmt.Println(f(3,4))
// }

// package main

// import "fmt"

// func main()  {
// 	for i := 0 ; i < 10 ; i++{
// 	for j := 0; j < 10; j++ {
// 		fmt.Print(i*j,"\t")
// 	}
// 	fmt.Println()
// 	}
// }

// package main

// import "fmt"

// func main()  {
// 	var users = [3]string {"ta","aki","xura"}
// 	for index := range users{
// 		fmt.Println(index)
// 	}
// }

// package main 

// import "fmt"


// func main()  {
// 	var users =[3]string{"tom","ali","xura"}
// 	for i := 0; i < len(users) ; i++{
// 		fmt.Println(users[i])
// 	}
// }
