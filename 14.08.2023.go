package main

import "fmt"

// func main()  {
// 	var x int = 4
// 	var p *int
// 	p = &x
// 	fmt.Println(p)
// }

// bu yerda manzilni olsih uchun & ishlatilyapti va o'zgaruvchi nomi (&x) ko'rsatkich x manzilini belgilamoqda
// x o'zgaruvchi int tipiga bog'liq  bo'lagni  uchun p ko'rsatkichi int tipidagi obyektga ishi=ora qlmoqda

// func main()  {
// 	var x int = 4
// 	var p *int
// 	p = &x
// 	fmt.Println("Adress" , p)
// 	fmt.Println("Value" , p)
// }

// func main()  {
// 	var x int = 4
// 	var p *int = &x
// 	*p = 25
// 	fmt.Println(x)
// }

// bu yerda ko'rsatkich yordamida ko'rsatkichda saqlangan qiymatni o'zgartirishimiz mumkin

// func main () {
// 	var pf *float64
// 	if pf != nil{
// 		fmt.Println(pf)
// 	}
// }

// func main() {

//     var x int = 4
//     var p *int  = &x
//     fmt.Println("Address:", p)
//     fmt.Println("Value:", *p)
// }

// func main()  {
// 	var pf * float64
// 	fmt.Println(pf)
// }

// obyekt manzili tayinlanmagan bo'lsa sukut bo'yicha nolga teng

// func main()  {
// 	p := new(int)
// 	fmt.Println("Value",*p)
// 	*p = 8
// 	fmt.Println(*p)
// }

// bu yerda new (yangi) o'zgaruvchi funksiya

// func change(x int)  {
// 	x = x * x
// }
// func main()  {
// 	d := 5
// 	fmt.Println("d before",d)
// 	change(d)
// 	fmt.Println(d)
// }

// func main()  {
// 	var x int = 5
// 	 var p *int
// 	 p = &x
// 	 fmt.Println(p)
// }

// func main()  {
// 	p :=new(int)
// 	fmt.Println("value",*p)
//      *p = 8
// 	 fmt.Println("value",*p)
// }

// func change(x *int)  {
// 	*x = (*x) * (*x)
// }

// func main()  {
// 	d := 5
// 	fmt.Println("d before" , d)
// 	change(&d)
//    fmt.Println(d)
// }

// type mile int

// func main()  {
// 	var distance mile = 5
// 	fmt.Println(distance)
// 	distance +=5
// 	fmt.Println(distance)
// }

// type mile int
// type kilometr int
// func distanceToEnemy (distance mile){

//     fmt.Println("расстояние для противника:")
//     fmt.Println(distance, "миль")
// }

// func main()  {
// 	var distance mile = 5
// 	distanceToEnemy(distance)
// 	var distance2 kilometr = 20
// 	fmt.Println(distance2)
// }

// type library [] string

// func change (lib library){
// 	for _, value := range lib{
// 		fmt.Println(value)
// 	}
// }

// func main()  {
// 	var myLibrary library = library{"book1","book2","book3"}
// 	printBooks(myLibrary)
// }

// type person struct{
// 	name string
// 	age int
// }

// func main()  {
//    var tom = person{name:"tom",age:41}
//    fmt.Println(tom.name)
//    fmt.Println(tom.age)

//    tom.age = 25
//    fmt.Println(tom.age)

//    tom.name ="hayrullox"
//    fmt.Println(tom.name)

// }

// type person struct {
// 	name string
// 	age  int
// }

// func main()  {
// 	tom := person{name:"hayri",age:27}
// 	var towPointer * person = &tom
// 	fmt.Println(tom.name)
// 	(*towPointer).age=32
// 	fmt.Println(tom.age)
// }

// type person struct {
// 	name string
// 	age int
// }

// func main () {
// 	tom := person {name: "hayri",age:27}
// 	var tooo * person = &tom
// 	fmt.Println(tom.name)
// 	(*tooo).age = 32
// 	fmt.Println(tom.age)
// }

// func main()  {
// 	var (
// 		a int = 8
// 		b int = 3
// 		c bool = a != b
// 		d bool = a !=8
// 	)

