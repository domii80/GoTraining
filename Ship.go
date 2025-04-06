package main

import (
	"fmt"
)

type Ship struct {
	length      int
	startPos    coord
	orientation int
}

func NewShip(length int, startPos coord, orientation int) Ship {
	return Ship{length: length, startPos: startPos, orientation: orientation}
}

func (ship *Ship) setShipLength(len int) {
	if len > 0 {
		ship.length = len
	}
}

func (ship Ship) getCoordinates() []coord {
	coords := []coord{}
	x := 0
	y := 0
	var coeff = getCoefficient(ship.orientation)
	for i := range ship.length {
		x = ship.startPos.x + i*coeff.xd
		y = ship.startPos.y + i*coeff.yd
		coords = append(coords, coord{x: x, y: y})
	}
	for i := range len(coords) {
		fmt.Printf("coords vom Schiff - x: %v y: %v\n", coords[i].x, coords[i].y)
	}
	return coords
}

func (ship Ship) isCollidating(s Ship) bool {
	result := false
	thisShip := []coord{}
	otherShip := []coord{}
	x := 0
	y := 0
	var coeff = getCoefficient(ship.orientation)
	for i := range ship.length {
		x = ship.startPos.x + i*coeff.xd
		y = ship.startPos.y + i*coeff.yd
		thisShip = append(thisShip, coord{x: x, y: y})
	}
	coeff = getCoefficient(s.orientation)
	for i := range s.length {
		x = s.startPos.x + i*coeff.xd
		y = s.startPos.y + i*coeff.yd
		otherShip = append(otherShip, coord{x: x, y: y})
	}
	for i := range thisShip {
		for j := range otherShip {
			res := thisShip[i] == otherShip[j]
			if i == len(thisShip)-1 && j == len(otherShip)-1 {
				fmt.Printf("thisShip: %v\n", thisShip)
				fmt.Printf("otherShip: %v\n", otherShip)
				fmt.Printf("i %v, j %v: %v\n", i, j, res)
			}
			if res {
				result = true
				break
			}
		}
	}
	return result
}
