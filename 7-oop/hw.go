package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.
type Coordinate struct {
	x, y float64
}

type Geom struct {
	start    Coordinate
	end      Coordinate
	distance float64
}

func (geom *Geom) CalculateDistance() float64 {
	if geom.start.x < 0 || geom.start.y < 0 {
		fmt.Println("Координаты начала не могут быть меньше нуля")
		return -1
	}
	if geom.end.x < 0 || geom.end.y < 0 {
		fmt.Println("Координаты конца не могут быть меньше нуля")
		return -1
	}
	// Если координаты совпадают - вернем 0
	if geom.start.x == geom.end.x && geom.start.y == geom.end.y {
		geom.distance = 0
		return geom.distance
	}
	// возврат расстояния между точками
	geom.distance = math.Sqrt(math.Pow(geom.end.x-geom.start.x, 2) + math.Pow(geom.end.y-geom.start.y, 2))
	return geom.distance
}
