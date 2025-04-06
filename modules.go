package main

import (
	"fmt"
	"os"
	"os/exec"
)

type ocean [20][20]int

type ship struct {
	length      int
	startPos    coord
	orientation int
}
type orientationCoefficient struct {
	xd int
	yd int
}

func ShipSinking() {
	clearScreen()
	var anzahl int
	fmt.Println("SCHIFFE VERSENKEN")
	fmt.Print("Wieviele Züge wollen Sie machen? ")
	fmt.Scan(&anzahl)
	ocean := createOcean()
	for range anzahl {
		displayOcean(ocean)
		var coord = aim()
		evalFire(coord, &ocean)
	}
	displayOcean(ocean)
	fmt.Println("")
	fmt.Println("Viel Spaß")
	fmt.Println("")
}

func clearScreen() {
	cmd := exec.Command("/usr/bin/clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func aim() coord {
	var c coord
	fmt.Printf("\nFeuer frei!\n")
	fmt.Printf("\nGeben Sie die x-Koordinate an: \n")
	fmt.Scan(&c.x)
	fmt.Println()
	fmt.Printf("Geben Sie die y-Koordinate an: \n")
	fmt.Scan(&c.y)
	fmt.Printf("\nFeuer auf Position (%v|%v)\n", c.x, c.y)
	return c
}

func evalFire(coord coord, ocean *ocean) {
	var x, y = coord.x, coord.y
	if ocean[x][y] == -1 {
		fmt.Println("Platsch")
	}
	if ocean[x][y] == 1 {
		fmt.Println("Treffer")
		ocean[x][y] = -1
	}
	if ocean[x][y] == 0 {
		fmt.Println("Daneben")
		ocean[x][y] = -1
	}
}

func displayOcean(ocean ocean) {
	for i := range len(ocean) {
		for j := range len(ocean[i]) {
			fmt.Printf(" %v", ocean[i][j])
		}
		fmt.Println()
	}
}

func createOcean() ocean {
	var ocean ocean
	var ship = ship{length: 5, startPos: coord{x: 9, y: 9}, orientation: 270}
	setShip(&ocean, ship)
	return ocean

}

func checkShip(ocean *ocean, ship ship) bool {
	result := true
	var shipsEnd coord
	coeff := getCoefficient(ship.orientation)
	shipsEnd.x = ship.startPos.x + coeff.xd*ship.length
	shipsEnd.y = ship.startPos.y + coeff.yd*ship.length
	if !(0 <= shipsEnd.x && shipsEnd.x <= 19) {
		result = false
	}
	return result
}

func getCoefficient(direction int) orientationCoefficient {
	m := getOrientationMatrix()
	oriCoeff := m[direction]
	return oriCoeff
}

func getOrientationMatrix() map[int]orientationCoefficient {
	return map[int]orientationCoefficient{
		0:   {1, 0},
		45:  {1, 1},
		90:  {0, 1},
		135: {-1, 1},
		180: {-1, 0},
		225: {-1, -1},
		270: {0, -1},
		315: {1, -1}}
}

func setShip(ocean *ocean, ship ship) bool {
	result := true
	x := 0
	y := 0
	var coeff = getCoefficient(ship.orientation)
	if checkShip(ocean, ship) {
		for i := range ship.length {
			x = ship.startPos.x + i*coeff.xd
			y = ship.startPos.y + i*coeff.yd
			ocean[x][y] = 1
		}
	} else {
		result = false
	}
	return result
}
