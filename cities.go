package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type City struct {
	Name   string
	X      int
	Y      int
	Images []*ebiten.Image
}

/*

REFERENCE CODE FOR REWORK

// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
	currentCalc   func() func(rect) float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func areaCalc() func(rect) float64  { return rect.area }
func perimCalc() func(rect) float64 { return rect.perim }

func main() {
	r := rect{width: 3, height: 4, currentCalc: perimCalc}
	fmt.Println(r.currentCalc()(r))
}
*/
