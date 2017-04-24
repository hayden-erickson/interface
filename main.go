package main

import "fmt"

// SUPPLY provides values to BOOL_EXPR which compares them
type SUPPLY func(a, b interface{})
type BOOL_EXPR func(a, b interface{}) bool

func AND(exprs ...BOOL_EXPR) BOOL_EXPR {
	return func(a, b interface{}) bool {
		result := true
		for _, e := range exprs {
			result = result && e(a, b)
		}
		return result
	}
}

func OR(exprs ...BOOL_EXPR) BOOL_EXPR {
	return func(a, b interface{}) bool {
		result := false
		for _, e := range exprs {
			result = result || e(a, b)
		}
		return result
	}
}

func IF(expr BOOL_EXPR, isBetter SUPPLY) SUPPLY {
	return func(a, b interface{}) {
		if expr(a, b) {
			isBetter(a, b)
			return
		}
		isBetter(b, a)
	}
}

type Fabric int

var (
	leather   Fabric = 0
	suede     Fabric = 1
	polyester Fabric = 2
	cotton    Fabric = 3
)

type Engine struct {
	size int
}

type Upholstry struct {
	fabric Fabric
}

type Car struct {
	name      string
	engine    Engine
	upholstry Upholstry
	MPG       int
}

var (
	jaguar Car
	kia    Car
)

func toCar(A, B interface{}) (Car, Car) {
	carA, _ := A.(Car)
	carB, _ := B.(Car)
	return carA, carB
}

func hasBiggerEngine(A, B interface{}) bool {
	carA, carB := toCar(A, B)
	return carA.engine.size > carB.engine.size
}

func hasBetterUpholstry(A, B interface{}) bool {
	carA, carB := toCar(A, B)
	return carA.upholstry.fabric < carB.upholstry.fabric
}

func hasBetterMPG(A, B interface{}) bool {
	carA, carB := toCar(A, B)
	return carA.MPG > carB.MPG
}

func buyCar(A, B interface{}) {
	carA, _ := toCar(A, B)
	fmt.Printf("I would like to buy the %s", carA.name)
}

func main() {
	jaguar = Car{
		name: "jaguar",
		MPG:  25,
		upholstry: Upholstry{
			fabric: leather,
		},
		engine: Engine{
			size: 300,
		},
	}
	kia = Car{
		name: "kia soul",
		MPG:  35,
		upholstry: Upholstry{
			fabric: polyester,
		},
		engine: Engine{
			size: 150,
		},
	}

	whichToBuy := IF(OR(hasBiggerEngine, AND(hasBetterMPG, hasBetterUpholstry)), buyCar)
	whichToBuy(jaguar, kia)
}
