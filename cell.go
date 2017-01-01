package main

type Cell interface {
	IsAlive() bool
	Evolve(neighbours int) Cell
}

type NullCell struct{}

func (cell NullCell) IsAlive() bool {
	return false
}

func (cell NullCell) Evolve(neighbours int) Cell {
	if neighbours == 3 {
		return new(AliveCell)
	}

	return cell
}

type AliveCell struct{}

func (cell AliveCell) IsAlive() bool {
	return true
}

func (cell AliveCell) Evolve(neighbours int) Cell {
	if neighbours == 0 || neighbours >= 4 {
		return new(NullCell)
	}

	return cell
}
