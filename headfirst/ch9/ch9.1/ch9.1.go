package main

import "fmt"

type Liters float64
type Gallons float64

func ToGallons(l Liters) Gallons {
	return Gallons(l * 0.264)
}
func ToLiters(g Gallons) Liters {
	return Liters(g * 3.785)
}
func main() {
	carFuel := Gallons(10.1)
	busFuel := Liters(240.2)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Println(carFuel)
	fmt.Println(busFuel)
}
