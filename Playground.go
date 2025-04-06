package main

type Playground [][]int

func NewPlayground(x int, y int) Playground {
	matrix := make(Playground, x)
	for i := range x {
		matrix[i] = make([]int, y)
	}
	return matrix
}

func (playground Playground) getValue(x int, y int) (int, error) {
	var err error = nil
	return playground[x][y], err
}

func (playGround Playground) setValue(x int, y int, v int) error {
	var err error = nil
	playGround[x][y] = v
	return err
}
