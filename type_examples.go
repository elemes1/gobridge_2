package main

import (
	"strconv"
)

type myInt int

func (i myInt) String() string {
	return "my int is" + strconv.Itoa(int(i))
}

type yourType myInt

var y myInt
var z yourType

type Person struct {
	name          string
	age           int
	getChildrenFn func() []string
}

// type myCancelFn context.CancelFunc
type myCancelFn func()

// type Person otherPkg.Person // in case I wanted to avoid exposing this external type to my immediate users

type car struct {
	numWheels int
	numDoors  int
}

type plane struct {
	numWings int
	numDoors int
}

type ford struct {
	car
	sound string
}

// func main() {
// 	var f myCancelFn
// 	f = func() {
// 		fmt.Println("cancel things")
// 	}
// 	f()
//
// 	// 	fmt.Println(x)
// 	// 	fmt.Println(y)
// 	// 	fmt.Println(z)
// 	//
// 	// 	fmt.Println(int(y) == int(z))
// 	// fmt.Println(struct{ name string }{name: "josh"})
// 	// i := myInt(6)
// 	// fmt.Println(i)
// 	//
// 	// p := Person{
// 	// 	name: "john",
// 	// 	age:  14,
// 	// }
// 	// fmt.Println(p)
//
// 	// 	f := ford{
// 	// 		car: car{
// 	// 			numWheels: 4,
// 	// 			numDoors:  2,
// 	// 		},
// 	// 		sound: "honk",
// 	// 	}
// 	//
// 	// 	fmt.Println(f.car.numDoors)
// 	// 	fmt.Println(f.numDoors)
//
// 	or := Order{
// 		OrderRec: OrderRec{
// 			ID: 1,
// 		},
// 		Customer: "",
// 	}
//
// 	// can call fields or methods on embedded types directly, or use the intermediary type to get it as well.
// 	// or.CustomerID
// 	// or.ID
// 	// or.OrderRec.ID
//
// 	b, err := json.Marshal(or)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(b))
// }

// package foo
// type MyType int
//
// package bar
// type MyType int

// OrderRec is a record in the order table.
type OrderRec struct {
	ID         int    `json:"my_id" db:"id"`
	CustomerID string `json:"customer_id" db:"customer_id"`
}

type Order struct {
	OrderRec
	Customer string `json:"customer" db:"customer"`
}

type stringer interface {
	string() string
}

type yoinker interface {
	stringer
	yoink() string
}

// // DON'T DO THIS !
// type lots interface {
// 	one
// 	two
// 	three
// }
