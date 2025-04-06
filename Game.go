package main

import "fmt"

type Game struct {
	playground Playground
	ships      []Ship
}

func NewGame(playground Playground, ships []Ship) Game {
	return Game{playground: playground, ships: ships}
}

func (game *Game) addShip(ship Ship) {
	game.ships = append(game.ships, ship)
}

func (game Game) getPlayground() Playground {
	playground := game.playground
	return playground
}

func (game Game) getShips() []Ship {
	ships := game.ships
	return ships
}

func (game Game) displayGame() {
	playground := game.getPlayground()
	ships := game.getShips()
	for i := range len(ships) {
		ship := ships[i]
		coeff := getCoefficient(ship.orientation)
		for l := range ship.length {
			x := ship.startPos.x + l*coeff.xd
			y := ship.startPos.y + l*coeff.yd
			playground[x][y] = 1
		}
	}
	for i := range len(playground) {
		for j := range len(playground[i]) {
			str := fmt.Sprintf("%2d ", playground[i][j])
			print(str)
		}
		fmt.Println()
	}
}
