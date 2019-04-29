package main
import (
	"fmt"
)

// type Person struct{
// 	name string
// 	age int
// }

type point struct{
	x,y int
}

type rect1 struct{
	leftUp,rightDown *point
}

func main() {
	// person := Person{"111",20}
	// person.name = "11111"
	// person.age = 18
	// var p1 *Person = new(Person) 
	// p1.name = "112233"
	// p1.age = 1
	// var p2 *Person = &Person{}
	// p2.name = "11"
	// p2.age = 111
	// (*p2).name = "33"
	// (*p2).age = 133
	// var p2 *Person = &Person{"44",144}
	// fmt.Println(*p2)
	rs1 := rect1{&point{10,20},&point{30,40}}
	// rs1.leftUp.x = 100
	fmt.Printf("%p\n",rs1.leftUp)	
	fmt.Printf("%p\n",rs1.rightDown)
}