// 	fmt.Println(c)
// 	fmt.Println(d)
// }

// func createPointer(x int) *int{
//     p := new(int)
//     *p = x
//     return p
// }

// func main() {

//     p1 := createPointer(7)
//     fmt.Println("p1:", *p1)
//     p2 := createPointer(10)
//     fmt.Println("p2:", *p2)
//     p3 := createPointer(28)
//     fmt.Println("p3:", *p3)
// }

// type mile int

// func main() {

//     var distance mile = 5
//     fmt.Println(distance)
//     distance += 5
//     fmt.Println(distance)
// }

// type library []string

// func printBooks(lib library){

//     for index, value := range lib{

//         fmt.Println(index,value)
//     }
// }

// func main() {

//     var my library = library{"Book1", "Book2", "Book3"}
//     printBooks(my)
// }

// func main()  {
// 	var x = 4
// 	var p *int
// 	p = &x
// 	fmt.Println(*p)
// }

// func main()  {
// 	var x = 23
// 	var p *int
// 	p = &x
// 	fmt.Println("adress",p)
// 	fmt.Println("value",*p)
// }

// func main()  {
// 	var pf * float64
// 	if pf !=nil{
// 		fmt.Println("value",*pf)
// 	}
// }

// func main()  {
// 	p :=new(int)
// 	fmt.Println("value",p)
// 	*p = 9
// 	fmt.Println("value",*p)
// }

// func main()  {
// 	p := new(int)
// 	fmt.Println("value",p)
// 	*p = 8
// 	fmt.Println("value",*p)
// }

// func changeValue(x int)  {
// 	x = x * x
// }

// func main()  {
// 	d := 5
// 	fmt.Println("value",d)
// 	changeValue(d)
// 	fmt.Println(d)
// }

// new int ya'ni (yangi o'zgaruvchilar)

// func changeValue(x *int){
// 	(*x) = (*x) * (*x)
// }
// func main()  {
// 	d := 5
// 	fmt.Println(d)
// 	changeValue(&d)
// 	fmt.Println(d)
// }
// bu yerda x  ning kvadrati

// func changeValue(x *int)  {
//  *x = 49
//  fmt.Println(*x)
// }

// func main()  {
// 	d := 6
// 	fmt.Println(d)
// 	changeValue(&d)
// 	fmt.Println(&d)
// }

// func createPointer(x int) *int{
// 	p := new(int)
// 	*p = x
// 	return p
// }

// func main()  {
// 	p1 := createPointer(7)
// 	fmt.Println("p1",*p1)
// }

// func changeValue(x int){
//     x = x * x
// }
// func main() {

//     d := 5
//     fmt.Println("d before:", d)
//     changeValue(d)
//     fmt.Println("d after:", d)
// }

// type person struct{
// 	name string
// 	age int
// }

// func (p person) print(){
// 	fmt.Println(p.name)
// 	fmt.Println(p.age)
// }

// func (p person)eat (meal string)  {
// 	fmt.Println(p.name, meal)
// }

// func main()  {
// 	var tom = person{name: "hayrullo",age: 24}
// 	tom.print()
// 	tom.eat("makaronni yemoqchi")
// }

// type person struct{
// 	name string
// 	age int
// }

// func (p person)print()  {
// 	fmt.Println(p.name)
// 	fmt.Println(p.age)
// }

// func (p person)eat (meal string)  {
// 	fmt.Println(p.name,meal)
// }

// func main()  {
// 	var tom = person {name: "xura",age: 18}
// 	tom.print()
// 	tom.eat("bir")
// }

// type person struct{
// 	name string
// 	age int
// }

// func (p person) updateAge (newAge int) {
// 	p.age = newAge
// }

// func main()  {
// 	var tom = person {name: "xura",age: 18}
// 	fmt.Println(tom.age)
// 	tom.updateAge(33)
// 	fmt.Println(tom.age)
// }

// import "math"

// func main()  {
// 	fmt.Println(math.Sqrt(25))
// }



func main()  {
	examples(13)
}


func examples(son int)  {
	// son = 23 
	fmt.Println("son:",son)
}


